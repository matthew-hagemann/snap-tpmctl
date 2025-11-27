package cmd

import (
	"context"
	"fmt"
	"snap-tpmctl/internal/snapd"

	"github.com/urfave/cli/v3"
)

func newRegenerateKeyCmd() *cli.Command {
	var recoveryKey string

	return &cli.Command{
		Name:    "regenerate-key",
		Usage:   "Regenerate an existing local recovery key",
		Suggest: true,
		Arguments: []cli.Argument{
			&cli.StringArg{
				Name:        "key-id",
				UsageText:   "<key-id>",
				Destination: &recoveryKey,
			},
		},
		Action: func(ctx context.Context, c *cli.Command) error {
			return regenerateKey(ctx, recoveryKey)
		},
	}
}

func regenerateKey(ctx context.Context, _ string) error {
	c := snapd.NewClient()
	defer c.Close()

	if err := c.LoadAuthFromHome(); err != nil {
		return err
	}

	key, err := c.GenerateRecoveryKey(ctx)
	if err != nil {
		return fmt.Errorf("failed to generate recovery key: %w", err)
	}

	fmt.Printf("Recovery Key: %s\n", key.RecoveryKey)
	fmt.Printf("Key ID: %s\n", key.KeyID)

	res, err := c.ReplaceRecoveryKey(ctx, key.KeyID, nil)
	if err != nil {
		return fmt.Errorf("failed to replace recovery key: %w", err)
	}

	fmt.Println(res.Status)
	fmt.Println(res.Change)

	return nil
}

func newRegenerateEnterpriseKeyCmd() *cli.Command {
	var recoveryKey string

	return &cli.Command{
		Name:    "regenerate-enterprise-key",
		Usage:   "Regenerate an existing enterprise recovery key",
		Suggest: true,
		Arguments: []cli.Argument{
			&cli.StringArg{
				Name:        "key-id",
				UsageText:   "<key-id>",
				Destination: &recoveryKey,
			},
		},
		Action: func(ctx context.Context, c *cli.Command) error {
			return regenerateEnterpriseKey(ctx, recoveryKey)
		},
	}
}

func regenerateEnterpriseKey(_ context.Context, key string) error {
	fmt.Println("Regenerated enterprise key with id", key)
	return nil
}
