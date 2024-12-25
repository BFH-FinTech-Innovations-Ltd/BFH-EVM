// Copyright Tharsis Labs Ltd.(Evmos)
// SPDX-License-Identifier:ENCL-1.0(https://github.com/evmos/evmos/blob/main/LICENSE)
package types

import (
	"math/big"

	sdkmath "cosmossdk.io/math"
)

const (
	// BaseDenom defines the default coin denomination used in Evmos in:
	//
	// - Staking parameters: denomination used as stake in the dPoS chain
	// - Mint parameters: denomination minted due to fee distribution rewards
	// - Governance parameters: denomination used for spam prevention in proposal deposits
	// - Crisis parameters: constant fee denomination used for spam prevention to check broken invariant
	// - EVM parameters: denomination used for running EVM state transitions in Evmos.
	BaseDenom        string = "ubfh"
	BaseDenomTestnet string = "ubfh"

	// BaseDenomUnit defines the base denomination unit for Evmos.
	// 1 bfh = 1x10^{BaseDenomUnit} ubfh
	BaseDenomUnit = 18

	// DisplayDenom defines the denomination displayed to users in client applications.
	DisplayDenom        string = "ubfh"
	DisplayDenomTestnet string = "ubfh"

	// DefaultGasPrice is default gas price for evm transactions
	DefaultGasPrice = 20
)

// PowerReduction defines the default power reduction value for staking
var PowerReduction = sdkmath.NewIntFromBigInt(new(big.Int).Exp(big.NewInt(10), big.NewInt(BaseDenomUnit), nil))
