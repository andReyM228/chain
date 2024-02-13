package chain_client

import (
	"context"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/cosmos/cosmos-sdk/crypto/keyring"
	"github.com/cosmos/cosmos-sdk/types"
	txTypes "github.com/cosmos/cosmos-sdk/types/tx"
	authTypes "github.com/cosmos/cosmos-sdk/x/auth/types"
)

func (c *chainClient) SignAndSend(ctx context.Context, signer *keyring.Record, memo string, msg ...types.Msg) (*types.TxResponse, error) {
	for _, m := range msg {
		err := m.ValidateBasic()
		if err != nil {
			return nil, err
		}
	}

	txBuilder := c.codec.GetTxConfig().NewTxBuilder()
	err := txBuilder.SetMsgs(msg...)
	if err != nil {
		return nil, err
	}

	txBuilder.SetGasLimit(c.gasLimit)

	signerAddress, err := signer.GetAddress()
	if err != nil {
		return nil, err
	}

	account, err := c.authQueryClient.Account(ctx, &authTypes.QueryAccountRequest{Address: signerAddress.String()})
	if err != nil {
		return nil, err
	}

	var authAccount authTypes.AccountI

	err = c.codec.UnpackAny(account.GetAccount(), &authAccount)
	if err != nil {
		return nil, err
	}

	factory := tx.Factory{}.
		WithKeybase(c.keyring).
		WithTxConfig(c.codec.GetTxConfig()).
		WithChainID(c.chainID).
		WithAccountNumber(authAccount.GetAccountNumber()).
		WithSequence(authAccount.GetSequence())

	txBuilder.SetMemo(memo)

	err = tx.Sign(factory, signer.Name, txBuilder, true)
	if err != nil {
		return nil, err
	}

	bytes, err := c.codec.GetTxConfig().TxEncoder()(txBuilder.GetTx())
	if err != nil {
		return nil, err
	}

	res, err := c.txClient.BroadcastTx(ctx, &txTypes.BroadcastTxRequest{TxBytes: bytes, Mode: txTypes.BroadcastMode_BROADCAST_MODE_SYNC})
	if err != nil {
		return nil, err
	}

	return res.TxResponse, nil
}
