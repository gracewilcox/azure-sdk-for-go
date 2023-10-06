//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.

package azquery_test

import (
	"context"
	"testing"

	"github.com/Azure/azure-sdk-for-go/sdk/monitor/azquery"
	"github.com/stretchr/testify/require"
)

var metricName = "requests/count"
var metricResourceProvider = "Microsoft.Insights/components"

func TestQueryBatch_BasicQuerySuccess(t *testing.T) {
	client := startMetricsBatchTest(t)
	timespan := azquery.TimeInterval("PT12H")

	res, err := client.QueryBatch(context.Background(), "", metricResourceProvider, [](metricName))
	require.NoError(t, err)
	_ = res

	// res, err := client.QueryResource(context.Background(), resourceURI,
	// 	&azquery.MetricsClientQueryResourceOptions{
	// 		Timespan:        to.Ptr(timespan),
	// 		Interval:        to.Ptr("PT1M"),
	// 		MetricNames:     nil,
	// 		Aggregation:     to.SliceOfPtrs(azquery.AggregationTypeAverage, azquery.AggregationTypeCount),
	// 		Top:             nil,
	// 		OrderBy:         to.Ptr("Average asc"),
	// 		Filter:          nil,
	// 		ResultType:      nil,
	// 		MetricNamespace: to.Ptr("Microsoft.AppConfiguration/configurationStores"),
	// 	})
	// require.NoError(t, err)
	// require.NotNil(t, res.Response.Timespan)
	// require.Equal(t, *res.Response.Value[0].ErrorCode, "Success")
	// require.Equal(t, *res.Response.Namespace, "Microsoft.AppConfiguration/configurationStores")

	// testSerde(t, &res)
	// testSerde(t, res.Value[0])
	// testSerde(t, res.Value[0].Name)
	// testSerde(t, res.Value[0].TimeSeries[0])
}
