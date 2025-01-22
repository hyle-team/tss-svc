package api

import (
	"context"
	"github.com/go-chi/chi/v5"
	"github.com/gorilla/websocket"
	"github.com/hyle-team/tss-svc/internal/api/common"
	"github.com/hyle-team/tss-svc/internal/api/ctx"
	"github.com/hyle-team/tss-svc/internal/api/requests"
	apiTypes "github.com/hyle-team/tss-svc/internal/api/types"
	database "github.com/hyle-team/tss-svc/internal/db"
	types "github.com/hyle-team/tss-svc/internal/types"
	"gitlab.com/distributed_lab/ape"
	"gitlab.com/distributed_lab/ape/problems"
	"google.golang.org/protobuf/encoding/protojson"
	"net/http"
	"slices"
	"strconv"
	"time"
)

const (
	paramOriginTxId = "origin_tx_id"
	pollingPeriod   = 1 * time.Second
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin:     func(r *http.Request) bool { return true },
}

func CheckWithdrawalWs(w http.ResponseWriter, r *http.Request) {
	var (
		ctxt   = r.Context()
		chains = ctx.Chains(ctxt)
	)

	//get incoming params
	chainId, txHash, txNonce, err := getUrlParams(w, r, chains)
	if err != nil {
		ape.RenderErr(w, problems.BadRequest(err)...)
	}

	deposit, err := requests.CheckTx(ctxt, &types.DepositIdentifier{
		TxHash:  txHash,
		TxNonce: int64(txNonce),
		ChainId: chainId,
	})
	if err != nil {
		ape.RenderErr(w, problems.BadRequest(err)...)
	}

	depositIdentifier := deposit.DepositIdentifier

	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		ape.RenderErr(w, problems.BadRequest(err)...)
		ctx.Logger(ctxt).WithError(err).Debug("websocket upgrade error")
		return
	}

	gracefulClose := make(chan struct{})
	go watchConnectionClosing(ws, gracefulClose)
	watchWithdrawalStatus(ctxt, ws, gracefulClose, depositIdentifier)
}

func watchConnectionClosing(ws *websocket.Conn, done chan struct{}) {
	defer close(done)

	for {
		// collecting errors and close message to signalize writer.
		// note: `ReadMessage` is a blocking operation.
		// note: infinite loop will be broken either by close message or
		//       closed connection by writer goroutine, which immediately
		//       sends an error to a reader.
		mt, _, err := ws.ReadMessage()
		if err != nil || mt == websocket.CloseMessage {
			break
		}
	}
}

func watchWithdrawalStatus(ctxt context.Context, ws *websocket.Conn, connClosed chan struct{}, id database.DepositIdentifier) {
	defer func() { _ = ws.Close() }()

	var (
		db                                         = ctx.DB(ctxt)
		logger                                     = ctx.Logger(ctxt)
		prevStatus          types.WithdrawalStatus = -1
		cancelled, graceful bool
		ticker              = time.NewTicker(pollingPeriod)

		// function to repeat iteration after some period or break the loop
		// in case of a cancellation signal. If the signal is produced by
		// app context, websocket connection would be closed gracefully with
		// the corresponding `CloseGoingAway` status
		tillCancel = func() {
			select {
			case <-connClosed:
				cancelled = true
			case <-ctxt.Done():
				cancelled, graceful = true, true
			case <-ticker.C:
				// doing nothing, just waiting some period
			}
		}
	)

	defer ticker.Stop()

	// fast-starting without waiting for initial tick.
	// This shenanigan is just a classic `do-while` construction
	// with missing init statement and condition expression.
	// Using `tillCancel` as a post statement allows us to run
	// first iteration without waiting for ticker to tick.
	for ; ; tillCancel() {
		if cancelled {
			if graceful {
				_ = ws.WriteMessage(websocket.CloseMessage, websocket.FormatCloseMessage(websocket.CloseGoingAway, "Server shutting down"))
			}
			return
		}

		deposit, err := db.Get(id)
		if err != nil {
			logger.WithError(err).Error("failed to get deposit")
			_ = ws.WriteMessage(websocket.CloseMessage, websocket.FormatCloseMessage(websocket.CloseInternalServerErr, "Internal server error"))
			return
		}
		if deposit == nil {
			_ = ws.WriteMessage(websocket.CloseMessage, websocket.FormatCloseMessage(4004, "deposit not found"))
			return
		}
		logger.Info(deposit.WithdrawalStatus.String())

		//poll until our tx won`t be finished
		if deposit.WithdrawalStatus == prevStatus {
			continue
		}

		raw, err := protojson.Marshal(common.ToStatusResponse(deposit))
		if err != nil {
			logger.WithError(err).Error("failed to marshal deposit status")
			_ = ws.WriteMessage(websocket.CloseMessage, websocket.FormatCloseMessage(websocket.CloseInternalServerErr, "Internal server error"))
			return
		}
		if err = ws.WriteMessage(websocket.TextMessage, raw); err != nil {
			logger.WithError(err).Error("failed to write message to websocket")
			_ = ws.WriteMessage(websocket.CloseMessage, websocket.FormatCloseMessage(websocket.CloseInternalServerErr, "Internal server error"))
			return
		}

		// is it a time for websocket closing
		if slices.Contains(database.FinalWithdrawalStatuses, deposit.WithdrawalStatus) {
			err = ws.WriteMessage(websocket.CloseMessage, websocket.FormatCloseMessage(websocket.CloseNormalClosure, ""))
			if err != nil {
				logger.WithError(err).Error("failed to send close msg after finish")
			}
			return
		}

		prevStatus = deposit.WithdrawalStatus
	}
}

func getUrlParams(w http.ResponseWriter, r *http.Request, chains apiTypes.ChainsMap) (chainId string, txHash string, txNonce int, err error) {
	chainId = chi.URLParam(r, "chain_id")
	if _, ok := chains[chainId]; !ok {
		return "", "", 0, apiTypes.ErrInvalidChainId
	}
	txHash = chi.URLParam(r, "tx_hash")
	if len(txHash) < 3 {
		return "", "", 0, apiTypes.ErrInvalidTxHash
	}
	txNonce, err = strconv.Atoi(chi.URLParam(r, "tx_nonce"))
	if err != nil || txNonce < 0 {
		return "", "", 0, apiTypes.ErrInvalidTxNonce
	}
	return chainId, txHash, txNonce, nil
}
