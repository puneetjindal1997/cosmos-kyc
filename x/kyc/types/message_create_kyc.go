package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgCreateKyc = "create_kyc"

var _ sdk.Msg = &MsgCreateKyc{}

func NewMsgCreateKyc(creator string, doc string, description string) *MsgCreateKyc {
	return &MsgCreateKyc{
		Creator:     creator,
		Doc:         doc,
		Description: description,
	}
}

func (msg *MsgCreateKyc) Route() string {
	return RouterKey
}

func (msg *MsgCreateKyc) Type() string {
	return TypeMsgCreateKyc
}

func (msg *MsgCreateKyc) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgCreateKyc) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgCreateKyc) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
