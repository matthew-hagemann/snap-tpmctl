package cmd

import (
	"context"

	"github.com/urfave/cli/v3"
)

func newMountVolumeCmd() *cli.Command {
	return &cli.Command{
		Name:      "mount-volume",
		Usage:     "Unlock and mount a LUKS encrypted volume",
		ArgsUsage: "<key-id>",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:  "mount-point",
				Usage: "Mount point for the volume",
			},
		},
		Action: mountVolume,
	}
}

func mountVolume(ctx context.Context, cmd *cli.Command) error {
	println("Mount volume in", cmd.String("mount-point"))
	return nil
}

func newGetLuksPassphraseCmd() *cli.Command {
	return &cli.Command{
		Name:      "get-luks-passphrase",
		Usage:     "Get LUKS passphrase from recovery key",
		ArgsUsage: "<key-id>",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:  "file",
				Usage: "Write passphrase to file",
			},
		},
		Action: getLuksPassphrase,
	}
}

func getLuksPassphrase(ctx context.Context, cmd *cli.Command) error {
	println("Get LUKS passphrase for key", cmd.String("key-id"))

	if file := cmd.String("file"); file == "" {
		println("print to stdout")
	} else {
		println("print to file", file)
	}

	return nil
}
