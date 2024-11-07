package service

import (
	"net/http"
	"sync"
	"time"

	"github.com/jonasjesusamerico/goexpert-stress-test/internal/domain"
)

type testerService struct {
	clientPool sync.Pool
}

func NewTesterService() domain.LoadTester {
	return &testerService{
		clientPool: sync.Pool{
			New: func() interface{} { return &http.Client{} },
		},
	}
}

func (s *testerService) run(wg *sync.WaitGroup, results chan<- int, config domain.Config) {
	defer wg.Done()
	client := s.clientPool.Get().(*http.Client)
	defer s.clientPool.Put(client)

	for j := 0; j < config.TotalRequests/config.Concurrency; j++ {
		resp, err := client.Get(config.URL)
		if err != nil {
			results <- 0 // Erro de conexão
			continue
		}
		results <- resp.StatusCode
		resp.Body.Close()
	}
}

func (s *testerService) RunTest(config domain.Config) domain.Result {
	var wg sync.WaitGroup
	results := make(chan int, config.TotalRequests)
	start := time.Now()

	// Distribui requisições entre workers
	for i := 0; i < config.Concurrency; i++ {
		wg.Add(1)
		go s.run(&wg, results, config)
	}

	// Aguarda finalização dos workers
	wg.Wait()
	close(results)

	return s.results(results, config.TotalRequests, start)
}

func (s *testerService) results(results <-chan int, totalRequests int, start time.Time) domain.Result {
	statusCounts := make(map[int]int)
	for result := range results {
		statusCounts[result]++
	}
	duration := time.Since(start)

	return domain.Result{
		Duration:      duration,
		TotalRequests: totalRequests,
		StatusCounts:  statusCounts,
	}
}
