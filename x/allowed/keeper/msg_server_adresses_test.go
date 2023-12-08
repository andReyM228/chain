package keeper_test

import (
	"testing"

	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/stretchr/testify/require"

	"github.com/andReyM228/one/x/allowed/types"
)

func TestAdressesMsgServerCreate(t *testing.T) {
	srv, ctx := setupMsgServer(t)
	creator := "A"
	for i := 0; i < 5; i++ {
		resp, err := srv.CreateAdresses(ctx, &types.MsgCreateAdresses{Creator: creator})
		require.NoError(t, err)
		require.Equal(t, i, int(resp.Id))
	}
}

func TestAdressesMsgServerUpdate(t *testing.T) {
	creator := "A"

	tests := []struct {
		desc    string
		request *types.MsgUpdateAdresses
		err     error
	}{
		{
			desc:    "Completed",
			request: &types.MsgUpdateAdresses{Creator: creator},
		},
		{
			desc:    "Unauthorized",
			request: &types.MsgUpdateAdresses{Creator: "B"},
			err:     sdkerrors.ErrUnauthorized,
		},
		{
			desc:    "Unauthorized",
			request: &types.MsgUpdateAdresses{Creator: creator, Id: 10},
			err:     sdkerrors.ErrKeyNotFound,
		},
	}
	for _, tc := range tests {
		t.Run(tc.desc, func(t *testing.T) {
			srv, ctx := setupMsgServer(t)
			_, err := srv.CreateAdresses(ctx, &types.MsgCreateAdresses{Creator: creator})
			require.NoError(t, err)

			_, err = srv.UpdateAdresses(ctx, tc.request)
			if tc.err != nil {
				require.ErrorIs(t, err, tc.err)
			} else {
				require.NoError(t, err)
			}
		})
	}
}

func TestAdressesMsgServerDelete(t *testing.T) {
	creator := "A"

	tests := []struct {
		desc    string
		request *types.MsgDeleteAdresses
		err     error
	}{
		{
			desc:    "Completed",
			request: &types.MsgDeleteAdresses{Creator: creator},
		},
		{
			desc:    "Unauthorized",
			request: &types.MsgDeleteAdresses{Creator: "B"},
			err:     sdkerrors.ErrUnauthorized,
		},
		{
			desc:    "KeyNotFound",
			request: &types.MsgDeleteAdresses{Creator: creator, Id: 10},
			err:     sdkerrors.ErrKeyNotFound,
		},
	}
	for _, tc := range tests {
		t.Run(tc.desc, func(t *testing.T) {
			srv, ctx := setupMsgServer(t)

			_, err := srv.CreateAdresses(ctx, &types.MsgCreateAdresses{Creator: creator})
			require.NoError(t, err)
			_, err = srv.DeleteAdresses(ctx, tc.request)
			if tc.err != nil {
				require.ErrorIs(t, err, tc.err)
			} else {
				require.NoError(t, err)
			}
		})
	}
}
