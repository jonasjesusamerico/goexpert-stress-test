package main

import (
	"github.com/jonasjesusamerico/goexpert-stress-test/internal/infra"
)

func main() {
	container := infra.NewContainer()

	cli := container.GetCLI()

	cli.Start()
}
