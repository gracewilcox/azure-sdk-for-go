// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License.

package tscore

import (
	"reflect"

	"github.com/Azure/azure-sdk-for-go/sdk/tscore/internal/exported"
	"github.com/Azure/azure-sdk-for-go/sdk/tscore/internal/shared"
	"github.com/Azure/azure-sdk-for-go/sdk/tscore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/tscore/runtime"
	"github.com/Azure/azure-sdk-for-go/sdk/tscore/tracing"
)

// AccessToken represents an Azure service bearer access token with expiry information.
type AccessToken = exported.AccessToken

// TokenCredential represents a credential capable of providing an OAuth token.
type TokenCredential = exported.TokenCredential

// KeyCredential contains an authentication key used to authenticate to an Azure service.
type KeyCredential = exported.KeyCredential

// NewKeyCredential creates a new instance of [KeyCredential] with the specified values.
//   - key is the authentication key
func NewKeyCredential(key string) *KeyCredential {
	return exported.NewKeyCredential(key)
}

// holds sentinel values used to send nulls
var nullables map[reflect.Type]interface{} = map[reflect.Type]interface{}{}

// NullValue is used to send an explicit 'null' within a request.
// This is typically used in JSON-MERGE-PATCH operations to delete a value.
func NullValue[T any]() T {
	t := shared.TypeOfT[T]()
	v, found := nullables[t]
	if !found {
		var o reflect.Value
		if k := t.Kind(); k == reflect.Map {
			o = reflect.MakeMap(t)
		} else if k == reflect.Slice {
			// empty slices appear to all point to the same data block
			// which causes comparisons to become ambiguous.  so we create
			// a slice with len/cap of one which ensures a unique address.
			o = reflect.MakeSlice(t, 1, 1)
		} else {
			o = reflect.New(t.Elem())
		}
		v = o.Interface()
		nullables[t] = v
	}
	// return the sentinel object
	return v.(T)
}

// IsNullValue returns true if the field contains a null sentinel value.
// This is used by custom marshallers to properly encode a null value.
func IsNullValue[T any](v T) bool {
	// see if our map has a sentinel object for this *T
	t := reflect.TypeOf(v)
	if o, found := nullables[t]; found {
		o1 := reflect.ValueOf(o)
		v1 := reflect.ValueOf(v)
		// we found it; return true if v points to the sentinel object.
		// NOTE: maps and slices can only be compared to nil, else you get
		// a runtime panic.  so we compare addresses instead.
		return o1.Pointer() == v1.Pointer()
	}
	// no sentinel object for this *t
	return false
}

// ClientOptions contains optional settings for a client's pipeline.
// Instances can be shared across calls to SDK client constructors when uniform configuration is desired.
// Zero-value fields will have their specified default values applied during use.
type ClientOptions = policy.ClientOptions

// Client is a basic HTTP client.  It consists of a pipeline and tracing provider.
// GRACE TODO remove modVer and namespace
type Client struct {
	pl runtime.Pipeline
	tr tracing.Tracer

	// cached on the client to support shallow copying with new values
	tp        tracing.Provider
	modVer    string
	namespace string
}

// TODO, make pipeline the only required one, put everything else in the struct
// TODO, gut as many fields as we can, see what we can get rid of
func NewCustomClient(pipeline runtime.Pipeline, tracer tracing.Tracer, tracingProvider tracing.Provider, moduleVersion string, namespace string) (*Client, error) {
	return &Client{
		pl:        pipeline,
		tr:        tracer,
		tp:        tracingProvider,
		modVer:    moduleVersion,
		namespace: namespace,
	}, nil
}

// GRACE TODO maybe rename pipeline options
// GRACE TODO rename to NewDefaultClient? Name TBD? or delete
// NewClient creates a new Client instance with the provided values.
//   - moduleName - the fully qualified name of the module where the client is defined; used by the tracing provider.
//   - moduleVersion - the semantic version of the module; used by the tracing provider.
//   - plOpts - pipeline configuration options; can be the zero-value
//   - options - optional client configurations; pass nil to accept the default values
func NewClient(moduleName, moduleVersion string, plOpts runtime.PipelineOptions, options *ClientOptions) (*Client, error) {
	if options == nil {
		options = &ClientOptions{}
	}

	pl := runtime.NewPipeline(plOpts, options)

	tr := options.TracingProvider.NewTracer(moduleName, moduleVersion)
	if tr.Enabled() && plOpts.Tracing.Namespace != "" {
		tr.SetAttributes(tracing.Attribute{Key: shared.TracingNamespaceAttrName, Value: plOpts.Tracing.Namespace})
	}

	// TODO CHANGE TO USE NEWCUSTOMCLIENT
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
	// TODO, change to use NewCustomClient
	return &Client{pl: c.pl, tr: tr, tp: c.tp, modVer: c.modVer, namespace: c.namespace}
}
