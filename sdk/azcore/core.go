//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License.

package azcore

import (
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/internal/exported"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/internal/shared"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/tracing"
	"github.com/Azure/azure-sdk-for-go/sdk/tscore"
)

// KEEP
// AccessToken represents an Azure service bearer access token with expiry information.
type AccessToken = exported.AccessToken

// KEEP
// TokenCredential represents a credential capable of providing an OAuth token.
type TokenCredential = exported.TokenCredential

// KEEP
// KeyCredential contains an authentication key used to authenticate to an Azure service.
type KeyCredential = tscore.KeyCredential

// KEEP
// NewKeyCredential creates a new instance of [KeyCredential] with the specified values.
//   - key is the authentication key
func NewKeyCredential(key string) *KeyCredential {
	return tscore.NewKeyCredential(key)
}

// REMOVE
// SASCredential contains a shared access signature used to authenticate to an Azure service.
type SASCredential = exported.SASCredential

// REMOVE
// NewSASCredential creates a new instance of [SASCredential] with the specified values.
//   - sas is the shared access signature
func NewSASCredential(sas string) *SASCredential {
	return exported.NewSASCredential(sas)
}

// KEEP
// NullValue is used to send an explicit 'null' within a request.
// This is typically used in JSON-MERGE-PATCH operations to delete a value.
func NullValue[T any]() T {
	return tscore.NullValue[T]()
}

// KEEP
// IsNullValue returns true if the field contains a null sentinel value.
// This is used by custom marshallers to properly encode a null value.
func IsNullValue[T any](v T) bool {
	return tscore.IsNullValue(v)
}

// KEEP- remove API version and cloud
// ClientOptions contains optional settings for a client's pipeline.
// Instances can be shared across calls to SDK client constructors when uniform configuration is desired.
// Zero-value fields will have their specified default values applied during use.
type ClientOptions = policy.ClientOptions

// Client is a basic HTTP client.  It consists of a pipeline and tracing provider.
type Client struct {
	pl runtime.Pipeline
	tr tracing.Tracer

	// cached on the client to support shallow copying with new values
	tp        tracing.Provider
	modVer    string
	namespace string
}

// KEEP
// NewClient creates a new Client instance with the provided values.
//   - moduleName - the fully qualified name of the module where the client is defined; used by the telemetry policy and tracing provider.
//   - moduleVersion - the semantic version of the module; used by the telemetry policy and tracing provider.
//   - plOpts - pipeline configuration options; can be the zero-value
//   - options - optional client configurations; pass nil to accept the default values
func NewClient(moduleName, moduleVersion string, plOpts runtime.PipelineOptions, options *ClientOptions) (*Client, error) {
	if options == nil {
		options = &ClientOptions{}
	}

	if !options.Telemetry.Disabled {
		if err := shared.ValidateModVer(moduleVersion); err != nil {
			return nil, err
		}
	}

	pl := runtime.NewPipeline(moduleName, moduleVersion, plOpts, options)

	tr := options.TracingProvider.NewTracer(moduleName, moduleVersion)
	if tr.Enabled() && plOpts.Tracing.Namespace != "" {
		tr.SetAttributes(tracing.Attribute{Key: shared.TracingNamespaceAttrName, Value: plOpts.Tracing.Namespace})
	}

	return &Client{
		pl:        pl,
		tr:        tr,
		tp:        options.TracingProvider,
		modVer:    moduleVersion,
		namespace: plOpts.Tracing.Namespace,
	}, nil
}

// Pipeline returns the pipeline for this client.
func (c *Client) Pipeline() runtime.Pipeline {
	return c.pl
}

// Tracer returns the tracer for this client.
func (c *Client) Tracer() tracing.Tracer {
	return c.tr
}

// WithClientName returns a shallow copy of the Client with its tracing client name changed to clientName.
// Note that the values for module name and version will be preserved from the source Client.
//   - clientName - the fully qualified name of the client ("package.Client"); this is used by the tracing provider when creating spans
func (c *Client) WithClientName(clientName string) *Client {
	tr := c.tp.NewTracer(clientName, c.modVer)
	if tr.Enabled() && c.namespace != "" {
		tr.SetAttributes(tracing.Attribute{Key: shared.TracingNamespaceAttrName, Value: c.namespace})
	}

	return &Client{
		pl:        c.pl,
		tr:        tr,
		tp:        c.tp,
		modVer:    c.modVer,
		namespace: c.namespace,
	}
}
