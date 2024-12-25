// Copyright Tharsis Labs Ltd.(Evmos)
// SPDX-License-Identifier:ENCL-1.0(https://github.com/evmos/evmos/blob/main/LICENSE)

package keeper

import (
	"cosmossdk.io/store/prefix"
	"github.com/BFH-FinTech-Innovations-Ltd/BFH-EVM/x/evm/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/ethereum/go-ethereum/common"
)

// IsContract determines if the given address is a smart contract.
func (k *Keeper) IsContract(ctx sdk.Context, addr common.Address) bool {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefixCodeHash)
	return store.Has(addr.Bytes())
}
