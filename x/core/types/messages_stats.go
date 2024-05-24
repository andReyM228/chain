package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const (
	TypeMsgCreateStats = "create_stats"
	TypeMsgUpdateStats = "update_stats"
	TypeMsgDeleteStats = "delete_stats"
)

var _ sdk.Msg = &MsgCreateStats{}

func NewMsgCreateStats(
	creator string,
	index string,
	date string,
	stats *DailyStats,

) *MsgCreateStats {
	return &MsgCreateStats{
		Creator: creator,
		Index:   index,
		Date:    date,
		Stats:   stats,
	}
}

func (msg *MsgCreateStats) Route() string {
	return RouterKey
}

func (msg *MsgCreateStats) Type() string {
	return TypeMsgCreateStats
}

func (msg *MsgCreateStats) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgCreateStats) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgCreateStats) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}

var _ sdk.Msg = &MsgUpdateStats{}

func NewMsgUpdateStats(
	creator string,
	index string,
	date string,
	stats *DailyStats,

) *MsgUpdateStats {
	return &MsgUpdateStats{
		Creator: creator,
		Index:   index,
		Date:    date,
		Stats:   stats,
	}
}

func (msg *MsgUpdateStats) Route() string {
	return RouterKey
}

func (msg *MsgUpdateStats) Type() string {
	return TypeMsgUpdateStats
}

func (msg *MsgUpdateStats) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgUpdateStats) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgUpdateStats) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}

var _ sdk.Msg = &MsgDeleteStats{}

func NewMsgDeleteStats(
	creator string,
	index string,

) *MsgDeleteStats {
	return &MsgDeleteStats{
		Creator: creator,
		Index:   index,
	}
}
func (msg *MsgDeleteStats) Route() string {
	return RouterKey
}

func (msg *MsgDeleteStats) Type() string {
	return TypeMsgDeleteStats
}

func (msg *MsgDeleteStats) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgDeleteStats) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgDeleteStats) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
