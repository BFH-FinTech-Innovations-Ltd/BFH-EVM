package cosmos_test

import (
	"testing"

	"github.com/stretchr/testify/suite"

	"github.com/BFH-FinTech-Innovations-Ltd/BFH-EVM/app/ante/testutils"
)

type AnteTestSuite struct {
	*testutils.AnteTestSuite
}

func TestAnteTestSuite(t *testing.T) {
	baseSuite := new(testutils.AnteTestSuite)
	baseSuite.WithLondonHardForkEnabled(true)
	baseSuite.WithFeemarketEnabled(true)

	suite.Run(t, &AnteTestSuite{baseSuite})
}
