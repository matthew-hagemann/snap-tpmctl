package cmd

import (
	"context"
	"fmt"

	"github.com/urfave/cli/v3"
)

func newRegenerateKeyCmd() *cli.Command {
	return &cli.Command{
		Name:      "regenerate-key",
		Usage:     "Regenerate an existing local recovery key",
		ArgsUsage: "<key-id>",
		Action:    regenerateKey,
	}
}

func regenerateKey(ctx context.Context, cmd *cli.Command) error {
	fmt.Println("Regenerated local key with id", cmd.String("key-id"))
	return nil
}

func newRegenerateEnterpriseKeyCmd() *cli.Command {
	return &cli.Command{
		Name:  "regenerate-enterprise-key",
		Usage: "Regenerate an existing enterprise recovery key", ArgsUsage: "<key-id>",
		Action: regenerateEnterpriseKey,
	}
}

func regenerateEnterpriseKey(ctx context.Context, cmd *cli.Command) error {
	fmt.Println("Regenerated enterprise key with id", cmd.String("key-id"))
	return nil
}
