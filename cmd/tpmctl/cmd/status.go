package cmd

import (
	"context"
	"fmt"
	"snap-tpmctl/internal/log"
	"snap-tpmctl/internal/validator"

	"github.com/urfave/cli/v3"
)

type statusArgs struct {
	key   int
	other int
}

func newStatusCmd() *cli.Command {
	var args statusArgs

	return &cli.Command{
		Name:    "status",
		Usage:   "Show TPM status",
		Suggest: true,
		Arguments: []cli.Argument{
			&cli.IntArg{
				Name:        "key-id",
				UsageText:   "<key-id>",
				Value:       -1,
				Destination: &args.key,
			},
		},
		Action: func(ctx context.Context, cmd *cli.Command) error {
			return status(ctx, cmd, args)
		},
	}
}

func status(ctx context.Context, _cmd *cli.Command, args statusArgs) error {
	key, err := validator.ValidateKey(args.key)
	if err != nil {
		return cli.Exit(fmt.Errorf("invalid key-id value, %s", err.Error()), 1)
	}

	fmt.Println("This is my status for key", key)

	// slog.Debug("detailed information for troubleshooting")
	// slog.Info("general operational information")
	// slog.Warn("something unexpected but not critical")

	log.Debug(ctx, "detailed information for troubleshooting")
	log.Info(ctx, "general operational information")
	log.Notice(ctx, "something unexpected and critical")
	log.Warning(ctx, "something unexpected but not critical")

	return nil
}
