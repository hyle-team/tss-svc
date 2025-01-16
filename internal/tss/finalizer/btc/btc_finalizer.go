package btc

import (
	"context"
	"github.com/bnb-chain/tss-lib/v2/common"
	"github.com/hyle-team/tss-svc/internal/bridge/client/btc"
	"github.com/hyle-team/tss-svc/internal/bridge/types"
	"github.com/hyle-team/tss-svc/internal/db"
	"github.com/pkg/errors"
)

type Finalizer struct {
	chainClient types.Proxy
	// TODO: add Bridge core connector
	// TODO: add Bitcoin network connection

	data    []byte
	rawData db.DepositData
	db      db.DepositsQ
}

func NewFinalizer(chainClient types.Proxy, data []byte, signatureData *common.SignatureData, db db.DepositsQ, deposit db.DepositData) *Finalizer {
	deposit.Signature = signatureData.Signature
	return &Finalizer{
		chainClient: chainClient,
		data:        data,
		db:          db,
		rawData:     deposit,
	}
}

func (bt *Finalizer) Run(_ context.Context) error {
	btcProxy, ok := bt.chainClient.(btc.BridgeProxy)
	if !ok {
		return errors.Wrap(errors.New("invalid proxy type"), "finalization failed")
	}
	bt.chainClient = btcProxy

	// Save the data with signature to db
	err := bt.db.SetDepositSignature(bt.rawData)
	if err != nil {
		return errors.Wrap(err, "finalization failed")
	}

	// Using core connector pass withdrawal tx info to Bridge core
	// TODO: Implement passing data to core

	//
	return nil

}
