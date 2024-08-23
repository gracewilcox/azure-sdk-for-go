package sdk

import (
	"github.com/gracewilcox/azure-sdk-for-go/sdk/tscore/internal/shared"
	"github.com/gracewilcox/azure-sdk-for-go/sdk/tscore/policy"
	"github.com/gracewilcox/azure-sdk-for-go/sdk/tscore/sdk/pipeline"
	sdkpolicy "github.com/gracewilcox/azure-sdk-for-go/sdk/tscore/sdk/policy"
	"github.com/gracewilcox/azure-sdk-for-go/sdk/tscore/tracing"
)

// PipelineOptions contains optional settings for a client's pipeline.
// Instances can be shared across calls to SDK client constructors when uniform configuration is desired.
// Zero-value fields will have their specified default values applied during use.
type PipelineOptions struct {
	// Logging configures the built-in logging policy.
	Logging policy.LogOptions

	// Retry configures the built-in retry policy.
	Retry policy.RetryOptions

	// Transport sets the transport for HTTP requests.
	Transport pipeline.Transporter

	Tracing policy.TracingOptions

	// PerCall contains custom policies to inject into the pipeline.
	// Each policy is executed once per request.
	PerCall []pipeline.Policy

	// PerRetry contains custom policies to inject into the pipeline.
	// Each policy is executed once per request, and for each retry of that request.
	PerRetry []pipeline.Policy
}

// NewPipeline creates a pipeline from connection options, with any additional policies as specified.
func NewPipeline(options *PipelineOptions) pipeline.Pipeline {
	if options == nil {
		options = &PipelineOptions{}
	}

	// we put the includeResponsePolicy at the very beginning so that the raw response
	// is populated with the final response (some policies might mutate the response)
	policies := []pipeline.Policy{sdkpolicy.NewIncludeResponsePolicy(nil)}
	policies = append(policies, options.PerCall...)
	policies = append(policies, sdkpolicy.NewRetryPolicy(&options.Retry))
	policies = append(policies, options.PerRetry...)
	policies = append(policies, sdkpolicy.NewHTTPHeaderPolicy(nil))
	policies = append(policies, sdkpolicy.NewHTTPTracePolicy(&policy.HTTPTraceOptions{AllowedQueryParams: options.Logging.AllowedQueryParams}))
	policies = append(policies, sdkpolicy.NewLogPolicy(&options.Logging))
	policies = append(policies, sdkpolicy.NewBodyDownloadPolicy(nil))
	transport := options.Transport
	if transport == nil {
		transport = defaultHTTPClient
	}
	return pipeline.New(transport, policies...)
}

type ClientOptions struct {
	Tracing policy.TracingOptions
}

// Client is a basic HTTP client.  It consists of a pipeline and tracing provider.
type Client struct {
	pl pipeline.Pipeline
	tr tracing.Tracer
}

// NewClient creates a new Client instance with the provided values.
//   - moduleName - the fully qualified name of the module where the client is defined; used by the telemetry policy and tracing provider.
//   - moduleVersion - the semantic version of the module; used by the telemetry policy and tracing provider.
//   - plOpts - pipeline configuration options; can be the zero-value
//   - options - optional client configurations; pass nil to accept the default values
func NewClient(pipeline pipeline.Pipeline, options *ClientOptions) (*Client, error) {
	if options == nil {
		options = &ClientOptions{}
	}

	// REEXAMINE WHEN TRACING is turned on or not, is it when moduleName and moduleVersion are empty strings??
	tr := options.Tracing.Provider.NewTracer(options.Tracing.ModuleName, options.Tracing.ModuleVersion)
	if tr.Enabled() && options.Tracing.Namespace != "" {
		tr.SetAttributes(tracing.Attribute{Key: shared.TracingNamespaceAttrName, Value: options.Tracing.Namespace})
	}

	return &Client{
		pl: pipeline,
		tr: tr,
	}, nil
}

// Pipeline returns the pipeline for this client.
func (c *Client) Pipeline() pipeline.Pipeline {
	return c.pl
}

// Tracer returns the tracer for this client.
func (c *Client) Tracer() tracing.Tracer {
	return c.tr
}

// TODO maybe "With" function????
// pretty up because she's public surface area
type ContextWithDeniedValues = shared.ContextWithDeniedValues
