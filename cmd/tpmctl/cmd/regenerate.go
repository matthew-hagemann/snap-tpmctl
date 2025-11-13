package cmd

import (
	"context"
	"fmt"

	"github.com/urfave/cli/v3"
)

func newRegenerateKeyCmd() *cli.Command {
	return &cli.Command{
		Name:  "regenerate-key",
		Usage: "Regenerate an existing local recovery key",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:  "key-id",
				Usage: "Recovery key ID to use for unlocking",
			},
		},
		Action: regenerateKey,
	}
}

func regenerateKey(ctx context.Context, cmd *cli.Command) error {
	fmt.Println("Regenerated local key with id", cmd.String("key-id"))
	return nil
}

func newRegenerateEnterpriseKeyCmd() *cli.Command {
	return &cli.Command{
		Name:  "regenerate-enterprise-key",
		Usage: "Regenerate an existing enterprise recovery key",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:  "key-id",
				Usage: "Recovery key ID to use for unlocking",
			},
		},
		Action: regenerateEnterpriseKey,
	}
}

func regenerateEnterpriseKey(ctx context.Context, cmd *cli.Command) error {
	fmt.Println("Regenerated enterprise key with id", cmd.String("key-id"))
	return nil
}
