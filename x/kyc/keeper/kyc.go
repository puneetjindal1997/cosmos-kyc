package keeper

import (
	"encoding/binary"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"

	"kyc/x/kyc/types"
)

func (k Keeper) GetKycCount(ctx sdk.Context) uint64 {
	// Get the store using storeKey (which is "blog") and PostCountKey (which is "Post/count/")
	store := prefix.NewStore(ctx.KVStore(k.storeKey), []byte(types.KycCountKey))

	// Convert the PostCountKey to bytes
	byteKey := []byte(types.KycCountKey)

	// Get the value of the count
	bz := store.Get(byteKey)

	// Return zero if the count value is not found (for example, it's the first post)
	if bz == nil {
		return 0
	}

	// Convert the count into a uint64
	return binary.BigEndian.Uint64(bz)
}

func (k Keeper) SetKycCount(ctx sdk.Context, count uint64) {
	// Get the store using storeKey (which is "blog") and PostCountKey (which is "Post/count/")
	store := prefix.NewStore(ctx.KVStore(k.storeKey), []byte(types.KycCountKey))

	// Convert the PostCountKey to bytes
	byteKey := []byte(types.KycCountKey)

	// Convert count from uint64 to string and get bytes
	bz := make([]byte, 8)
	binary.BigEndian.PutUint64(bz, count)

	// Set the value of Post/count/ to count
	store.Set(byteKey, bz)
}

func (k Keeper) AppendKyc(ctx sdk.Context, kyc types.Kyc) uint64 {
	// Get the current number of kycs in the store
	count := k.GetKycCount(ctx)

	// Assign an ID to the kyc based on the number of kycs in the store
	kyc.Id = count

	// Get the store
	store := prefix.NewStore(ctx.KVStore(k.storeKey), []byte(types.KycKey))

	// Convert the kyc ID into bytes
	byteKey := make([]byte, 8)
	binary.BigEndian.PutUint64(byteKey, kyc.Id)

	// Marshal the kyc into bytes
	appendedValue := k.cdc.MustMarshal(&kyc)

	// Insert the kyc bytes using kyc ID as a key
	store.Set(byteKey, appendedValue)

	// Update the kyc count
	k.SetKycCount(ctx, count+1)
	return count
}
