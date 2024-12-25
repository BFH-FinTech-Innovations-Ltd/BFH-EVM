// Copyright Tharsis Labs Ltd.(Evmos)
// SPDX-License-Identifier:ENCL-1.0(https://github.com/evmos/evmos/blob/main/LICENSE)

package testdata

import (
	contractutils "github.com/BFH-FinTech-Innovations-Ltd/BFH-EVM/contracts/utils"
	evmtypes "github.com/BFH-FinTech-Innovations-Ltd/BFH-EVM/x/evm/types"
)

func LoadERC20AllowanceCaller() (evmtypes.CompiledContract, error) {
	return contractutils.LoadContractFromJSONFile("ERC20AllowanceCaller.json")
}
