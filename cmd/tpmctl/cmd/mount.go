package cmd
/*
import (
	"context"

	"github.com/urfave/cli/v3"
)

func MountVolumeCommand() *cli.Command {
	return &cli.Command{
		Name:      "mount-volume",
		Usage:     "Unlock and mount a LUKS encrypted volume",
		ArgsUsage: "<volume-id>",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:  "key-id",
				Usage: "Recovery key ID to use for unlocking",
			},
			&cli.StringFlag{
				Name:  "mount-point",
				Usage: "Mount point for the volume",
			},
		},
		Action: mountVolumeAction,
	}
}

func mountVolumeAction(ctx context.Context, cmd *cli.Command) error {
	return nil
}

func GetLuksPassphraseCommand() *cli.Command {
	return &cli.Command{
		Name:      "get-luks-passphrase",
		Usage:     "Get LUKS passphrase from recovery key",
		ArgsUsage: "<key-id>",
		Flags: []cli.Flag{
			&cli.BoolFlag{
				Name:  "reveal",
				Usage: "Show passphrase on screen (insecure)",
			},
			&cli.StringFlag{
				Name:  "output",
				Usage: "Write passphrase to file",
			},
		},
		Action: getLuksPassphraseAction,
	}
}

func getLuksPassphraseAction(ctx context.Context, cmd *cli.Command) error {
	return nil
}
*/