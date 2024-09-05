//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License.

package runtime

import (
	"context"
	"errors"
	"net/http"
	"strings"
	"testing"

	"github.com/gracewilcox/azure-sdk-for-go/sdk/azcore/internal/log"
	"github.com/gracewilcox/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/gracewilcox/azure-sdk-for-go/sdk/internal/mock"
	"github.com/gracewilcox/azure-sdk-for-go/sdk/tscore/sdk/pipeline"
	"github.com/stretchr/testify/require"
)

func TestPolicyLoggingSuccess(t *testing.T) {
	rawlog := map[log.Event]string{}
	log.SetListener(func(cls log.Event, s string) {
		rawlog[cls] = s
	})
	srv, close := mock.NewServer()
	defer close()
	srv.SetResponse()
	pl := pipeline.New(srv, NewLogPolicy(nil))
	req, err := NewRequest(context.Background(), http.MethodGet, srv.URL())
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	qp := req.Raw().URL.Query()
	qp.Set("api-version", "12345")
	qp.Set("sig", "redact_me")
	req.Raw().URL.RawQuery = qp.Encode()
	resp, err := pl.Do(req)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if resp.StatusCode != http.StatusOK {
		t.Fatalf("unexpected status code: %d", resp.StatusCode)
	}
	if logReq, ok := rawlog[log.EventRequest]; ok {
		// Request ==> OUTGOING REQUEST (Try=1)
		// 	GET http://127.0.0.1:49475?one=fish&sig=REDACTED
		// 	(no headers)
		if !strings.Contains(logReq, "(no headers)") {
			t.Fatal("missing (no headers)")
		}
		if !strings.Contains(logReq, "api-version=12345") {
			t.Fatal("didn't find api-version query param")
		}
		if strings.Contains(logReq, "sig=redact_me") {
			t.Fatal("sig query param wasn't redacted")
		}
	} else {
		t.Fatal("missing LogRequest")
	}
	if logResp, ok := rawlog[log.EventResponse]; ok {
		// Response ==> REQUEST/RESPONSE (Try=1/1.0034ms, OpTime=1.0034ms) -- RESPONSE SUCCESSFULLY RECEIVED
		// 	GET http://127.0.0.1:49475?one=fish&sig=REDACTED
		// 	(no headers)
		// 	--------------------------------------------------------------------------------
		// 	RESPONSE Status: 200 OK
		// 	Content-Length: [0]
		// 	Date: [Fri, 22 Nov 2019 23:48:02 GMT]
		if !strings.Contains(logResp, "RESPONSE Status: 200 OK") {
			t.Fatal("missing response status")
		}
	} else {
		t.Fatal("missing LogResponse")
	}
}

func TestPolicyLoggingError(t *testing.T) {
	rawlog := map[log.Event]string{}
	log.SetListener(func(cls log.Event, s string) {
		rawlog[cls] = s
	})
	srv, close := mock.NewServer()
	defer close()
	srv.SetError(errors.New("bogus error"))
	pl := pipeline.New(srv, NewLogPolicy(nil))
	req, err := NewRequest(context.Background(), http.MethodGet, srv.URL())
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	req.Raw().Header.Add("header", "one")
	req.Raw().Header.Add("Authorization", "redact")
	resp, err := pl.Do(req)
	if err == nil {
		t.Fatal("unexpected nil error")
	}
	if resp != nil {
		t.Fatal("unexpected respose")
	}
	if logReq, ok := rawlog[log.EventRequest]; ok {
		// Request ==> OUTGOING REQUEST (Try=1)
		// 	GET http://127.0.0.1:50057
		// 	Authorization: REDACTED
		// 	Header: [one]
		if !strings.Contains(logReq, "Authorization: REDACTED") {
			t.Fatal("missing redacted authorization header")
		}
	} else {
		t.Fatal("missing LogRequest")
	}
	if logResponse, ok := rawlog[log.EventResponse]; ok {
		// Response ==> REQUEST/RESPONSE (Try=1/0s, OpTime=0s) -- REQUEST ERROR
		// 	GET http://127.0.0.1:50057
		// 	Authorization: REDACTED
		// 	Header: [one]
		// 	--------------------------------------------------------------------------------
		// 	ERROR:
		// 	bogus error
		// 	 ...stack track...
		if !strings.Contains(logResponse, "Authorization: REDACTED") {
			t.Fatal("missing redacted authorization header")
		}
		if !strings.Contains(logResponse, "bogus error") {
			t.Fatal("missing error message")
		}
	} else {
		t.Fatal("missing LogResponse")
	}
}

func TestWithAllowedHeadersQueryParams(t *testing.T) {
	rawlog := map[log.Event]string{}
	log.SetListener(func(cls log.Event, s string) {
		rawlog[cls] = s
	})

	const (
		plAllowedHeader = "pipeline-allowed"
		plAllowedQP     = "pipeline-allowed-qp"
		clAllowedHeader = "client-allowed"
		clAllowedQP     = "client-allowed-qp"
		redactedHeader  = "redacted-header"
		redactedQP      = "redacted-qp"
	)

	srv, close := mock.NewServer()
	defer close()
	srv.AppendResponse(mock.WithHeader(plAllowedHeader, "received1"), mock.WithHeader(clAllowedHeader, "received2"), mock.WithHeader(redactedHeader, "cantseeme"))

	pl := NewPipeline("", "", PipelineOptions{
		AllowedHeaders:         []string{plAllowedHeader},
		AllowedQueryParameters: []string{plAllowedQP},
	}, &policy.ClientOptions{
		Logging: policy.LogOptions{
			AllowedHeaders:     []string{clAllowedHeader},
			AllowedQueryParams: []string{clAllowedQP},
		},
		Transport: srv,
	})

	req, err := NewRequest(context.Background(), http.MethodGet, srv.URL())
	require.NoError(t, err)
	// don't use Header.Set() as it canonicalizes the headers (our SDKs don't either)
	req.Raw().Header[plAllowedHeader] = []string{"sent1"}
	req.Raw().Header[clAllowedHeader] = []string{"sent2"}
	req.Raw().Header[redactedHeader] = []string{"cantseeme"}
	qp := req.Raw().URL.Query()
	qp.Add(plAllowedQP, "sent1")
	qp.Add(clAllowedQP, "sent2")
	qp.Add(redactedQP, "cantseeme")
	req.Raw().URL.RawQuery = qp.Encode()

	resp, err := pl.Do(req)
	require.NoError(t, err)
	require.NotNil(t, resp)

	require.Len(t, rawlog, 3)
	require.Contains(t, rawlog[log.EventRequest], "?client-allowed-qp=sent2&pipeline-allowed-qp=sent1&redacted-qp=REDACTED")
	require.Regexp(t, `client-allowed: sent2\s+pipeline-allowed: sent1\s+redacted-header: REDACTED`, rawlog[log.EventRequest])
	require.Regexp(t, `Client-Allowed: received2\s+Content-Length: 0\s+Date: (?:[a-zA-Z0-9:,\s]+)\s+Pipeline-Allowed: received1\s+Redacted-Header: REDACTED`, rawlog[log.EventResponse])
}
