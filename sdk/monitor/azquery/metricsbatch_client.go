//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package azquery

import (
	"context"
	"errors"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

// MetricsBatchClient contains the methods for the MetricsBatch group.
// Don't use this type directly, use a constructor function instead.
type MetricsBatchClient struct {
	internal *azcore.Client
}

// QueryBatch - Lists the metric values for multiple resources.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-05-01-preview
//   - baseURL - The regional endpoint to use, for example https://eastus.metrics.monitor.azure.com. The region should match the
//     region of the requested resources. For global resources, the region should be 'global'.
//   - subscriptionID - The subscription identifier for the resources in this batch.
//   - metricnamespace - Metric namespace that contains the requested metric names.
//   - metricnames - The names of the metrics (comma separated) to retrieve.
//   - resourceIDs - The comma separated list of resource IDs to query metrics for.
//   - options - MetricsBatchClientQueryBatchOptions contains the optional parameters for the MetricsBatchClient.QueryBatch method.
func (client *MetricsBatchClient) QueryBatch(ctx context.Context, baseURL string, subscriptionID string, metricnamespace string, metricnames []string, resourceIDs ResourceIDList, options *MetricsBatchClientQueryBatchOptions) (MetricsBatchClientQueryBatchResponse, error) {
	var err error
	req, err := client.queryBatchCreateRequest(ctx, baseURL, subscriptionID, metricnamespace, metricnames, resourceIDs, options)
	if err != nil {
		return MetricsBatchClientQueryBatchResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return MetricsBatchClientQueryBatchResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return MetricsBatchClientQueryBatchResponse{}, err
	}
	resp, err := client.queryBatchHandleResponse(httpResp)
	return resp, err
}

// queryBatchCreateRequest creates the QueryBatch request.
func (client *MetricsBatchClient) queryBatchCreateRequest(ctx context.Context, baseURL string, subscriptionID string, metricnamespace string, metricnames []string, resourceIDs ResourceIDList, options *MetricsBatchClientQueryBatchOptions) (*policy.Request, error) {
	host := "{baseUrl}"
	host = strings.ReplaceAll(host, "{baseUrl}", baseURL)
	urlPath := "/subscriptions/{subscriptionId}/metrics:getBatch"
	if subscriptionID == "" {
		return nil, errors.New("parameter subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	if options != nil && options.Starttime != nil {
		reqQP.Set("starttime", *options.Starttime)
	}
	if options != nil && options.Endtime != nil {
		reqQP.Set("endtime", *options.Endtime)
	}
	if options != nil && options.Interval != nil {
		reqQP.Set("interval", *options.Interval)
	}
	reqQP.Set("metricnamespace", metricnamespace)
	reqQP.Set("metricnames", strings.Join(metricnames, ","))
	if options != nil && options.Aggregation != nil {
		reqQP.Set("aggregation", *options.Aggregation)
	}
	if options != nil && options.Top != nil {
		reqQP.Set("top", strconv.FormatInt(int64(*options.Top), 10))
	}
	if options != nil && options.Orderby != nil {
		reqQP.Set("orderby", *options.Orderby)
	}
	if options != nil && options.Filter != nil {
		reqQP.Set("filter", *options.Filter)
	}
	reqQP.Set("api-version", "2023-05-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, resourceIDs); err != nil {
		return nil, err
	}
	return req, nil
}

// queryBatchHandleResponse handles the QueryBatch response.
func (client *MetricsBatchClient) queryBatchHandleResponse(resp *http.Response) (MetricsBatchClientQueryBatchResponse, error) {
	result := MetricsBatchClientQueryBatchResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.MetricResultsResponse); err != nil {
		return MetricsBatchClientQueryBatchResponse{}, err
	}
	return result, nil
}
