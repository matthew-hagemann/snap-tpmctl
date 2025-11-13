package cmd

import (
	"context"
	"log/slog"
	"os"

	"github.com/urfave/cli/v3"
)

// TODO: make logging with a more human output. (look at authd for slog)
// TODO: add other commands, similar to status
// TODO: add tests for main.go
// 		Look for table testing
// 		Look at parallel testing
// TODO: look at verbosity with -vvv
// TODO maybe? Look at the offset of 1 in count()
// TODO: rename the project snap-tpmctl

// App is the main application structure.
type App struct {
	root cli.Command
}

// New returns a new App.
func New() App {
	return App{
		root: newRootCmd(),
	}
}

func newRootCmd() cli.Command {
	var verbosity int

	return cli.Command{
		Name:    "snap-tpmctl",
		Usage:   "Ubuntu TPM and FDE management tool",
		Version: "0.1.0",
		Commands: []*cli.Command{
			newStatusCmd(),
			/*newCreateKeyCmd(),
			newCreateEnterpriseKeyCmd(),
			newRegenerateKeyCmd(),
			newRegenerateEnterpriseKeyCmd(),
			newMountVolumeCmd(),
			newGetLuksPassphraseCmd(),*/
		},
		UseShortOptionHandling: true,
		Flags: []cli.Flag{
			&cli.BoolFlag{
				Name:    "verbosity",
				Usage:   "Increase verbosity level",
				Aliases: []string{"f"},
				Config: cli.BoolConfig{
					Count: &verbosity,
				},
			},
			&cli.StringFlag{
				Name:  "tpm-path",
				Usage: "TPM device path",
				Value: "/dev/tpm0",
			},
		},
		Before: func(ctx context.Context, c *cli.Command) error {
			setupLogging(verbosity)
			return nil
		},
	}
}

// Run is the main entry point of the app.
func (a App) Run() error {
	return a.root.Run(context.Background(), os.Args)
}

func setupLogging(level int) {
	l := slog.LevelWarn
	// TODO: this looks weird (the offset by 1)
	switch level {
	case 0:
	case 2:
		l = slog.LevelInfo
	default:
		l = slog.LevelDebug
	}

	h := slog.NewTextHandler(os.Stderr, &slog.HandlerOptions{
		Level: l,
	})
	slog.SetDefault(slog.New(h))
}
