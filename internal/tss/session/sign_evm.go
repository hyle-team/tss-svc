package session

import (
	"context"
	"fmt"
	"sync"
	"time"

	connector "github.com/hyle-team/tss-svc/internal/core/connector"
	"github.com/hyle-team/tss-svc/internal/tss/finalizer"

	"github.com/hyle-team/tss-svc/internal/bridge"
	"github.com/hyle-team/tss-svc/internal/bridge/withdrawal"
	"github.com/hyle-team/tss-svc/internal/core"
	"github.com/hyle-team/tss-svc/internal/db"
	"github.com/hyle-team/tss-svc/internal/p2p"
	"github.com/hyle-team/tss-svc/internal/tss"
	"github.com/hyle-team/tss-svc/internal/tss/consensus"
	"github.com/hyle-team/tss-svc/internal/types"
	"github.com/pkg/errors"
	"gitlab.com/distributed_lab/logan/v3"
	"go.uber.org/atomic"
)

var _ p2p.TssSession = &EvmSigningSession{}

type EvmSigningSession struct {
	sessionId        *atomic.String
	idChangeListener func(oldId string, newId string)
	mu               *sync.RWMutex

	parties []p2p.Party
	self    tss.LocalSignParty
	db      db.DepositsQ

	params SigningSessionParams
	logger *logan.Entry

	coreConnector *connector.Connector
	fetcher       *bridge.DepositFetcher
	constructor   *withdrawal.EvmWithdrawalConstructor

	signingParty   *tss.SignParty
	consensusParty *consensus.Consensus[withdrawal.EvmWithdrawalData]
	finalizer      *finalizer.EvmFinalizer
}

func NewEvmSigningSession(
	self tss.LocalSignParty,
	parties []p2p.Party,
	params SigningSessionParams,
	db db.DepositsQ,
	logger *logan.Entry,
) *EvmSigningSession {
	sessionId := GetConcreteSigningSessionIdentifier(params.ChainId, params.Id)

	return &EvmSigningSession{
		sessionId: atomic.NewString(sessionId),
		mu:        &sync.RWMutex{},

		parties: parties,
		self:    self,
		db:      db,

		params: params,
		logger: logger,
	}
}

func (s *EvmSigningSession) WithDepositFetcher(fetcher *bridge.DepositFetcher) *EvmSigningSession {
	s.fetcher = fetcher
	return s
}

func (s *EvmSigningSession) WithConstructor(constructor *withdrawal.EvmWithdrawalConstructor) *EvmSigningSession {
	s.constructor = constructor
	return s
}

func (s *EvmSigningSession) WithCoreConnector(conn *connector.Connector) *EvmSigningSession {
	s.coreConnector = conn
	return s
}

func (s *EvmSigningSession) Run(ctx context.Context) error {
	runDelay := time.Until(s.params.StartTime)
	if runDelay <= 0 {
		return errors.New("target time is in the past")
	}

	nextSessionStartTime := s.params.StartTime
	for {
		s.mu.Lock()
		s.logger = s.logger.WithField("session_id", s.Id())
		s.consensusParty = consensus.New[withdrawal.EvmWithdrawalData](
			consensus.LocalConsensusParty{
				SessionId: s.Id(),
				Threshold: s.self.Threshold,
				Self:      s.self.Address,
				ChainId:   s.params.ChainId,
			},
			s.parties,
			s.db,
			s.fetcher,
			s.constructor,
			s.logger.WithField("phase", "consensus"),
		)
		s.signingParty = tss.NewSignParty(s.self, s.Id(), s.logger.WithField("phase", "signing"))
		s.finalizer = finalizer.NewEVMFinalizer(s.db, s.coreConnector, s.logger.WithField("phase", "finalizing"))
		s.mu.Unlock()

		s.logger.Info(fmt.Sprintf("waiting for next signing session %s to start in %s", s.Id(), time.Until(nextSessionStartTime)))

		select {
		case <-ctx.Done():
			s.logger.Info("signing session cancelled")
			return nil
		case <-time.After(time.Until(nextSessionStartTime)):
			nextSessionStartTime = time.Now().Add(tss.BoundarySigningSession)
		}

		s.logger.Info(fmt.Sprintf("signing session %s started", s.Id()))
		if err := s.runSession(ctx); err != nil {
			s.logger.WithError(err).Error("failed to run signing session")
		}
		s.logger.Info(fmt.Sprintf("signing session %s finished", s.Id()))

		s.incrementSessionId()
	}
}

func (s *EvmSigningSession) runSession(ctx context.Context) error {
	// consensus phase
	consensusCtx, consCtxCancel := context.WithTimeout(ctx, tss.BoundaryConsensus)
	defer consCtxCancel()

	s.consensusParty.Run(consensusCtx)
	result, err := s.consensusParty.WaitFor()
	if err != nil {
		if !errors.Is(err, context.DeadlineExceeded) {
			return errors.Wrap(err, "consensus phase error occurred")

		}
		if err = ctx.Err(); err != nil {
			s.logger.Info("session cancelled")
			return nil
		}
		if err = consensusCtx.Err(); err != nil {
			if result.SigData != nil {
				s.logger.Info("local party is not the signer in the current session")
			} else {
				s.logger.Info("consensus phase timeout")
			}
			return nil
		}
	}
	if result.SigData == nil {
		s.logger.Info("no data to sign in the current session")
		return nil
	}
	if err = s.db.UpdateStatus(result.SigData.DepositIdentifier(), types.WithdrawalStatus_WITHDRAWAL_STATUS_PROCESSING); err != nil {
		return errors.Wrap(err, "failed to update deposit status")
	}
	if result.Signers == nil {
		s.logger.Info("local party is not the signer in the current session")
		return nil
	}

	// signing phase
	signingCtx, sigCtxCancel := context.WithTimeout(ctx, tss.BoundarySign)
	defer sigCtxCancel()

	s.signingParty.
		WithParties(result.Signers).
		WithSigningData(result.SigData.ProposalData.SigData).
		Run(signingCtx)
	signature := s.signingParty.WaitFor()
	if signature == nil {
		return errors.New("signing phase error occurred")
	}

	// finalization phase
	finalizerCtx, finalizerCancel := context.WithTimeout(ctx, tss.BoundaryFinalize)
	defer finalizerCancel()

	err = s.finalizer.
		WithData(result.SigData).
		WithSignature(signature).
		WithLocalPartyProposer(s.self.Address == result.Proposer).
		Finalize(finalizerCtx)
	if err != nil {
		return errors.Wrap(err, "finalizer phase error occurred")
	}

	return nil
}

func (s *EvmSigningSession) Id() string {
	return s.sessionId.Load()
}

func (s *EvmSigningSession) incrementSessionId() {
	prevSessionId := s.Id()
	nextSessionId := IncrementSessionIdentifier(prevSessionId)
	s.sessionId.Store(nextSessionId)
	s.idChangeListener(prevSessionId, nextSessionId)
}

func (s *EvmSigningSession) Receive(request *p2p.SubmitRequest) error {
	if request == nil {
		return errors.New("nil request")
	}

	switch request.Type {
	case p2p.RequestType_RT_PROPOSAL, p2p.RequestType_RT_ACCEPTANCE, p2p.RequestType_RT_SIGN_START:
		s.mu.RLock()
		err := s.consensusParty.Receive(request)
		s.mu.RUnlock()

		return err
	case p2p.RequestType_RT_SIGN:
		data := &p2p.TssData{}
		if err := request.Data.UnmarshalTo(data); err != nil {
			return errors.Wrap(err, "failed to unmarshal TSS request signingData")
		}

		sender, err := core.AddressFromString(request.Sender)
		if err != nil {
			return errors.Wrap(err, "failed to parse sender address")
		}

		s.mu.RLock()
		s.signingParty.Receive(sender, data)
		s.mu.RUnlock()

		return nil
	default:
		return errors.New(fmt.Sprintf("unsupported request type %s from '%s'", request.Type, request.Sender))
	}
}

func (s *EvmSigningSession) RegisterIdChangeListener(f func(oldId string, newId string)) {
	s.idChangeListener = f
}
