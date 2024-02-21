//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License.

// This is an internal helper package to combine the complete logging APIs.
package log

import (
	azlog "github.com/Azure/azure-sdk-for-go/sdk/azcore/log"
	"github.com/Azure/azure-sdk-for-go/sdk/internal/log"
)

// KEEP
type Event = log.Event

const (
	EventRequest     = azlog.EventRequest
	EventResponse    = azlog.EventResponse
	EventRetryPolicy = azlog.EventRetryPolicy
	EventLRO         = azlog.EventLRO // REMOVE
)

// KEEP
func Write(cls log.Event, msg string) {
	log.Write(cls, msg)
}

// KEEP
func Writef(cls log.Event, format string, a ...interface{}) {
	log.Writef(cls, format, a...)
}

// KEEP
func SetListener(lst func(Event, string)) {
	log.SetListener(lst)
}

// KEEP
func Should(cls log.Event) bool {
	return log.Should(cls)
}
