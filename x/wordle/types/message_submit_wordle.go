package types

import (
	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var _ sdk.Msg = &MsgSubmitWordle{}

func NewMsgSubmitWordle(creator string, word string) *MsgSubmitWordle {
	return &MsgSubmitWordle{
		Creator: creator,
		Word:    word,
	}
}

func (msg *MsgSubmitWordle) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
