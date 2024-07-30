// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License.

package runtime

import (
	"github.com/gracewilcox/azure-sdk-for-go/sdk/tscore"
	"github.com/gracewilcox/azure-sdk-for-go/sdk/tscore/runtime"
)

// KeyCredentialPolicy authorizes requests with a [azcore.KeyCredential].
type KeyCredentialPolicy = runtime.KeyCredentialPolicy

// KeyCredentialPolicyOptions contains the optional values configuring [KeyCredentialPolicy].
type KeyCredentialPolicyOptions = runtime.KeyCredentialPolicyOptions

// NewKeyCredentialPolicy creates a new instance of [KeyCredentialPolicy].
//   - cred is the [azcore.KeyCredential] used to authenticate with the service
//   - header is the name of the HTTP request header in which the key is placed
//   - options contains optional configuration, pass nil to accept the default values
func NewKeyCredentialPolicy(cred *tscore.KeyCredential, header string, options *KeyCredentialPolicyOptions) *KeyCredentialPolicy {
	return runtime.NewKeyCredentialPolicy(cred, header, options)
}
