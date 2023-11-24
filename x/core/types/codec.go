package types

import (
	"github.com/cosmos/cosmos-sdk/codec"
	cdctypes "github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/msgservice"
)

func RegisterCodec(cdc *codec.LegacyAmino) {
	cdc.RegisterConcrete(&MsgCreateStats{}, "core/CreateStats", nil)
	cdc.RegisterConcrete(&MsgUpdateStats{}, "core/UpdateStats", nil)
	cdc.RegisterConcrete(&MsgDeleteStats{}, "core/DeleteStats", nil)
	cdc.RegisterConcrete(&MsgIssue{}, "core/Issue", nil)
	// this line is used by starport scaffolding # 2
}

func RegisterInterfaces(registry cdctypes.InterfaceRegistry) {
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgCreateStats{},
		&MsgUpdateStats{},
		&MsgDeleteStats{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgIssue{},
	)
	// this line is used by starport scaffolding # 3

	msgservice.RegisterMsgServiceDesc(registry, &_Msg_serviceDesc)
}

var (
	Amino     = codec.NewLegacyAmino()
	ModuleCdc = codec.NewProtoCodec(cdctypes.NewInterfaceRegistry())
)
