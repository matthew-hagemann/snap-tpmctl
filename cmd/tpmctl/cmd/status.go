package cmd

import (
	"context"
	"fmt"
	"snap-tpmctl/internal/log"

	"github.com/urfave/cli/v3"
)

func newStatusCmd() *cli.Command {
	var recoveryKey string

	return &cli.Command{
		Name:    "status",
		Usage:   "Show TPM status",
		Suggest: true,
		Arguments: []cli.Argument{
			&cli.StringArg{
				// check
				Name:        "key-id",
				UsageText:   "<key-id>",
				Destination: &recoveryKey,
			},
		},
		Action: func(ctx context.Context, cmd *cli.Command) error {
			if err := isValidRecoveryKey(recoveryKey); err != nil {
				return err
			}

			return status(ctx, recoveryKey)
		},
	}
}

func status(ctx context.Context, recoveryKey string) error {
	fmt.Println("This is my status for key", recoveryKey)

	// slog.Debug("detailed information for troubleshooting")
	// slog.Info("general operational information")
	// slog.Warn("something unexpected but not critical")

	log.Debug(ctx, "detailed information for troubleshooting")
	log.Info(ctx, "general operational information")
	log.Notice(ctx, "something unexpected and critical")
	log.Warning(ctx, "something unexpected but not critical")

	return nil
}

func isValidRecoveryKey(k string) error {
	// TODO: regexp validation
	return nil
}
