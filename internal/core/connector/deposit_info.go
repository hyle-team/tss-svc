package connector

import (
	"context"
	bridgetypes "github.com/hyle-team/bridgeless-core/v12/x/bridge/types"
	"github.com/hyle-team/tss-svc/internal/types"
	"github.com/pkg/errors"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

var errTxNotFound = status.Error(codes.NotFound, "transaction not found")

func (c *Connector) GetDepositInfo(identifier *types.DepositIdentifier) (*bridgetypes.Transaction, error) {
	req := bridgetypes.QueryTransactionByIdRequest{
		ChainId: identifier.ChainId,
		TxHash:  identifier.TxHash,
		TxNonce: uint64(identifier.TxNonce),
	}

	resp, err := c.querier.TransactionById(context.Background(), &req)
	if err != nil {
		if errors.Is(err, errTxNotFound) {
			return nil, nil
		}

		return nil, errors.Wrap(err, "failed to get token info")
	}

	return &resp.Transaction, nil
}
