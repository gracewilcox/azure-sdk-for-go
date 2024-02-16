//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License.

package tracing

import (
	"github.com/Azure/azure-sdk-for-go/sdk/tscore/tracing"
)

// SpanKind represents the role of a Span inside a Trace. Often, this defines how a Span will be processed and visualized by various backends.
type SpanKind = tracing.SpanKind

// SpanStatus represents the status of a span.
type SpanStatus = tracing.SpanStatus

const (
	// SpanKindInternal indicates the span represents an internal operation within an application.
	SpanKindInternal = tracing.SpanKindInternal

	// SpanKindServer indicates the span covers server-side handling of a request.
	SpanKindServer SpanKind = tracing.SpanKindServer

	// SpanKindClient indicates the span describes a request to a remote service.
	SpanKindClient SpanKind = tracing.SpanKindClient

	// SpanKindProducer indicates the span was created by a messaging producer.
	SpanKindProducer SpanKind = tracing.SpanKindProducer

	// SpanKindConsumer indicates the span was created by a messaging consumer.
	SpanKindConsumer SpanKind = tracing.SpanKindConsumer
)

const (
	// SpanStatusUnset is the default status code.
	SpanStatusUnset = tracing.SpanStatusUnset

	// SpanStatusError indicates the operation contains an error.
	SpanStatusError = tracing.SpanStatusError

	// SpanStatusOK indicates the operation completed successfully.
	SpanStatusOK = tracing.SpanStatusOK
)
