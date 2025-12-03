package cmd_test

import (
	"context"
	"testing"

	"github.com/nalgeon/be"
	"snap-tpmctl/cmd/tpmctl/cmd"
	"snap-tpmctl/internal/testutils"
)

func TestCreateKey(t *testing.T) {
	t.Parallel()

	tests := map[string]struct {
		recoveryKeyName string

		authFails        bool
		generateKeyFails bool
		enumerateFails   bool
		addKeyFails      bool

		wantErr bool
	}{
		// Success cases
		"Success": {
			recoveryKeyName: "my-key",
		},

		// FIXME: drop wantInErr
		// Validation errors
		"Error when name empty": {
			recoveryKeyName: "",
			wantErr:         true,
		},
		"Error when name starts with snap": {
			recoveryKeyName: "snap-key",
			wantErr:         true,
		},
		"Error when name starts with snapd": {
			recoveryKeyName: "snapd-key",
			wantErr:         true,
		},
		"Error when name starts with default": {
			recoveryKeyName: "default-key",
			wantErr:         true,
		},
		"Error when name matches additional recovery": {
			recoveryKeyName: "additional-recovery",
			wantErr:         true,
		},

		// Snapd errors
		"Error when auth fails": {
			recoveryKeyName: "my-key",
			authFails:       true,
			wantErr:         true,
		},

		"Error when generate key fails": {
			recoveryKeyName:  "my-key",
			generateKeyFails: true,
			wantErr:          true,
		},

		"Error when enumerate fails": {
			recoveryKeyName: "my-key",
			enumerateFails:  true,
			wantErr:         true,
		},

		"Error when add key fails": {
			recoveryKeyName: "my-key",
			addKeyFails:     true,
			wantErr:         true,
		},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			ctx := context.Background()
			mockClient := testutils.NewMockSnapdClient()

			// Configure mock based on test case flags
			mockClient.LoadAuthError = tc.authFails
			mockClient.GenerateKeyError = tc.generateKeyFails
			mockClient.EnumerateError = tc.enumerateFails
			mockClient.AddKeyError = tc.addKeyFails

			err := cmd.CreateKey(ctx, mockClient, tc.recoveryKeyName)

			// FIXME: this pattern needs work
			if tc.wantErr {
				be.Err(t, err)
				return
			}
			be.Err(t, err, nil)
		})
	}
}
