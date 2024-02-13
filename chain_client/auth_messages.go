package chain_client

import (
	"context"
	"github.com/cosmos/cosmos-sdk/crypto/hd"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (c *chainClient) AddAccount(ctx context.Context, name string, mnemonic string) error {
	_, err := c.keyring.NewAccount(name, mnemonic, "", sdk.FullFundraiserPath, hd.Secp256k1)
	if err != nil {
		return err
	}

	return nil
}
