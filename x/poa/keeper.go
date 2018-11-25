package poa

import (
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	abci "github.com/tendermint/tendermint/abci/types"
	"github.com/tendermint/tendermint/crypto"
	tmtypes "github.com/tendermint/tendermint/types"
)

// Keeper exposes the data store
type Keeper struct {
	storeKey sdk.StoreKey

	cdc *codec.Codec
}

// ValidatorValue declares the main validator struct
type ValidatorValue struct {
	Address sdk.ValAddress
	PubKey  crypto.PubKey
}

// ABCIValidatorUpdate returns the abci.ValidatorUpdate for handling updates
func (v ValidatorValue) ABCIValidatorUpdate() abci.ValidatorUpdate {
	return abci.ValidatorUpdate{
		PubKey: tmtypes.TM2PB.PubKey(v.PubKey),
		Power:  10, //NOTE: Set arbitrary value - can be changed for setting a weight on each validator
	}
}

// CreateValidator creates a new validator (the address is the only data handled for now)
func (k Keeper) CreateValidator(ctx sdk.Context, valAddr sdk.ValAddress, pubKey crypto.PubKey) {
	store := ctx.KVStore(k.storeKey)
	//NOTE: second argument is unused, can be used later for a better handling with voting power and moniker...
	val := &ValidatorValue{valAddr, pubKey}
	store.Set([]byte(valAddr), k.cdc.MustMarshalBinary(*val))
}

// ValidatorExists returns true if the validator has been registered already
func (k Keeper) ValidatorExists(ctx sdk.Context, validatorAddr sdk.ValAddress) bool {
	store := ctx.KVStore(k.storeKey)
	val := store.Get([]byte(validatorAddr))
	return (val != nil)
}

// GetAllValidators returns all validators of the datastore
func (k Keeper) GetAllValidators(ctx sdk.Context) []ValidatorValue {
	store := ctx.KVStore(k.storeKey)
	it := store.Iterator(nil, nil)
	validators := []ValidatorValue{}
	for ; it.Valid(); it.Next() {
		val := &ValidatorValue{}
		if err := k.cdc.UnmarshalBinary(it.Value(), val); err != nil {
			continue
		}
		validators = append(validators, *val)
	}
	return validators
}
