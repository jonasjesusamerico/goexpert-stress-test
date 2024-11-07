package infra

import (
	"flag"
	"os"
	"testing"

	"github.com/jonasjesusamerico/goexpert-stress-test/internal/domain"
	mock_domain "github.com/jonasjesusamerico/goexpert-stress-test/internal/domain/mock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type CLITestSuite struct {
	suite.Suite
	loadTesterMock      *mock_domain.LoadTester
	reportGeneratorMock *mock_domain.Reporter
	cli                 *cli
	exitCode            int
	usageCalled         bool
}

func (suite *CLITestSuite) SetupTest() {
	suite.loadTesterMock = mock_domain.NewTester(suite.T())
	suite.reportGeneratorMock = mock_domain.NewReporter(suite.T())

	suite.cli = &cli{
		loadTester: suite.loadTesterMock,
		reporter:   suite.reportGeneratorMock,
	}

	suite.exitCode = 0
	suite.usageCalled = false
}

func (suite *CLITestSuite) TestCLI_Execute() {
	testConfig := domain.Config{
		URL:           "http://test.com",
		TotalRequests: 100,
		Concurrency:   10,
	}
	testResult := domain.Result{
		Duration:      0,
		TotalRequests: 100,
		StatusCounts:  map[int]int{200: 100},
	}

	suite.loadTesterMock.On("RunTest", testConfig).Return(testResult)
	suite.reportGeneratorMock.On("Generate", testResult).Return()

	suite.T().Run("Valid Arguments", func(t *testing.T) {
		os.Args = []string{"cmd", "--url=http://test.com", "--requests=100", "--concurrency=10"}
		flag.CommandLine = flag.NewFlagSet(os.Args[0], flag.ExitOnError) // Reset flags

		suite.cli.Start()

		suite.loadTesterMock.AssertExpectations(t)
		suite.reportGeneratorMock.AssertExpectations(t)
		assert.Equal(t, 0, suite.exitCode)
		assert.False(t, suite.usageCalled)
	})

}

func TestCLITestSuite(t *testing.T) {
	suite.Run(t, new(CLITestSuite))
}
