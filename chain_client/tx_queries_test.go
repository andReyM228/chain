package chain_client

import (
	"context"
	"github.com/cosmos/cosmos-sdk/crypto/keyring"
	"testing"
)

func Test_chainClient_GetTx(t *testing.T) {
	client := NewClient(ClientConfig{
		ChainID:     "one",
		BaseUrl:     "localhost:9090",
		KeyringType: keyring.BackendMemory,
		GasLimit:    100000,
	})

	type args struct {
		ctx  context.Context
		hash string
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "success",
			args: args{
				ctx:  context.Background(),
				hash: "1659B27EC00BD947EE51E089916F46DF7F187B25FE626ADF37C2F25CBC271908",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := client.GetTx(tt.args.ctx, tt.args.hash)
			if err != nil {
				t.Errorf("GetStats() error = %v", err)
				return
			}

			t.Log(got)
		})
	}
}
