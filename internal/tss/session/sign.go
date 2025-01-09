package session

import (
	"context"
	"fmt"
	"github.com/bnb-chain/tss-lib/v2/common"
	"github.com/hyle-team/tss-svc/internal/core"
	"github.com/hyle-team/tss-svc/internal/p2p"
	"github.com/hyle-team/tss-svc/internal/tss"
	"github.com/pkg/errors"
	"gitlab.com/distributed_lab/logan/v3"
	"sync"
	"time"
)

type SigningSessionParams struct {
	Id        string
	StartTime time.Time
	Threshold int
}

type SigningSession struct {
	params SigningSessionParams
	logger *logan.Entry
	wg     *sync.WaitGroup

	connectedPartiesCount func() int
	partiesCount          int

	signingParty interface {
		Run(ctx context.Context)
		WaitFor() *common.SignatureData
		Receive(sender core.Address, data *p2p.TssData)
	}

	data   string
	result *common.SignatureData
	err    error
}

func NewSigningSession(self tss.LocalSignParty, params SigningSessionParams, logger *logan.Entry, parties []p2p.Party, data string, connectedPartiesCountFunc func() int) *SigningSession {
	return &SigningSession{
		params:                params,
		wg:                    &sync.WaitGroup{},
		logger:                logger,
		connectedPartiesCount: connectedPartiesCountFunc,
		partiesCount:          len(parties),
		data:                  data,
		signingParty:          tss.NewSignParty(self, parties, data, params.Id, logger),
	}
}

func (s *SigningSession) Run(ctx context.Context) error {
	runDelay := time.Until(s.params.StartTime)
	if runDelay <= 0 {
		return errors.New("target time is in the past")
	}

	s.logger.Info(fmt.Sprintf("signing session will start in %s", runDelay))

	select {
	case <-ctx.Done():
		s.logger.Info("signing session cancelled")
		return nil
	case <-time.After(runDelay):
	}

	if s.connectedPartiesCount() != s.partiesCount {
		return errors.New("cannot start signing session: not all parties connected")
	}

	s.wg.Add(1)
	go s.run(ctx)
	return nil
}

func (s *SigningSession) run(ctx context.Context) {
	defer s.wg.Done()

	boundedCtx, cancel := context.WithTimeout(ctx, BoundarySigningSession)
	defer cancel()

	s.signingParty.Run(boundedCtx)
	s.result = s.signingParty.WaitFor()
	s.logger.Info("signing session finished")
	if s.result != nil {
		return
	}

	if err := boundedCtx.Err(); err != nil {
		s.err = err
	} else {
		s.err = errors.New("signing session error occurred")
	}
}

func (s *SigningSession) WaitFor() (*common.SignatureData, error) {
	s.wg.Wait()
	return s.result, s.err
}

func (s *SigningSession) Id() string {
	return s.params.Id
}

func (s *SigningSession) Receive(request *p2p.SubmitRequest) error {
	if request.Type != p2p.RequestType_SIGN {
		return errors.New("invalid request type")
	}

	var data *p2p.TssData

	if err := request.Data.UnmarshalTo(data); err != nil {
		return errors.Wrap(err, "failed to unmarshal TSS request data")
	}

	sender, _ := core.AddressFromString(request.Sender)
	s.signingParty.Receive(sender, data)
	return nil
}

// RegisterIdChangeListener is a no-op for SigningSession
func (s *SigningSession) RegisterIdChangeListener(func(oldId, newId string)) {}