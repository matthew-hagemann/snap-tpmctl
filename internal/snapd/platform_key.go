package snapd

import (
	"context"
	"net/http"
)

type AuthMode string

// AuthMode represents the authentication mode for platform keys.
const (
	AuthModePin        AuthMode = "pin"
	AuthModePassphrase AuthMode = "passphrase"
	AuthModeNone       AuthMode = "none"
)

// KDFType represents the key derivation function type.
type KDFType string

const (
	KDFTypeArgon2id KDFType = "argon2id"
	KDFTypeArgon2i  KDFType = "argon2i"
	KDFTypePBKDF2   KDFType = "pbkdf2"
)

// ReplacePlatformKey replaces the platform key with the specified authentication.
type PlatformKeyRequest struct {
	Action     string    `json:"action"`
	AuthMode   AuthMode  `json:"auth-mode"`
	Passphrase string    `json:"passphrase,omitempty"`
	Pin        string    `json:"pin,omitempty"`
	KDFTime    *int      `json:"kdf-time,omitempty"`
	KDFType    KDFType   `json:"kdf-type,omitempty"`
	KeySlots   []KeySlot `json:"keyslots,omitempty"`
}

// ReplacePlatformKey replaces the platform key with the specified authentication.
func (c *Client) ReplacePlatformKey(ctx context.Context, authMode AuthMode, pin, passphrase string) (*AsyncResponse, error) {
	body := PlatformKeyRequest{
		Action:     "replace-platform-key",
		AuthMode:   authMode,
		Pin:        pin,
		Passphrase: passphrase,
		KDFTime:    nil,
		KDFType:    "",
		KeySlots:   nil,
	}

	resp, err := c.doAsyncRequest(ctx, http.MethodPost,
		"/v2/system-volumes", nil, body)
	if err != nil {
		return nil, err
	}

	return resp, nil
}
