package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgIssue = "issue"

var _ sdk.Msg = &MsgIssue{}

func NewMsgIssue(creator string, amount string, address string, denom string) *MsgIssue {
	return &MsgIssue{
		Creator: creator,
		Amount:  amount,
		Address: address,
		Denom:   denom,
	}
}

func (msg *MsgIssue) Route() string {
	return RouterKey
}

func (msg *MsgIssue) Type() string {
	return TypeMsgIssue
}

func (msg *MsgIssue) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgIssue) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgIssue) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
