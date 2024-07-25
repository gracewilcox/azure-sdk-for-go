//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License.

// This is an internal helper package to combine the complete logging APIs.
package log

import (
	"fmt"
)

const (
	// EventRequest entries contain information about HTTP requests.
	// This includes information like the URL, query parameters, and headers.
	EventRequest Event = "Request"

	// EventResponse entries contain information about HTTP responses.
	// This includes information like the HTTP status code, headers, and request URL.
	EventResponse Event = "Response"

	// EventResponseError entries contain information about HTTP responses that returned
	// an *tscore.ResponseError (i.e. responses with a non 2xx HTTP status code).
	// This includes the contents of ResponseError.Error().
	EventResponseError Event = "ResponseError"

	// EventRetryPolicy entries contain information specific to the retry policy in use.
	EventRetryPolicy Event = "Retry"
)

///////////////////////////////////////////////////////////////////////////////////////////////////
// NOTE: The following are exported as public surface area from tscore.  DO NOT MODIFY
///////////////////////////////////////////////////////////////////////////////////////////////////

// Event is used to group entries.  Each group can be toggled on or off.
type Event string

// SetEvents is used to control which events are written to
// the log.  By default all log events are writen.
func SetEvents(cls ...Event) {
	log.cls = cls
}

// SetListener will set the Logger to write to the specified listener.
func SetListener(lst func(Event, string)) {
	log.lst = lst
}

///////////////////////////////////////////////////////////////////////////////////////////////////
// END PUBLIC SURFACE AREA
///////////////////////////////////////////////////////////////////////////////////////////////////

// Should returns true if the specified log event should be written to the log.
// By default all log events will be logged.  Call SetEvents() to limit
// the log events for logging.
// If no listener has been set this will return false.
// Calling this method is useful when the message to log is computationally expensive
// and you want to avoid the overhead if its log event is not enabled.
func Should(cls Event) bool {
	if log.lst == nil {
		return false
	}
	if log.cls == nil || len(log.cls) == 0 {
		return true
	}
	for _, c := range log.cls {
		if c == cls {
			return true
		}
	}
	return false
}

// Write invokes the underlying listener with the specified event and message.
// If the event shouldn't be logged or there is no listener then Write does nothing.
func Write(cls Event, message string) {
	if !Should(cls) {
		return
	}
	log.lst(cls, message)
}

// Writef invokes the underlying listener with the specified event and formatted message.
// If the event shouldn't be logged or there is no listener then Writef does nothing.
func Writef(cls Event, format string, a ...interface{}) {
	if !Should(cls) {
		return
	}
	log.lst(cls, fmt.Sprintf(format, a...))
}

// TestResetEvents is used for TESTING PURPOSES ONLY.
func TestResetEvents() {
	log.cls = nil
}

// logger controls which events to log and writing to the underlying log.
type logger struct {
	cls []Event
	lst func(Event, string)
}

// the process-wide logger
var log logger
