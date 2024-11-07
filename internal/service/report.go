package service

import (
	"fmt"

	"github.com/jonasjesusamerico/goexpert-stress-test/internal/domain"
)

type reportService struct{}

func NewReportService() domain.Reporter {
	return &reportService{}
}

func (s *reportService) Generate(result domain.Result) {
	fmt.Printf("Tempo total gasto: %v\n", result.Duration)
	fmt.Printf("Total de requests: %d\n", result.TotalRequests)
	fmt.Printf("Requests com status 200: %d\n", result.StatusCounts[200])

	for status, count := range result.StatusCounts {
		if status != 200 && status != 0 {
			fmt.Printf("Status %d: %d\n", status, count)
		}
	}
}
