package tpm_test

import (
	"context"
	"testing"

	"github.com/nalgeon/be"
	"snap-tpmctl/internal/testutils"
	"snap-tpmctl/internal/tpm"
)

func TestValidateRecoveryKeyName(t *testing.T) {
	t.Parallel()

	tests := map[string]struct {
		recoveryKeyName string
		enumerateFails  bool
		wantErr         bool
	}{
		"Success": {
			recoveryKeyName: "my-key",
		},
		"Error when name empty": {
			recoveryKeyName: "",
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
		"Error when name matches existing recovery Key": {
			recoveryKeyName: "additional-recovery",
			wantErr:         true,
		},
		"Error when enumerate fails": {
			recoveryKeyName: "my-key",
			enumerateFails:  true,
			wantErr:         true,
		},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			ctx := context.Background()

			// Configure mock based on test case flags
			mockClient := testutils.NewMockSnapdClient()
			mockClient.EnumerateError = tc.enumerateFails

			err := tpm.ValidateRecoveryKeyName(ctx, mockClient, tc.recoveryKeyName)

			if tc.wantErr {
				be.Err(t, err)
				return
			}
			be.Err(t, err, nil)
		})
	}
}

func TestCreateKey(t *testing.T) {
	t.Parallel()

	tests := map[string]struct {
		recoveryKeyName string

		generateKeyFails bool
		addKeyFails      bool

		wantErr bool
	}{
		"Success": {
			recoveryKeyName: "my-key",
		},
		"Error when generate key fails": {
			recoveryKeyName:  "my-key",
			generateKeyFails: true,
			wantErr:          true,
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
			mockClient.GenerateKeyError = tc.generateKeyFails
			mockClient.AddKeyError = tc.addKeyFails

			_, err := tpm.CreateKey(ctx, mockClient, tc.recoveryKeyName)

			if tc.wantErr {
				be.Err(t, err)
				return
			}
			be.Err(t, err, nil)
		})
	}
}
