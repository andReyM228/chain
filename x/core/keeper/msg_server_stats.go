package keeper

import (
	"context"

	"github.com/andReyM228/one/x/core/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

func (k msgServer) CreateStats(goCtx context.Context, msg *types.MsgCreateStats) (*types.MsgCreateStatsResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Check if the value already exists
	_, isFound := k.GetStats(
		ctx,
		msg.Index,
	)
	if isFound {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "index already set")
	}

	var stats = types.Stats{
		Creator: msg.Creator,
		Index:   msg.Index,
		Date:    msg.Date,
		Stats:   msg.Stats,
	}

	k.SetStats(
		ctx,
		stats,
	)

	return &types.MsgCreateStatsResponse{}, nil
}

func (k msgServer) UpdateStats(goCtx context.Context, msg *types.MsgUpdateStats) (*types.MsgUpdateStatsResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Check if the value exists
	valFound, isFound := k.GetStats(
		ctx,
		msg.Index,
	)
	if !isFound {
		return nil, sdkerrors.Wrap(sdkerrors.ErrKeyNotFound, "index not set")
	}

	// Checks if the the msg creator is the same as the current owner
	if msg.Creator != valFound.Creator {
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "incorrect owner")
	}

	var stats = types.Stats{
		Creator: msg.Creator,
		Index:   msg.Index,
		Date:    msg.Date,
		Stats:   msg.Stats,
	}

	k.SetStats(ctx, stats)

	return &types.MsgUpdateStatsResponse{}, nil
}

func (k msgServer) DeleteStats(goCtx context.Context, msg *types.MsgDeleteStats) (*types.MsgDeleteStatsResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Check if the value exists
	valFound, isFound := k.GetStats(
		ctx,
		msg.Index,
	)
	if !isFound {
		return nil, sdkerrors.Wrap(sdkerrors.ErrKeyNotFound, "index not set")
	}

	// Checks if the the msg creator is the same as the current owner
	if msg.Creator != valFound.Creator {
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "incorrect owner")
	}

	k.RemoveStats(
		ctx,
		msg.Index,
	)

	return &types.MsgDeleteStatsResponse{}, nil
}
