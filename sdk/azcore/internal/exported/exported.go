//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License.

package exported

import (
	"context"
	"sync/atomic"

	"github.com/Azure/azure-sdk-for-go/sdk/tscore"
)

// TODO we keeping??
// AccessToken represents an Azure service bearer access token with expiry information.
// Exported as azcore.AccessToken.
type AccessToken = tscore.AccessToken

// KEEP- removing EnableCAE, keep the others??
// TokenRequestOptions contain specific parameter that may be used by credentials types when attempting to get a token.
// Exported as policy.TokenRequestOptions.
type TokenRequestOptions struct {
	// Claims are any additional claims required for the token to satisfy a conditional access policy, such as a
	// service may return in a claims challenge following an authorization failure. If a service returned the
	// claims value base64 encoded, it must be decoded before setting this field.
	Claims string

	// EnableCAE indicates whether to enable Continuous Access Evaluation (CAE) for the requested token. When true,
	// azidentity credentials request CAE tokens for resource APIs supporting CAE. Clients are responsible for
	// handling CAE challenges. If a client that doesn't handle CAE challenges receives a CAE token, it may end up
	// in a loop retrying an API call with a token that has been revoked due to CAE.
	EnableCAE bool

	// Scopes contains the list of permission scopes required for the token.
	Scopes []string

	// TenantID identifies the tenant from which to request the token. azidentity credentials authenticate in
	// their configured default tenants when this field isn't set.
	TenantID string
}

// KEEP
// TokenCredential represents a credential capable of providing an OAuth token.
// Exported as azcore.TokenCredential.
type TokenCredential interface {
	// GetToken requests an access token for the specified set of scopes.
	GetToken(ctx context.Context, options TokenRequestOptions) (AccessToken, error)
}

// REMOVE
// SASCredential contains a shared access signature used to authenticate to an Azure service.
// Exported as azcore.SASCredential.
type SASCredential struct {
	cred *keyCredential
}

// REMOVE
// NewSASCredential creates a new instance of [SASCredential] with the specified values.
//   - sas is the shared access signature
func NewSASCredential(sas string) *SASCredential {
	return &SASCredential{cred: newKeyCredential(sas)}
}

// Update replaces the existing shared access signature with the specified value.
func (k *SASCredential) Update(sas string) {
	k.cred.Update(sas)
}

// REMOVE
// SASCredentialGet returns the shared access sig for cred.
func SASCredentialGet(cred *SASCredential) string {
	return cred.cred.Get()
}

// KEEP
type keyCredential struct {
	key atomic.Value // string
}

// KEEP
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
