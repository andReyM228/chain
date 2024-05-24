package keeper_test

import (
	"testing"

	testkeeper "github.com/andReyM228/one/testutil/keeper"
	"github.com/andReyM228/one/x/one/types"
	"github.com/stretchr/testify/require"
)

func TestGetParams(t *testing.T) {
	k, ctx := testkeeper.OneKeeper(t)
	params := types.DefaultParams()

	k.SetParams(ctx, params)

	require.EqualValues(t, params, k.GetParams(ctx))
}
