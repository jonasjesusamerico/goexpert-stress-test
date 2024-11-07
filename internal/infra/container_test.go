package infra

import (
	"testing"

	"github.com/jonasjesusamerico/goexpert-stress-test/internal/domain"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type ContainerTestSuite struct {
	suite.Suite
	container *Container
}

func (suite *ContainerTestSuite) SetupTest() {
	suite.container = NewContainer()
}

func (suite *ContainerTestSuite) TestNewContainer() {
	assert.NotNil(suite.T(), suite.container)
	assert.NotNil(suite.T(), suite.container.loadTester)
	assert.NotNil(suite.T(), suite.container.generator)
	assert.NotNil(suite.T(), suite.container.cli)
}

func (suite *ContainerTestSuite) TestGetLoadTester() {
	loadTester := suite.container.GetLoadTester()
	assert.NotNil(suite.T(), loadTester)
	assert.Implements(suite.T(), (*domain.LoadTester)(nil), loadTester)
}

func (suite *ContainerTestSuite) TestGetReportGenerator() {
	reportGenerator := suite.container.GetReporter()
	assert.NotNil(suite.T(), reportGenerator)
	assert.Implements(suite.T(), (*domain.Reporter)(nil), reportGenerator)
}

func (suite *ContainerTestSuite) TestGetCLI() {
	cli := suite.container.GetCLI()
	assert.NotNil(suite.T(), cli)
}

func TestContainerTestSuite(t *testing.T) {
	suite.Run(t, new(ContainerTestSuite))
}
