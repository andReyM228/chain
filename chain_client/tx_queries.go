package chain_client

import (
	"context"
	"github.com/cosmos/cosmos-sdk/types/tx"
)

func (c chainClient) GetTx(ctx context.Context, hash string) (*tx.GetTxResponse, error) {
	trx, err := c.txClient.GetTx(ctx, &tx.GetTxRequest{Hash: hash})
	if err != nil {
		return nil, err
	}

	return trx, nil
}
