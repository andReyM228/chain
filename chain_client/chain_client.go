package chain_client

import (
	client_codec "github.com/andReyM228/chain_client/codec"
	coremoduletypes "github.com/andReyM228/one/x/core/types"
	"github.com/cosmos/cosmos-sdk/crypto/keyring"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/tx"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	"github.com/cosmos/cosmos-sdk/codec"
	cdctypes "github.com/cosmos/cosmos-sdk/codec/types"
	cryptocodec "github.com/cosmos/cosmos-sdk/crypto/codec"
)

const (
	DenomOne = "one"
)

type (
	ClientConfig struct {
		ChainID     string `yaml:"chain_id"`
		BaseUrl     string `yaml:"base_url"`
		KeyringType string `yaml:"keyring_type"`
		GasLimit    uint64 `yaml:"gas_limit"`
	}

	chainClient struct {
		chainID         string
		baseUrl         string
		keyringType     string
		gasLimit        uint64
		keyring         keyring.Keyring
		txClient        tx.ServiceClient
		coreQueryClient coremoduletypes.QueryClient
		bankQueryClient banktypes.QueryClient
		codec           client_codec.Codec
		authQueryClient authtypes.QueryClient
	}
)

func NewClient(cfg ClientConfig) Client {
	grpcConnection := grpc.WithTransportCredentials(insecure.NewCredentials())

	conn, err := grpc.Dial(cfg.BaseUrl, grpcConnection)
	if err != nil {
		return nil
	}

	client := &chainClient{
		chainID:         cfg.ChainID,
		baseUrl:         cfg.BaseUrl,
		keyringType:     cfg.KeyringType,
		gasLimit:        cfg.GasLimit,
		txClient:        tx.NewServiceClient(conn),
		coreQueryClient: coremoduletypes.NewQueryClient(conn),
		bankQueryClient: banktypes.NewQueryClient(conn),
		codec:           client_codec.NewCodec(),
		authQueryClient: authtypes.NewQueryClient(conn),
	}

	client.keyring, err = keyring.New(sdk.KeyringServiceName(), client.keyringType, "", nil, getProtoCodec())
	if err != nil {
		return nil
	}

	return client
}

func getProtoCodec() codec.Codec {
	registry := cdctypes.NewInterfaceRegistry()

	coremoduletypes.RegisterInterfaces(registry)
	banktypes.RegisterInterfaces(registry)
	authtypes.RegisterInterfaces(registry)

	cryptocodec.RegisterInterfaces(registry)
	return codec.NewProtoCodec(registry)
}
