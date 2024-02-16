package chain_client

import (
	"context"
	"github.com/cosmos/cosmos-sdk/crypto/hd"
	"github.com/cosmos/cosmos-sdk/crypto/keyring"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (c *chainClient) AddAccount(ctx context.Context, name string, mnemonic string) error {
	_, err := c.keyring.NewAccount(name, mnemonic, "", sdk.FullFundraiserPath, hd.Secp256k1)
	if err != nil {
		return err
	}

	return nil
}

// GenerateAccount - generates a new account by name.
func (c *chainClient) GenerateAccount(name string) (*keyring.Record, string, error) {
	r, mnemonic, err := c.keyring.NewMnemonic(name, keyring.English, sdk.FullFundraiserPath, keyring.DefaultBIP39Passphrase, hd.Secp256k1)
	if err != nil {
		return nil, "", err
	}

	return r, mnemonic, nil
}
