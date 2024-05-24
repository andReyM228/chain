package keeper

import (
	"context"
	"fmt"

	"github.com/andReyM228/one/x/allowed/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

func (k msgServer) CreateAdresses(goCtx context.Context, msg *types.MsgCreateAdresses) (*types.MsgCreateAdressesResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	var adresses = types.Adresses{
		Creator: msg.Creator,
		Adress:  msg.Adress,
	}

	id := k.AppendAdresses(
		ctx,
		adresses,
	)

	return &types.MsgCreateAdressesResponse{
		Id: id,
	}, nil
}

func (k msgServer) UpdateAdresses(goCtx context.Context, msg *types.MsgUpdateAdresses) (*types.MsgUpdateAdressesResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	var adresses = types.Adresses{
		Creator: msg.Creator,
		Id:      msg.Id,
		Adress:  msg.Adress,
	}

	// Checks that the element exists
	val, found := k.GetAdresses(ctx, msg.Id)
	if !found {
		return nil, sdkerrors.Wrap(sdkerrors.ErrKeyNotFound, fmt.Sprintf("key %d doesn't exist", msg.Id))
	}

	// Checks if the msg creator is the same as the current owner
	if msg.Creator != val.Creator {
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "incorrect owner")
	}

	k.SetAdresses(ctx, adresses)

	return &types.MsgUpdateAdressesResponse{}, nil
}

func (k msgServer) DeleteAdresses(goCtx context.Context, msg *types.MsgDeleteAdresses) (*types.MsgDeleteAdressesResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Checks that the element exists
	val, found := k.GetAdresses(ctx, msg.Id)
	if !found {
		return nil, sdkerrors.Wrap(sdkerrors.ErrKeyNotFound, fmt.Sprintf("key %d doesn't exist", msg.Id))
	}

	// Checks if the msg creator is the same as the current owner
	if msg.Creator != val.Creator {
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "incorrect owner")
	}

	k.RemoveAdresses(ctx, msg.Id)

	return &types.MsgDeleteAdressesResponse{}, nil
}
