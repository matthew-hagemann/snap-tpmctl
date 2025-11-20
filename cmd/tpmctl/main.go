package main

import (
	"log/slog"
	"os"
	"snap-tpmctl/cmd/tpmctl/cmd"
)

type app interface {
	Run() error
}

func main() {
	a := cmd.New(os.Args)
	os.Exit(run(a))
}

func run(a app) int {
	if err := a.Run(); err != nil {
		slog.Error(err.Error())
		return 1
	}

	return 0
}
