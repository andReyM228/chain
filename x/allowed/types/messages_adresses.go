package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const (
	TypeMsgCreateAdresses = "create_adresses"
	TypeMsgUpdateAdresses = "update_adresses"
	TypeMsgDeleteAdresses = "delete_adresses"
)

var _ sdk.Msg = &MsgCreateAdresses{}

func NewMsgCreateAdresses(creator string, adress string) *MsgCreateAdresses {
	return &MsgCreateAdresses{
		Creator: creator,
		Adress:  adress,
	}
}

func (msg *MsgCreateAdresses) Route() string {
	return RouterKey
}

func (msg *MsgCreateAdresses) Type() string {
	return TypeMsgCreateAdresses
}

func (msg *MsgCreateAdresses) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgCreateAdresses) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgCreateAdresses) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}

var _ sdk.Msg = &MsgUpdateAdresses{}

func NewMsgUpdateAdresses(creator string, id uint64, adress string) *MsgUpdateAdresses {
	return &MsgUpdateAdresses{
		Id:      id,
		Creator: creator,
		Adress:  adress,
	}
}

func (msg *MsgUpdateAdresses) Route() string {
	return RouterKey
}

func (msg *MsgUpdateAdresses) Type() string {
	return TypeMsgUpdateAdresses
}

func (msg *MsgUpdateAdresses) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgUpdateAdresses) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgUpdateAdresses) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}

var _ sdk.Msg = &MsgDeleteAdresses{}

func NewMsgDeleteAdresses(creator string, id uint64) *MsgDeleteAdresses {
	return &MsgDeleteAdresses{
		Id:      id,
		Creator: creator,
	}
}
func (msg *MsgDeleteAdresses) Route() string {
	return RouterKey
}

func (msg *MsgDeleteAdresses) Type() string {
	return TypeMsgDeleteAdresses
}

func (msg *MsgDeleteAdresses) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgDeleteAdresses) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgDeleteAdresses) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
