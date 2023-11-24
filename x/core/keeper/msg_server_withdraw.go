package keeper

import (
	"context"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"one/x/core/types"
)

func (k msgServer) Withdraw(goCtx context.Context, msg *types.MsgWithdraw) (*types.MsgWithdrawResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	amount, ok := sdk.NewIntFromString(msg.Amount)
	if !ok {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "invalid amount")
	}

	address, err := sdk.AccAddressFromBech32(msg.Address)
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "invalid address")
	}

	coin := sdk.NewCoin(msg.Denom, amount)

	coins := sdk.NewCoins(coin)

	err = k.Keeper.bankKeeper.SendCoinsFromAccountToModule(ctx, address, types.ModuleName, coins)
	if err != nil {
		return nil, err
	}

	err = k.Keeper.bankKeeper.BurnCoins(ctx, types.ModuleName, coins)
	if err != nil {
		return nil, err
	}

	err = ctx.EventManager().EmitTypedEvents(msg)
	if err != nil {
		return nil, err
	}

	return &types.MsgWithdrawResponse{}, nil
}
