package keeper

import (
	"context"
	sdkioerrors "cosmossdk.io/errors"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"one/x/core/types"
)

func (k msgServer) Issue(goCtx context.Context, msg *types.MsgIssue) (*types.MsgIssueResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	amount, ok := sdk.NewIntFromString(msg.Amount)
	if !ok {
		return nil, sdkioerrors.Wrap(sdkerrors.ErrInvalidRequest, "invalid amount")
	}

	address, err := sdk.AccAddressFromBech32(msg.Address)
	if err != nil {
		return nil, sdkioerrors.Wrap(sdkerrors.ErrInvalidRequest, "invalid address")
	}

	// Check if the address is allowed
	_, ok = k.Keeper.allowedKeeper.GetAdressesByAdress(ctx, msg.Creator)
	if !ok {
		return nil, sdkioerrors.Wrap(sdkerrors.ErrUnauthorized, "address is not allowed")
	}

	coin := sdk.NewCoin(msg.Denom, amount)

	coins := sdk.NewCoins(coin)

	err = k.Keeper.bankKeeper.MintCoins(ctx, types.ModuleName, coins)
	if err != nil {
		return nil, err
	}

	err = k.Keeper.bankKeeper.SendCoinsFromModuleToAccount(ctx, types.ModuleName, address, coins)
	if err != nil {
		return nil, err
	}

	k.SaveStatsIssue(ctx, coins)

	err = ctx.EventManager().EmitTypedEvents(msg)
	if err != nil {
		return nil, err
	}

	return &types.MsgIssueResponse{}, nil
}
