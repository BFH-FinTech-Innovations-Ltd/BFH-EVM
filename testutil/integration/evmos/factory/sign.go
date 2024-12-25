// Copyright Tharsis Labs Ltd.(Evmos)
// SPDX-License-Identifier:ENCL-1.0(https://github.com/evmos/evmos/blob/main/LICENSE)
package factory

import (
	errorsmod "cosmossdk.io/errors"
	"github.com/BFH-FinTech-Innovations-Ltd/BFH-EVM/testutil/tx"
	evmtypes "github.com/BFH-FinTech-Innovations-Ltd/BFH-EVM/x/evm/types"
	cryptotypes "github.com/cosmos/cosmos-sdk/crypto/types"
	gethtypes "github.com/ethereum/go-ethereum/core/types"
)

// SignMsgEthereumTx signs a MsgEthereumTx with the provided private key and chainID.
func (tf *IntegrationTxFactory) SignMsgEthereumTx(privKey cryptotypes.PrivKey, msgEthereumTx evmtypes.MsgEthereumTx) (evmtypes.MsgEthereumTx, error) {
	ethChainID := tf.network.GetEIP155ChainID()
	signer := gethtypes.LatestSignerForChainID(ethChainID)
	err := msgEthereumTx.Sign(signer, tx.NewSigner(privKey))
	if err != nil {
		return evmtypes.MsgEthereumTx{}, errorsmod.Wrap(err, "failed to sign transaction")
	}
	return msgEthereumTx, nil
}
