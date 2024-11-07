package infra

import (
	"github.com/jonasjesusamerico/goexpert-stress-test/internal/domain"
	"github.com/jonasjesusamerico/goexpert-stress-test/internal/service"
)

type Container struct {
	loadTester domain.LoadTester
	generator  domain.Reporter
	cli        CLI
}

func NewContainer() *Container {
	container := &Container{
		loadTester: service.NewTesterService(),
		generator:  service.NewReportService(),
	}
	container.cli = NewCLI(container.loadTester, container.generator)
	return container
}

func (c *Container) GetLoadTester() domain.LoadTester {
	return c.loadTester
}

func (c *Container) GetReporter() domain.Reporter {
	return c.generator
}

func (c *Container) GetCLI() CLI {
	return c.cli
}
