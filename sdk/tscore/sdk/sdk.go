package sdk

import (
	"github.com/gracewilcox/azure-sdk-for-go/sdk/tscore/options"
	"github.com/gracewilcox/azure-sdk-for-go/sdk/tscore/sdk/pipeline"
	sdkpolicy "github.com/gracewilcox/azure-sdk-for-go/sdk/tscore/sdk/policy"
)

// PipelineOptions contains optional settings for a client's pipeline.
// Instances can be shared across calls to SDK client constructors when uniform configuration is desired.
// Zero-value fields will have their specified default values applied during use.
type PipelineOptions struct {
	// Logging configures the built-in logging policy.
	Logging options.LogOptions

	// Retry configures the built-in retry policy.
	Retry sdkpolicy.RetryPolicyOptions

	// Transport sets the transport for HTTP requests.
	Transport pipeline.Transporter

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

	// we put the captureResponsePolicy at the very beginning so that the raw response
	// is populated with the final response (some policies might mutate the response)
	policies := []pipeline.Policy{sdkpolicy.NewCaptureResponsePolicy(nil)}
	policies = append(policies, options.PerCall...)
	policies = append(policies, sdkpolicy.NewRetryPolicy(&options.Retry))
	policies = append(policies, options.PerRetry...)
	policies = append(policies, sdkpolicy.NewHTTPHeaderPolicy(nil))
	policies = append(policies, sdkpolicy.NewLogPolicy(&options.Logging))
	policies = append(policies, sdkpolicy.NewBodyDownloadPolicy(nil))
	transport := options.Transport
	if transport == nil {
		transport = defaultHTTPClient
	}
	return pipeline.New(transport, policies...)
}

type ClientOptions struct {
	// placeholder for future options
}

// Client is a basic HTTP client.  It consists of a pipeline.
type Client struct {
	pl pipeline.Pipeline
}

// NewClient creates a new Client instance with the provided values.
//   - options - optional client configurations; pass nil to accept the default values
func NewClient(pipeline pipeline.Pipeline, options *ClientOptions) (*Client, error) {
	if options == nil {
		options = &ClientOptions{}
	}

	return &Client{pl: pipeline}, nil
}

// Pipeline returns the pipeline for this client.
func (c *Client) Pipeline() pipeline.Pipeline {
	return c.pl
}
