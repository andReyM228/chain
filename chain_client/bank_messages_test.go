package chain_client

import (
	"context"
	"github.com/cosmos/cosmos-sdk/crypto/keyring"
	"testing"
)

func Test_chainClient_Send(t *testing.T) {
	client := NewClient(ClientConfig{
		ChainID:     "one",
		BaseUrl:     "localhost:9090",
		KeyringType: keyring.BackendMemory,
		GasLimit:    100000,
	})

	err := client.AddAccount(context.Background(), "bob", "panda property decide blush omit cross oak deposit banana wise warrior burden depth oil beach stumble swim tattoo ancient next grab merge manual remain")
	if err != nil {
		t.Fatal(err)
		return
	}

	type args struct {
		ctx       context.Context
		toAddress string
		amount    int64
		memo      string
		signBy    string
		denom     string
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "success",
			args: args{
				ctx:       context.Background(),
				toAddress: "cosmos1yhmva39w3q7uy593u8h58xaxmtglh6h94tq8ch",
				amount:    10,
				memo:      "test",
				denom:     DenomOne,
				signBy:    "bob",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := client.Send(tt.args.ctx, tt.args.toAddress, tt.args.amount, tt.args.memo, tt.args.denom, tt.args.signBy)
			if err != nil {
				t.Fatal(err)
				return
			}

			t.Log("Got: ", got)
		})
	}
}

func Test_chainClient_Issue(t *testing.T) {
	c := NewClient(ClientConfig{
		ChainID:     "one",
		BaseUrl:     "localhost:9090",
		KeyringType: keyring.BackendMemory,
		GasLimit:    100000,
	})

	err := c.AddAccount(context.Background(), "bob", "panda property decide blush omit cross oak deposit banana wise warrior burden depth oil beach stumble swim tattoo ancient next grab merge manual remain")
	if err != nil {
		t.Fatal(err)
		return
	}

	type args struct {
		ctx       context.Context
		toAddress string
		amount    int64
		memo      string
		denom     string
		signBy    string
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "success",
			args: args{
				ctx:       context.Background(),
				toAddress: "cosmos1du7e552z3r88mr59wutp0mqxse0p43n7s7fu44",
				amount:    10,
				memo:      "test",
				denom:     DenomOne,
				signBy:    "bob",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := c.Issue(tt.args.ctx, tt.args.toAddress, tt.args.amount, tt.args.memo, tt.args.denom, tt.args.signBy)
			if err != nil {
				t.Fatal(err)
				return
			}

			t.Log("Got: ", got)
		})
	}
}

func Test_chainClient_Withdraw(t *testing.T) {
	c := NewClient(ClientConfig{
		ChainID:     "one",
		BaseUrl:     "localhost:9090",
		KeyringType: keyring.BackendMemory,
		GasLimit:    100000,
	})

	err := c.AddAccount(context.Background(), "bob", "panda property decide blush omit cross oak deposit banana wise warrior burden depth oil beach stumble swim tattoo ancient next grab merge manual remain")
	if err != nil {
		t.Fatal(err)
		return
	}

	type args struct {
		ctx       context.Context
		toAddress string
		amount    int64
		memo      string
		denom     string
		signBy    string
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "success",
			args: args{
				ctx:       context.Background(),
				toAddress: "cosmos1du7e552z3r88mr59wutp0mqxse0p43n7s7fu44",
				amount:    10,
				memo:      "test",
				denom:     DenomOne,
				signBy:    "bob",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := c.Withdraw(tt.args.ctx, tt.args.toAddress, tt.args.amount, tt.args.memo, tt.args.denom, tt.args.signBy)
			if err != nil {
				t.Fatal(err)
				return
			}

			t.Log("Got: ", got)
		})
	}
}
