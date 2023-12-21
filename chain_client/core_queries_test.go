package chain_client

import (
	"context"
	coremoduletypes "github.com/andReyM228/one/x/core/types"

	"github.com/cosmos/cosmos-sdk/crypto/keyring"

	"testing"
)

func Test_chainClient_GetStats(t *testing.T) {
	client := NewClient(ClientConfig{
		ChainID:     "one",
		BaseUrl:     "localhost:9090",
		KeyringType: keyring.BackendMemory,
		GasLimit:    100000,
	})

	type args struct {
		ctx             context.Context
		requestAllStats *coremoduletypes.QueryAllStatsRequest
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "success",
			args: args{
				ctx:             context.Background(),
				requestAllStats: &coremoduletypes.QueryAllStatsRequest{},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := client.GetStats(tt.args.ctx, tt.args.requestAllStats)
			if err != nil {
				t.Errorf("GetStats() error = %v", err)
				return
			}

			t.Log(got)
		})
	}
}
