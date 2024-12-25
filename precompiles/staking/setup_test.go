package staking_test

import (
	"testing"

	"github.com/BFH-FinTech-Innovations-Ltd/BFH-EVM/precompiles/staking"
	"github.com/BFH-FinTech-Innovations-Ltd/BFH-EVM/testutil/integration/evmos/factory"
	"github.com/BFH-FinTech-Innovations-Ltd/BFH-EVM/testutil/integration/evmos/grpc"
	testkeyring "github.com/BFH-FinTech-Innovations-Ltd/BFH-EVM/testutil/integration/evmos/keyring"
	"github.com/BFH-FinTech-Innovations-Ltd/BFH-EVM/testutil/integration/evmos/network"
	"github.com/stretchr/testify/suite"
)

type PrecompileTestSuite struct {
	suite.Suite

	network     *network.UnitTestNetwork
	factory     factory.TxFactory
	grpcHandler grpc.Handler
	keyring     testkeyring.Keyring

	bondDenom  string
	precompile *staking.Precompile
}

func TestPrecompileUnitTestSuite(t *testing.T) {
	suite.Run(t, new(PrecompileTestSuite))
}

func (s *PrecompileTestSuite) SetupTest() {
	keyring := testkeyring.New(2)
	nw := network.NewUnitTestNetwork(
		network.WithPreFundedAccounts(keyring.GetAllAccAddrs()...),
	)
	grpcHandler := grpc.NewIntegrationHandler(nw)
	txFactory := factory.New(nw, grpcHandler)

	ctx := nw.GetContext()
	sk := nw.App.StakingKeeper
	bondDenom, err := sk.BondDenom(ctx)
	if err != nil {
		panic(err)
	}

	s.bondDenom = bondDenom
	s.factory = txFactory
	s.grpcHandler = grpcHandler
	s.keyring = keyring
	s.network = nw

	if s.precompile, err = staking.NewPrecompile(
		s.network.App.StakingKeeper,
		s.network.App.AuthzKeeper,
	); err != nil {
		panic(err)
	}
}
