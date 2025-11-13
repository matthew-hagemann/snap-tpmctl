package cmd

import (
	"context"
	"fmt"
	"log/slog"

	"github.com/urfave/cli/v3"
)

func newStatusCmd() *cli.Command {
	return &cli.Command{
		Name:      "status",
		Usage:     "Show TPM status",
		ArgsUsage: "<key-id>",
		Action:    status,
	}
}

func status(ctx context.Context, cmd *cli.Command) error {
	fmt.Println("This is my status")
	slog.Debug("this is my debug log")
	return nil
}
