//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License.

package uuid

import (
	"reflect"
	"regexp"
	"testing"
)

func TestNew(t *testing.T) {
	u, err := New()
	if err != nil {
		t.Fatal(err)
	}
	if reflect.ValueOf(u).IsZero() {
		t.Fatal("unexpected zero-value UUID")
	}
	s := u.String()
	match, err := regexp.MatchString(`[a-z0-9]{8}-[a-z0-9]{4}-[a-z0-9]{4}-[a-z0-9]{4}-[a-z0-9]{12}`, s)
	if err != nil {
		t.Fatal(err)
	}
	if !match {
		t.Fatalf("invalid UUID string %s", s)
	}
}
