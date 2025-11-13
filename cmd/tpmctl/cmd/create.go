package cmd

import (
	"context"
	"fmt"

	"github.com/urfave/cli/v3"
)

func newCreateKeyCmd() *cli.Command {
	return &cli.Command{
		Name:   "create-key",
		Usage:  "Create a new local recovery key",
		Action: createKey,
	}
}

func createKey(ctx context.Context, cmd *cli.Command) error {
	fmt.Println("Created local key")
	return nil
}

func newCreateEnterpriseKeyCmd() *cli.Command {
	return &cli.Command{
		Name:   "create-enterprise-key",
		Usage:  "Create a new enterprise recovery key for Landscape",
		Action: createEnterpriseKey,
	}
}

func createEnterpriseKey(ctx context.Context, cmd *cli.Command) error {
	fmt.Println("Created enterprise key")
	return nil
}
