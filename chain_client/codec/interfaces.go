package client_codec

import (
	"github.com/cosmos/cosmos-sdk/client"
	cdctypes "github.com/cosmos/cosmos-sdk/codec/types"
)

type Codec interface {
	GetTxConfig() client.TxConfig
	GetInterfaceRegistry() cdctypes.InterfaceRegistry
	UnpackAny(any *cdctypes.Any, receiver interface{}) error
}
