package poa

import (
	"encoding/json"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/tendermint/tendermint/crypto"
)

// MsgCreateValidator defines a CreateValidator message
type MsgCreateValidator struct {
	ValidatorAddr sdk.ValAddress
	PubKey        crypto.PubKey
}

// NewMsgCreateValidator is a constructor function for MsgSetName
func NewMsgCreateValidator(valAddr sdk.ValAddress, pubKey crypto.PubKey) MsgCreateValidator {
	return MsgCreateValidator{
		ValidatorAddr: valAddr,
		PubKey:        pubKey,
	}
}

// Route Implements Msg.
func (msg MsgCreateValidator) Route() string { return "poa" }

// Type Implements Msg.
func (msg MsgCreateValidator) Type() string { return "create_validator" }

// ValidateBasic Implements Msg.
func (msg MsgCreateValidator) ValidateBasic() sdk.Error {
	return nil
}

// GetSignBytes Implements Msg.
func (msg MsgCreateValidator) GetSignBytes() []byte {
	b, err := json.Marshal(msg)
	if err != nil {
		panic(err)
	}
	return sdk.MustSortJSON(b)
}

// GetSigners Implements Msg.
func (msg MsgCreateValidator) GetSigners() []sdk.AccAddress {
	// The validator to be added needs to sign
	return []sdk.AccAddress{sdk.AccAddress(msg.ValidatorAddr)}
}
