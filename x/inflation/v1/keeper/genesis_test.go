package keeper_test

import (
	"testing"

	"cosmossdk.io/math"
	"github.com/BFH-FinTech-Innovations-Ltd/BFH-EVM/testutil/integration/evmos/network"
	"github.com/BFH-FinTech-Innovations-Ltd/BFH-EVM/x/inflation/v1/types"
	"github.com/stretchr/testify/require"
)

func TestInitGenesis(t *testing.T) {
	nw := network.NewUnitTestNetwork()
	ctx := nw.GetContext()
	// check calculated epochMintProvision at genesis
	epochMintProvision := nw.App.InflationKeeper.GetEpochMintProvision(ctx)
	expMintProvision := math.LegacyMustNewDecFromStr("282534246575342465753425").Quo(math.LegacyNewDec(types.ReductionFactor))
	require.Equal(t, expMintProvision, epochMintProvision)
}
