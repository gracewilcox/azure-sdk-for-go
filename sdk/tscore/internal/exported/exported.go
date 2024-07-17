//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License.

package exported

import (
	"context"
	"encoding/base64"
	"fmt"
	"io"
	"net/http"
	"sync/atomic"
	"time"
)

type nopCloser struct {
	io.ReadSeeker
}

func (n nopCloser) Close() error {
	return nil
}

// NopCloser returns a ReadSeekCloser with a no-op close method wrapping the provided io.ReadSeeker.
// Exported as streaming.NopCloser().
func NopCloser(rs io.ReadSeeker) io.ReadSeekCloser {
	return nopCloser{rs}
}

// HasStatusCode returns true if the Response's status code is one of the specified values.
// Exported as runtime.HasStatusCode().
func HasStatusCode(resp *http.Response, statusCodes ...int) bool {
	if resp == nil {
		return false
	}
	for _, sc := range statusCodes {
		if resp.StatusCode == sc {
			return true
		}
	}
	return false
}

// AccessToken represents a bearer access token with expiry information.
// Exported as tscore.AccessToken.
type AccessToken struct {
	Token     string
	ExpiresOn time.Time
}

// TokenRequestOptions contain specific parameter that may be used by credentials types when attempting to get a token.
// Exported as policy.TokenRequestOptions.
type TokenRequestOptions struct {
	// Scopes contains the list of permission scopes required for the token.
	// TODO from charles, should be audience?? read RFC OAuth2
	Scopes []string
}

// TokenCredential represents a credential capable of providing an OAuth token.
// Exported as tscore.TokenCredential.
type TokenCredential interface {
	// GetToken requests an access token for the specified set of scopes.
	GetToken(ctx context.Context, options TokenRequestOptions) (AccessToken, error)
}

// DecodeByteArray will base-64 decode the provided string into v.
// Exported as runtime.DecodeByteArray()
func DecodeByteArray(s string, v *[]byte, format Base64Encoding) error {
	if len(s) == 0 {
		return nil
	}
	payload := string(s)
	if payload[0] == '"' {
		// remove surrounding quotes
		payload = payload[1 : len(payload)-1]
	}
	switch format {
	case Base64StdFormat:
		decoded, err := base64.StdEncoding.DecodeString(payload)
		if err == nil {
			*v = decoded
			return nil
		}
		return err
	case Base64URLFormat:
		// use raw encoding as URL format should not contain any '=' characters
		decoded, err := base64.RawURLEncoding.DecodeString(payload)
		if err == nil {
			*v = decoded
			return nil
		}
		return err
	default:
		return fmt.Errorf("unrecognized byte array format: %d", format)
	}
}

// KeyCredential contains an authentication key used to authenticate to an Azure service.
// Exported as tscore.KeyCredential.
type KeyCredential struct {
	cred *keyCredential
}

// NewKeyCredential creates a new instance of [KeyCredential] with the specified values.
//   - key is the authentication key
func NewKeyCredential(key string) *KeyCredential {
	return &KeyCredential{cred: newKeyCredential(key)}
}

// Update replaces the existing key with the specified value.
func (k *KeyCredential) Update(key string) {
	k.cred.Update(key)
}

// KeyCredentialGet returns the key for cred.
func KeyCredentialGet(cred *KeyCredential) string {
	return cred.cred.Get()
}

type keyCredential struct {
	key atomic.Value // string
}

func newKeyCredential(key string) *keyCredential {
	keyCred := keyCredential{}
	keyCred.key.Store(key)
	return &keyCred
}

func (k *keyCredential) Get() string {
	return k.key.Load().(string)
}

func (k *keyCredential) Update(key string) {
	k.key.Store(key)
}
