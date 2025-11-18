package cmd

import (
	"context"
	"fmt"
	"snap-tpmctl/internal/log"

	"github.com/urfave/cli/v3"
)

func newStatusCmd() *cli.Command {
	return &cli.Command{
		Name:    "status",
		Usage:   "Show TPM status",
		Suggest: true,
		Arguments: []cli.Argument{
			&cli.IntArg{
				Name:      "key-id",
				UsageText: "<key-id>",
				Value:     -1,
			},
		},
		Action: status,
	}
}

func status(ctx context.Context, cmd *cli.Command) error {
	// TODO: add validator for key-id

	// TODO: is it safe calling arguments like this?
	if cmd.IntArg("key-id") < 0 {
		return cli.Exit("Missing key-id argument", 1)
	}

	fmt.Println("This is my status for key", cmd.IntArg("key-id"))

	// slog.Debug("detailed information for troubleshooting")
	// slog.Info("general operational information")
	// slog.Warn("something unexpected but not critical")

	log.Debug(ctx, "detailed information for troubleshooting")
	log.Info(ctx, "general operational information")
	log.Notice(ctx, "something unexpected and critical")
	log.Warning(ctx, "something unexpected but not critical")

	return nil
}
