package domain

import "time"

type Config struct {
	URL           string
	TotalRequests int
	Concurrency   int
}

type Result struct {
	Duration      time.Duration
	TotalRequests int
	StatusCounts  map[int]int
}
