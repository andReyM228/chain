package chain_client

import (
	"context"
	coremoduletypes "github.com/andReyM228/one/x/core/types"
	"github.com/cosmos/cosmos-sdk/types"
	"strconv"
)

// Issue a new token to the given address
func (c *chainClient) Issue(ctx context.Context, toAddress string, amount int64, memo, denom, signBy string) (*types.TxResponse, error) {
	record, err := c.keyring.Key(signBy)
	if err != nil {
		return nil, err
	}

	fromAddress, err := record.GetAddress()
	if err != nil {
		return nil, err
	}

	msg := coremoduletypes.MsgIssue{
		Creator: fromAddress.String(),
		Amount:  strconv.FormatInt(amount, 10),
		Address: toAddress,
		Denom:   denom,
	}

	txResp, err := c.SignAndSend(ctx, record, memo, &msg)
	if err != nil {
		return nil, err
	}

	return txResp, nil
}

func (c *chainClient) Withdraw(ctx context.Context, address string, amount int64, memo, denom, signBy string) (*types.TxResponse, error) {
	record, err := c.keyring.Key(signBy)
	if err != nil {
		return nil, err
	}

	fromAddress, err := record.GetAddress()
	if err != nil {
		return nil, err
	}

	msg := coremoduletypes.MsgWithdraw{
		Creator: fromAddress.String(),
		Amount:  strconv.FormatInt(amount, 10),
		Address: address,
		Denom:   denom,
	}

	txResp, err := c.SignAndSend(ctx, record, memo, &msg)
	if err != nil {
		return nil, err
	}

	return txResp, nil
}
