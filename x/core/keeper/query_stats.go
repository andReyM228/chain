package keeper

import (
	"context"

	"github.com/andReyM228/one/x/core/types"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/query"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) StatsAll(goCtx context.Context, req *types.QueryAllStatsRequest) (*types.QueryAllStatsResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var statss []types.Stats
	ctx := sdk.UnwrapSDKContext(goCtx)

	store := ctx.KVStore(k.storeKey)
	statsStore := prefix.NewStore(store, types.KeyPrefix(types.StatsKeyPrefix))

	pageRes, err := query.Paginate(statsStore, req.Pagination, func(key []byte, value []byte) error {
		var stats types.Stats
		if err := k.cdc.Unmarshal(value, &stats); err != nil {
			return err
		}

		statss = append(statss, stats)
		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryAllStatsResponse{Stats: statss, Pagination: pageRes}, nil
}

func (k Keeper) Stats(goCtx context.Context, req *types.QueryGetStatsRequest) (*types.QueryGetStatsResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}
	ctx := sdk.UnwrapSDKContext(goCtx)

	val, found := k.GetStats(
		ctx,
		req.Index,
	)
	if !found {
		return nil, status.Error(codes.NotFound, "not found")
	}

	return &types.QueryGetStatsResponse{Stats: val}, nil
}
