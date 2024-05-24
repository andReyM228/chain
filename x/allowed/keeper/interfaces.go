package keeper

import (
	"context"

	allowedtypes "github.com/andReyM228/one/x/allowed/types"
)

type AllowedKeeper interface {
	AddressByAddress(goCtx context.Context, req *allowedtypes.QueryGetAddressByAddressRequest) (*allowedtypes.QueryGetAdressesResponse, error)
}
