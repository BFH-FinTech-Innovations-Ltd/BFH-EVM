// Copyright Tharsis Labs Ltd.(Evmos)
// SPDX-License-Identifier:ENCL-1.0(https://github.com/evmos/evmos/blob/main/LICENSE)

package ante_test

import (
	storetypes "cosmossdk.io/store/types"
	"github.com/BFH-FinTech-Innovations-Ltd/BFH-EVM/testutil/integration/evmos/network"
	evmante "github.com/BFH-FinTech-Innovations-Ltd/BFH-EVM/x/evm/ante"
)

func (suite *EvmAnteTestSuite) TestBuildEvmExecutionCtx() {
	network := network.New()

	ctx := evmante.BuildEvmExecutionCtx(network.GetContext())

	suite.Equal(storetypes.GasConfig{}, ctx.KVGasConfig())
	suite.Equal(storetypes.GasConfig{}, ctx.TransientKVGasConfig())
}
