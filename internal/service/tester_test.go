package service

import (
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/jonasjesusamerico/goexpert-stress-test/internal/domain"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type MockHTTPClient struct {
	DoFunc func(req *http.Request) (*http.Response, error)
}

func (m *MockHTTPClient) Do(req *http.Request) (*http.Response, error) {
	return m.DoFunc(req)
}

type LoadTesterSuite struct {
	suite.Suite
	mockClient *MockHTTPClient
	loadTester domain.LoadTester
	mockServer *httptest.Server
}

func (suite *LoadTesterSuite) SetupTest() {
	suite.mockServer = httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	}))

	suite.mockClient = &MockHTTPClient{
		DoFunc: func(req *http.Request) (*http.Response, error) {
			if req.URL.String() != suite.mockServer.URL {
				return nil, errors.New("URL inesperada")
			}
			return &http.Response{
				StatusCode: http.StatusOK,
				Body:       nil,
			}, nil
		},
	}

	suite.loadTester = NewTesterService()
}

func (suite *LoadTesterSuite) TearDownTest() {
	suite.mockServer.Close()
}

func (suite *LoadTesterSuite) TestRunLoadTest() {
	totalRequests := 100
	concurrency := 5

	suite.mockClient.DoFunc = func(req *http.Request) (*http.Response, error) {
		return &http.Response{
			StatusCode: http.StatusOK,
			Body:       nil,
		}, nil
	}

	config := domain.Config{
		URL:           suite.mockServer.URL,
		TotalRequests: totalRequests,
		Concurrency:   concurrency,
	}
	result := suite.loadTester.RunTest(config)

	assert.Equal(suite.T(), totalRequests, result.TotalRequests)
	assert.True(suite.T(), result.Duration > 0)

	// Calcula o número esperado de status http.StatusOK
	expectedStatusCount := totalRequests // Cada request deve retornar http.StatusOK
	assert.Equal(suite.T(), expectedStatusCount, result.StatusCounts[http.StatusOK])
}

func TestLoadTesterSuite(t *testing.T) {
	suite.Run(t, new(LoadTesterSuite))
}
