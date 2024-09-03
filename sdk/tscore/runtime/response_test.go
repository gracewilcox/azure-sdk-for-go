//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License.

package runtime

import (
	"context"
	"net/http"
	"testing"

	"github.com/gracewilcox/azure-sdk-for-go/sdk/tscore/internal/mock"
	"github.com/gracewilcox/azure-sdk-for-go/sdk/tscore/sdk"
	"github.com/gracewilcox/azure-sdk-for-go/sdk/tscore/sdk/pipeline"
	sdkpolicy "github.com/gracewilcox/azure-sdk-for-go/sdk/tscore/sdk/policy"
)

func TestResponseUnmarshalXML(t *testing.T) {
	srv, close := mock.NewServer()
	defer close()
	// include UTF8 BOM
	srv.SetResponse(mock.WithBody([]byte("\xef\xbb\xbf<testXML><SomeInt>1</SomeInt><SomeString>s</SomeString></testXML>")))
	pl := sdk.NewPipeline(&sdk.PipelineOptions{Transport: srv})
	req, err := pipeline.NewRequest(context.Background(), http.MethodGet, srv.URL())
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	resp, err := pl.Do(req)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if !HasStatusCode(resp, http.StatusOK) {
		t.Fatalf("unexpected status code: %d", resp.StatusCode)
	}
	var tx testXML
	if err := UnmarshalAsXML(resp, &tx); err != nil {
		t.Fatalf("unexpected error unmarshalling: %v", err)
	}
	if tx.SomeInt != 1 || tx.SomeString != "s" {
		t.Fatal("unexpected value")
	}
}

func TestResponseFailureStatusCode(t *testing.T) {
	srv, close := mock.NewServer()
	defer close()
	srv.SetResponse(mock.WithStatusCode(http.StatusForbidden))
	pl := sdk.NewPipeline(&sdk.PipelineOptions{Transport: srv})
	req, err := pipeline.NewRequest(context.Background(), http.MethodGet, srv.URL())
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	resp, err := pl.Do(req)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if HasStatusCode(resp, http.StatusOK) {
		t.Fatalf("unexpected status code: %d", resp.StatusCode)
	}
}

func TestResponseUnmarshalJSON(t *testing.T) {
	srv, close := mock.NewServer()
	defer close()
	srv.SetResponse(mock.WithBody([]byte(`{ "someInt": 1, "someString": "s" }`)))
	pl := sdk.NewPipeline(&sdk.PipelineOptions{Transport: srv})
	req, err := pipeline.NewRequest(context.Background(), http.MethodGet, srv.URL())
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	resp, err := pl.Do(req)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if !HasStatusCode(resp, http.StatusOK) {
		t.Fatalf("unexpected status code: %d", resp.StatusCode)
	}
	var tx testJSON
	if err := UnmarshalAsJSON(resp, &tx); err != nil {
		t.Fatalf("unexpected error unmarshalling: %v", err)
	}
	if tx.SomeInt != 1 || tx.SomeString != "s" {
		t.Fatal("unexpected value")
	}
}

func TestResponseUnmarshalJSONskipDownload(t *testing.T) {
	srv, close := mock.NewServer()
	defer close()
	srv.SetResponse(mock.WithBody([]byte(`{ "someInt": 1, "someString": "s" }`)))
	pl := sdk.NewPipeline(&sdk.PipelineOptions{Transport: srv})
	req, err := pipeline.NewRequest(context.Background(), http.MethodGet, srv.URL())
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	sdkpolicy.SkipBodyDownload(req)
	resp, err := pl.Do(req)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if !HasStatusCode(resp, http.StatusOK) {
		t.Fatalf("unexpected status code: %d", resp.StatusCode)
	}
	var tx testJSON
	if err := UnmarshalAsJSON(resp, &tx); err != nil {
		t.Fatalf("unexpected error unmarshalling: %v", err)
	}
	if tx.SomeInt != 1 || tx.SomeString != "s" {
		t.Fatal("unexpected value")
	}
}

func TestResponseUnmarshalJSONNoBody(t *testing.T) {
	srv, close := mock.NewServer()
	defer close()
	srv.SetResponse(mock.WithBody([]byte{}))
	pl := sdk.NewPipeline(&sdk.PipelineOptions{Transport: srv})
	req, err := pipeline.NewRequest(context.Background(), http.MethodGet, srv.URL())
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	resp, err := pl.Do(req)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if !HasStatusCode(resp, http.StatusOK) {
		t.Fatalf("unexpected status code: %d", resp.StatusCode)
	}
	if err := UnmarshalAsJSON(resp, nil); err != nil {
		t.Fatalf("unexpected error unmarshalling: %v", err)
	}
}

func TestResponseUnmarshalXMLNoBody(t *testing.T) {
	srv, close := mock.NewServer()
	defer close()
	srv.SetResponse(mock.WithBody([]byte{}))
	pl := sdk.NewPipeline(&sdk.PipelineOptions{Transport: srv})
	req, err := pipeline.NewRequest(context.Background(), http.MethodGet, srv.URL())
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	resp, err := pl.Do(req)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if !HasStatusCode(resp, http.StatusOK) {
		t.Fatalf("unexpected status code: %d", resp.StatusCode)
	}
	if err := UnmarshalAsXML(resp, nil); err != nil {
		t.Fatalf("unexpected error unmarshalling: %v", err)
	}
}

func TestResponseUnmarshalAsByteArrayURLFormat(t *testing.T) {
	srv, close := mock.NewServer()
	defer close()
	srv.SetResponse(mock.WithBody([]byte(`"YSBzdHJpbmcgdGhhdCBnZXRzIGVuY29kZWQgd2l0aCBiYXNlNjR1cmw"`)))
	pl := sdk.NewPipeline(&sdk.PipelineOptions{Transport: srv})
	req, err := pipeline.NewRequest(context.Background(), http.MethodGet, srv.URL())
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	resp, err := pl.Do(req)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if !HasStatusCode(resp, http.StatusOK) {
		t.Fatalf("unexpected status code: %d", resp.StatusCode)
	}
	var ba []byte
	if err := UnmarshalAsByteArray(resp, &ba, Base64URLFormat); err != nil {
		t.Fatalf("unexpected error unmarshalling: %v", err)
	}
	if string(ba) != "a string that gets encoded with base64url" {
		t.Fatalf("bad payload, got %s", string(ba))
	}
}

func TestResponseUnmarshalAsByteArrayStdFormat(t *testing.T) {
	srv, close := mock.NewServer()
	defer close()
	srv.SetResponse(mock.WithBody([]byte(`"YSBzdHJpbmcgdGhhdCBnZXRzIGVuY29kZWQgd2l0aCBiYXNlNjR1cmw="`)))
	pl := sdk.NewPipeline(&sdk.PipelineOptions{Transport: srv})
	req, err := pipeline.NewRequest(context.Background(), http.MethodGet, srv.URL())
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	resp, err := pl.Do(req)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if !HasStatusCode(resp, http.StatusOK) {
		t.Fatalf("unexpected status code: %d", resp.StatusCode)
	}
	var ba []byte
	if err := UnmarshalAsByteArray(resp, &ba, Base64StdFormat); err != nil {
		t.Fatalf("unexpected error unmarshalling: %v", err)
	}
	if string(ba) != "a string that gets encoded with base64url" {
		t.Fatalf("bad payload, got %s", string(ba))
	}
}
