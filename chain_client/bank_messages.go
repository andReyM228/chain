package chain_client

import (
	"context"
	"github.com/cosmos/cosmos-sdk/types"
	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"
)

func (c *chainClient) Send(ctx context.Context, toAddress string, amount int64, memo, denom, signBy string) (*types.TxResponse, error) {
	record, err := c.keyring.Key(signBy)
	if err != nil {
		return nil, err
	}

	fromAddress, err := record.GetAddress()
	if err != nil {
		return nil, err
	}

	msg := banktypes.MsgSend{
		FromAddress: fromAddress.String(),
		ToAddress:   toAddress,
		Amount:      types.NewCoins(types.NewInt64Coin(denom, amount)),
	}

	txResp, err := c.SignAndSend(ctx, record, memo, &msg)
	if err != nil {
		return nil, err
	}

	return txResp, nil
}
