package cmd

import (
	"context"
	"fmt"
	"os"
	"snap-tpmctl/internal/snapd"
	"snap-tpmctl/internal/tpm"
	"snap-tpmctl/internal/tui"

	"github.com/urfave/cli/v3"
)

func newRemovePassphraseCmd() *cli.Command {
	return &cli.Command{
		Name:  "remove-passphrase",
		Usage: "Remove passphrase authentication",
		Action: func(ctx context.Context, cmd *cli.Command) error {
			// Ensure that the user's effective ID is root
			if os.Geteuid() != 0 {
				return fmt.Errorf("this command requires elevated privileges. Please run with sudo")
			}

			c := snapd.NewClient()
			defer c.Close()

			// Load auth before validation
			if err := c.LoadAuthFromHome(); err != nil {
				return fmt.Errorf("failed to load auth: %w", err)
			}

			// TODO:
			// Validate auth is a Passphrase

			//TODO:
			// if err := tmp.AddPassphrase(ctx, c, newPassphrase); err != nil {
			// 	return err
			// }
			fmt.Println("Passphrase removed successfully")
			return nil
		},
	}
}

func newRemovePINCmd() *cli.Command {
	return &cli.Command{
		Name:  "remove-pin",
		Usage: "Remove pin authentication",
		Action: func(ctx context.Context, cmd *cli.Command) error {
			// Ensure that the user's effective ID is root
			if os.Geteuid() != 0 {
				return fmt.Errorf("this command requires elevated privileges. Please run with sudo")
			}

			c := snapd.NewClient()
			defer c.Close()

			// Load auth before validation
			if err := c.LoadAuthFromHome(); err != nil {
				return fmt.Errorf("failed to load auth: %w", err)
			}

			// TODO:
			// Validate auth is a PIN

			// TODO:
			// if err := tmp.AddPIN(ctx, c, newPIN); err != nil {
			// 	return err
			// }
			fmt.Println("PIN removed successfully")
			return nil
		},
	}
}
