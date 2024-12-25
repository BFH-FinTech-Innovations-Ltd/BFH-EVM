package keeper_test

import (
	"github.com/stretchr/testify/suite"

	"github.com/BFH-FinTech-Innovations-Ltd/BFH-EVM/testutil/integration/evmos/factory"
	"github.com/BFH-FinTech-Innovations-Ltd/BFH-EVM/testutil/integration/evmos/grpc"
	"github.com/BFH-FinTech-Innovations-Ltd/BFH-EVM/testutil/integration/evmos/keyring"
	"github.com/BFH-FinTech-Innovations-Ltd/BFH-EVM/testutil/integration/evmos/network"
)

type KeeperTestSuite struct {
	suite.Suite

	network *network.UnitTestNetwork
	handler grpc.Handler
	keyring keyring.Keyring
	factory factory.TxFactory
}

func (suite *KeeperTestSuite) SetupTest() {
	keys := keyring.New(2)
	nw := network.NewUnitTestNetwork(
		network.WithPreFundedAccounts(keys.GetAllAccAddrs()...),
	)
	gh := grpc.NewIntegrationHandler(nw)
	tf := factory.New(nw, gh)
	suite.network = nw
	suite.factory = tf
	suite.handler = gh
	suite.keyring = keys
}
