//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License.

package exported

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestNewSASCredential(t *testing.T) {
	const val1 = "foo"
	cred := NewSASCredential(val1)
	require.NotNil(t, cred)
	require.EqualValues(t, val1, SASCredentialGet(cred))
	const val2 = "bar"
	cred.Update(val2)
	require.EqualValues(t, val2, SASCredentialGet(cred))
}
