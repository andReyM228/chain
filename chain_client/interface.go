package chain_client

import (
	"context"
	client_codec "github.com/andReyM228/one/chain_client/codec"
	"github.com/cosmos/cosmos-sdk/crypto/keyring"
	"github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/tx"
)

type (
	Client interface {
		Bank
		Core
		Auth
		Tx
		Unpacker
	}

	Bank interface {
		Send(ctx context.Context, toAddress string, amount int64, memo, denom, signBy string) (*types.TxResponse, error)
	}

	Core interface {
		Issue(ctx context.Context, toAddress string, amount int64, memo, denom, signBy string) (*types.TxResponse, error)
		Withdraw(ctx context.Context, address string, amount int64, memo, denom, signBy string) (*types.TxResponse, error)
	}

	Auth interface {
		AddAccount(ctx context.Context, name string, mnemonic string) error
		GenerateAccount(name string) (*keyring.Record, string, error)
	}

	Tx interface {
		GetTx(ctx context.Context, hash string) (*tx.GetTxResponse, error)
	}

	Unpacker interface {
		GetUnpacker() client_codec.Codec
	}
)
