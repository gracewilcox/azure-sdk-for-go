//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License.

package log

import (
	tslog "github.com/gracewilcox/azure-sdk-for-go/sdk/tscore/log"
)

///////////////////////////////////////////////////////////////////////////////////////////////////
// NOTE: The following are exported as public surface area from azcore.  DO NOT MODIFY
///////////////////////////////////////////////////////////////////////////////////////////////////

// Event is used to group entries.  Each group can be toggled on or off.
type Event = tslog.Event

// SetEvents is used to control which events are written to
// the log.  By default all log events are writen.
func SetEvents(cls ...Event) {
	tslog.SetEvents(cls...)
}

// SetListener will set the Logger to write to the specified listener.
func SetListener(lst func(Event, string)) {
	tslog.SetListener(lst)
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
	return tslog.Should(cls)
}

// Write invokes the underlying listener with the specified event and message.
// If the event shouldn't be logged or there is no listener then Write does nothing.
func Write(cls Event, message string) {
	tslog.Write(cls, message)
}

// Writef invokes the underlying listener with the specified event and formatted message.
// If the event shouldn't be logged or there is no listener then Writef does nothing.
func Writef(cls Event, format string, a ...interface{}) {
	tslog.Writef(cls, format, a...)
}
