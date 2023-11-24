package keeper

import (
	"context"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/cosmos/cosmos-sdk/types/query"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"one/x/allowed/types"
)

func (k Keeper) AdressesAll(goCtx context.Context, req *types.QueryAllAdressesRequest) (*types.QueryAllAdressesResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var adressess []types.Adresses
	ctx := sdk.UnwrapSDKContext(goCtx)

	store := ctx.KVStore(k.storeKey)
	adressesStore := prefix.NewStore(store, types.KeyPrefix(types.AdressesKey))

	pageRes, err := query.Paginate(adressesStore, req.Pagination, func(key []byte, value []byte) error {
		var adresses types.Adresses
		if err := k.cdc.Unmarshal(value, &adresses); err != nil {
			return err
		}

		adressess = append(adressess, adresses)
		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryAllAdressesResponse{Adresses: adressess, Pagination: pageRes}, nil
}

func (k Keeper) Adresses(goCtx context.Context, req *types.QueryGetAdressesRequest) (*types.QueryGetAdressesResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(goCtx)
	adresses, found := k.GetAdresses(ctx, req.Id)
	if !found {
		return nil, sdkerrors.ErrKeyNotFound
	}

	return &types.QueryGetAdressesResponse{Adresses: adresses}, nil
}

// AddressByAddress returns the adress with the given address
func (k Keeper) AddressByAddress(goCtx context.Context, req *types.QueryGetAddressByAddressRequest) (*types.QueryGetAdressesResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(goCtx)
	address, ok := k.GetAdressesByAdress(ctx, req.Address)
	if !ok {
		return nil, sdkerrors.ErrKeyNotFound
	}

	return &types.QueryGetAdressesResponse{Adresses: address}, nil
}
