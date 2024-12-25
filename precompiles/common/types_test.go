package common_test

import (
	"testing"

	"cosmossdk.io/math"
	"github.com/BFH-FinTech-Innovations-Ltd/BFH-EVM/precompiles/common"
	evmostypes "github.com/BFH-FinTech-Innovations-Ltd/BFH-EVM/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
)

var largeAmt, _ = math.NewIntFromString("1000000000000000000000000000000000000000")

func TestNewCoinsResponse(t *testing.T) {
	testCases := []struct {
		amount math.Int
	}{
		{amount: math.NewInt(1)},
		{amount: largeAmt},
	}

	for _, tc := range testCases {
		coin := sdk.NewCoin(evmostypes.BaseDenom, tc.amount)
		coins := sdk.NewCoins(coin)
		res := common.NewCoinsResponse(coins)
		require.Equal(t, 1, len(res))
		require.Equal(t, tc.amount.BigInt(), res[0].Amount)
	}
}

func TestNewDecCoinsResponse(t *testing.T) {
	testCases := []struct {
		amount math.Int
	}{
		{amount: math.NewInt(1)},
		{amount: largeAmt},
	}

	for _, tc := range testCases {
		coin := sdk.NewDecCoin(evmostypes.BaseDenom, tc.amount)
		coins := sdk.NewDecCoins(coin)
		res := common.NewDecCoinsResponse(coins)
		require.Equal(t, 1, len(res))
		require.Equal(t, tc.amount.BigInt(), res[0].Amount)
	}
}
