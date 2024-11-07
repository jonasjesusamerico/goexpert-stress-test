package infra

import (
	"flag"
	"fmt"
	"os"

	"github.com/jonasjesusamerico/goexpert-stress-test/internal/domain"
)

type CLI interface {
	Start()
}

type cli struct {
	loadTester domain.LoadTester
	reporter   domain.Reporter
}

func NewCLI(loadTester domain.LoadTester, reporter domain.Reporter) CLI {
	return &cli{
		loadTester: loadTester,
		reporter:   reporter,
	}
}

func (c *cli) Start() {
	var config domain.Config

	flag.StringVar(&config.URL, "url", "", "URL do serviço a ser testado")
	flag.IntVar(&config.TotalRequests, "requests", 0, "Número total de requests")
	flag.IntVar(&config.Concurrency, "concurrency", 1, "Número de chamadas simultâneas")
	flag.Parse()

	if err := validateConfig(config); err != nil {
		fmt.Println("Erro de configuração:", err)
		flag.Usage()
		os.Exit(1)
	}

	result := c.loadTester.RunTest(config)
	c.reporter.Generate(result)
}

// validateConfig verifica se os valores de configuração são válidos
func validateConfig(config domain.Config) error {
	if config.URL == "" {
		return fmt.Errorf("a URL do serviço é obrigatória")
	}
	if config.TotalRequests <= 0 {
		return fmt.Errorf("o número total de requests deve ser maior que zero")
	}
	if config.Concurrency <= 0 {
		return fmt.Errorf("o número de chamadas simultâneas deve ser maior que zero")
	}
	return nil
}
