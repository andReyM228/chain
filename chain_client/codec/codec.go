package client_codec

import (
	coremoduletypes "github.com/andReyM228/one/x/core/types"
	"github.com/cosmos/cosmos-sdk/client"
	sdkcodec "github.com/cosmos/cosmos-sdk/codec"
	cdctypes "github.com/cosmos/cosmos-sdk/codec/types"
	"github.com/cosmos/cosmos-sdk/std"
	"github.com/cosmos/cosmos-sdk/x/auth/tx"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"
)

type (
	encodingConfig struct {
		InterfaceRegistry cdctypes.InterfaceRegistry
		Marshaler         sdkcodec.ProtoCodecMarshaler
		TxConfig          client.TxConfig
		Amino             *sdkcodec.LegacyAmino
	}

	codec struct {
		enc encodingConfig
	}
)

func makeEncodingConfig() encodingConfig {
	interfaceRegistry := cdctypes.NewInterfaceRegistry()
	marshaler := sdkcodec.NewProtoCodec(interfaceRegistry)
	txConfig := tx.NewTxConfig(marshaler, tx.DefaultSignModes)
	amino := sdkcodec.NewLegacyAmino()

	return encodingConfig{
		InterfaceRegistry: interfaceRegistry,
		Marshaler:         marshaler,
		TxConfig:          txConfig,
		Amino:             amino,
	}
}

func NewCodec() Codec {
	cfg := makeEncodingConfig()

	std.RegisterInterfaces(cfg.InterfaceRegistry)
	coremoduletypes.RegisterInterfaces(cfg.InterfaceRegistry)
	banktypes.RegisterInterfaces(cfg.InterfaceRegistry)
	authtypes.RegisterInterfaces(cfg.InterfaceRegistry)

	return codec{enc: cfg}
}

func (c codec) GetTxConfig() client.TxConfig {
	return c.enc.TxConfig
}

func (c codec) GetInterfaceRegistry() cdctypes.InterfaceRegistry {
	return c.enc.InterfaceRegistry
}

func (c codec) UnpackAny(any *cdctypes.Any, receiver interface{}) error {
	return c.enc.InterfaceRegistry.UnpackAny(any, receiver)
}
