package keeper_test

import (
	"context"
	"testing"

	keepertest "github.com/andReyM228/one/testutil/keeper"
	"github.com/andReyM228/one/x/allowed/keeper"
	"github.com/andReyM228/one/x/allowed/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
)

func setupMsgServer(t testing.TB) (types.MsgServer, context.Context) {
	k, ctx := keepertest.AllowedKeeper(t)
	return keeper.NewMsgServerImpl(*k), sdk.WrapSDKContext(ctx)
}

func TestMsgServer(t *testing.T) {
	ms, ctx := setupMsgServer(t)
	require.NotNil(t, ms)
	require.NotNil(t, ctx)
}
