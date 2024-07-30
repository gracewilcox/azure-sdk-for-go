//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License.

// Package log provides functionality for configuring logging facilities.
package log

import (
	"github.com/gracewilcox/azure-sdk-for-go/sdk/tscore/internal/log"
)

// Event is used to group entries.  Each group can be toggled on or off.
type Event = log.Event

const (
	// EventRequest entries contain information about HTTP requests.
	// This includes information like the URL, query parameters, and headers.
	EventRequest = log.EventRequest

	// EventResponse entries contain information about HTTP responses.
	// This includes information like the HTTP status code, headers, and request URL.
	EventResponse Event = log.EventResponse

	// EventResponseError entries contain information about HTTP responses that returned
	// an *tscore.ResponseError (i.e. responses with a non 2xx HTTP status code).
	// This includes the contents of ResponseError.Error().
	EventResponseError Event = log.EventResponseError

	// EventRetryPolicy entries contain information specific to the retry policy in use.
	EventRetryPolicy Event = log.EventRetryPolicy
)

// SetEvents is used to control which events are written to
// the log.  By default all log events are writen.
// NOTE: this is not goroutine safe and should be called before using SDK clients.
func SetEvents(cls ...Event) {
	log.SetEvents(cls...)
}

// SetListener will set the Logger to write to the specified Listener.
// NOTE: this is not goroutine safe and should be called before using SDK clients.
func SetListener(lst func(Event, string)) {
	log.SetListener(lst)
}

func Should(cls Event) bool {
	return log.Should(cls)
}

func Write(cls Event, message string) {
	log.Write(cls, message)
}

func Writef(cls Event, format string, a ...interface{}) {
	log.Writef(cls, format, a...)
}

// for testing purposes
func resetEvents() {
	log.TestResetEvents()
}
