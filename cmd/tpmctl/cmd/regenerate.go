package cmd

import (
	"context"
	"fmt"
	"snap-tpmctl/internal/snapd"

	"github.com/urfave/cli/v3"
)

func newRegenerateKeyCmd() *cli.Command {
	// TODO: discuss with the snap team about how the user can get a key id
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
			// if err := isValidRecoveryKey(recoveryKey); err != nil {
			// 	return err
			// }

			return regenerateKey(ctx, recoveryKey)
		},
	}
}

func regenerateKey(ctx context.Context, _ string) error {
	c := snapd.NewClient()
	defer c.Close()

	key, err := c.GenerateRecoveryKey(ctx)
	if err != nil {
		return err
	}

	fmt.Println(key)

	res, err := c.ReplaceRecoveryKey(ctx, key.KeyID, nil)
	if err != nil {
		return err
	}

	fmt.Println(res)

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
			// if err := isValidRecoveryKey(recoveryKey); err != nil {
			// 	return err
			// }

			return regenerateEnterpriseKey(ctx, recoveryKey)
		},
	}
}

func regenerateEnterpriseKey(_ context.Context, key string) error {
	fmt.Println("Regenerated enterprise key with id", key)
	return nil
}

// func isValidRecoveryKey(k string) error {
// 	if k == "" {
// 		return fmt.Errorf("key-id cannot be empty")
// 	}

// 	matched, err := regexp.MatchString(`^[a-zA-Z0-9_-]{10}$`, k)
// 	if err != nil {
// 		return fmt.Errorf("regex validation error: %w", err)
// 	}
// 	if !matched {
// 		return fmt.Errorf("invalid key-id format: must contain only alphanumeric characters, hyphens, or underscores")
// 	}

// 	return nil
// }
