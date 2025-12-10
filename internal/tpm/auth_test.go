package tpm_test

import (
	"context"
	"testing"

	"github.com/nalgeon/be"
	"snap-tpmctl/internal/testutils"
	"snap-tpmctl/internal/tpm"
)

//nolint:dupl // Similar test pattern for different function (ReplacePassphrase vs ReplacePIN)
func TestReplacePassphrase(t *testing.T) {
	t.Parallel()

	tests := map[string]struct {
		oldPassphrase string
		newPassphrase string

		replacePassphraseError bool
		replacePassphraseNotOK bool

		wantErr bool
	}{
		"Success": {oldPassphrase: "old-passphrase", newPassphrase: "new-passphrase"},

		"Error when snapd down":      {oldPassphrase: "old-passphrase", newPassphrase: "new-passphrase", replacePassphraseError: true, wantErr: true},
		"Error when response not ok": {oldPassphrase: "old-passphrase", newPassphrase: "new-passphrase", replacePassphraseNotOK: true, wantErr: true},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			ctx := context.Background()
			mockClient := testutils.NewMockSnapdClient(testutils.MockConfig{
				ReplacePassphraseError: tc.replacePassphraseError,
				ReplacePassphraseNotOK: tc.replacePassphraseNotOK,
			})

			err := tpm.ReplacePassphrase(ctx, mockClient, tc.oldPassphrase, tc.newPassphrase)

			if tc.wantErr {
				be.Err(t, err)
				return
			}
			be.Err(t, err, nil)
		})
	}
}

//nolint:dupl // Similar test pattern for different function (ReplacePassphrase vs ReplacePIN)
func TestReplacePIN(t *testing.T) {
	t.Parallel()

	tests := map[string]struct {
		oldPin string
		newPin string

		replacePINError bool
		replacePINNotOK bool

		wantErr bool
	}{
		"Success": {oldPin: "123456", newPin: "654321"},

		"Error when snapd down":      {oldPin: "123456", newPin: "654321", replacePINError: true, wantErr: true},
		"Error when response not ok": {oldPin: "123456", newPin: "654321", replacePINNotOK: true, wantErr: true},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			ctx := context.Background()
			mockClient := testutils.NewMockSnapdClient(testutils.MockConfig{
				ReplacePINError: tc.replacePINError,
				ReplacePINNotOK: tc.replacePINNotOK,
			})

			err := tpm.ReplacePIN(ctx, mockClient, tc.oldPin, tc.newPin)

			if tc.wantErr {
				be.Err(t, err)
				return
			}
			be.Err(t, err, nil)
		})
	}
}
