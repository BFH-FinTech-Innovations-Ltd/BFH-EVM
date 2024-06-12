// Copyright Tharsis Labs Ltd.(Evmos)
// SPDX-License-Identifier:ENCL-1.0(https://github.com/evmos/evmos/blob/main/LICENSE)

package tokenfactory

import (
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"math/big"

	commonerr "github.com/evmos/evmos/v18/precompiles/common"

	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"
	erc20types "github.com/evmos/evmos/v18/x/erc20/types"
)

// EventMint is emitted when a new ERC20 token is minted.
type EventMint struct {
	To    common.Address
	Value *big.Int
}

// EventERC20Created is emitted when a new ERC20 token is created.
type EventERC20Created struct {
	Creator       common.Address
	Name          string
	Symbol        string
	Decimals      uint8
	InitialSupply *big.Int
	TokenAddress  common.Address
}

// ParseCreateErc20Args parses the arguments for creating an ERC20 token factory token.
func ParseCreateErc20Args(args []interface{}) (string, string, uint8, *big.Int, error) {
	if len(args) != 4 {
		return "", "", 0, nil, fmt.Errorf(commonerr.ErrInvalidNumberOfArgs, 4, len(args))
	}

	return parseERC20Args(args)
}

// ParseCreate2Erc20Args parses the arguments for creating an ERC20 token factory token with create2.
func ParseCreate2Erc20Args(args []interface{}) (string, string, uint8, *big.Int, [32]byte, error) {
	if len(args) != 5 {
		return "", "", 0, nil, [32]byte{}, fmt.Errorf(commonerr.ErrInvalidNumberOfArgs, 5, len(args))
	}

	name, symbol, decimals, initialSupply, err := parseERC20Args(args)
	if err != nil {
		return "", "", 0, nil, [32]byte{}, err
	}

	salt, ok := args[4].([32]byte)
	if !ok {
		return "", "", 0, nil, [32]byte{}, fmt.Errorf(ErrInvalidSaltArgumentType, args[4])
	}

	return name, symbol, decimals, initialSupply, salt, nil
}

// parseERC20Args parses the arguments for creating an ERC20 token factory token.
func parseERC20Args(args []interface{}) (string, string, uint8, *big.Int, error) {
	name, ok := args[0].(string)
	if !ok {
		return "", "", 0, nil, fmt.Errorf(ErrInvalidNameArgumentType, args[0])
	}

	symbol, ok := args[1].(string)
	if !ok {
		return "", "", 0, nil, fmt.Errorf(ErrInvalidSymbolArgumentType, args[1])
	}

	decimals, ok := args[2].(uint8)
	if !ok {
		return "", "", 0, nil, fmt.Errorf(ErrInvalidDecimalArgumentType, args[2])
	}

	initialSupply, ok := args[3].(*big.Int)
	if !ok {
		return "", "", 0, nil, fmt.Errorf(ErrInvalidInitialSupplyArgumentType, args[3])
	}

	return name, symbol, decimals, initialSupply, nil
}

// FIXME: make the one from Evmos ERC-20 precompile public
func NewDenomMetaData(contract, baseDenom, name, symbol string, decimals uint8) banktypes.Metadata {
	// create a bank denom metadata based on the ERC20 token ABI details
	// metadata name is should always be the contract since it's the key
	// to the bank store
	metadata := banktypes.Metadata{
		Description: erc20types.CreateDenomDescription(contract),
		Base:        baseDenom,
		// NOTE: Denom units MUST be increasing
		DenomUnits: []*banktypes.DenomUnit{
			{
				Denom:    baseDenom,
				Exponent: 0,
			},
		},
		Name:    name,
		Symbol:  symbol,
		Display: baseDenom,
	}

	// only append metadata if decimals > 0, otherwise validation fails
	if decimals > 0 {
		nameSanitized := erc20types.SanitizeERC20Name(name)
		metadata.DenomUnits = append(
			metadata.DenomUnits,
			&banktypes.DenomUnit{
				Denom:    nameSanitized,
				Exponent: uint32(decimals),
			},
		)
		metadata.Display = nameSanitized
	}

	return metadata
}