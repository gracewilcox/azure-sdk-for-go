//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License.

package pipeline

import (
	"bytes"
	"context"
	"errors"
	"io"
	"net/http"
	"strings"
	"testing"

	"github.com/gracewilcox/azure-sdk-for-go/sdk/tscore/internal/exported"
	"github.com/gracewilcox/azure-sdk-for-go/sdk/tscore/internal/shared"
	"github.com/stretchr/testify/require"
)

const testURL = "http://test.contoso.com/"

func TestPipelineErrors(t *testing.T) {
	pl := New(nil)
	resp, err := pl.Do(nil)
	if err == nil {
		t.Fatal("unexpected nil error")
	}
	if resp != nil {
		t.Fatal("expected nil response")
	}
	req, err := NewRequest(context.Background(), http.MethodGet, testURL)
	if err != nil {
		t.Fatal(err)
	}
	resp, err = pl.Do(req)
	if err == nil {
		t.Fatal("unexpected nil error")
	}
	if resp != nil {
		t.Fatal("expected nil response")
	}
	req.Raw().Header["Invalid"] = []string{string([]byte{0})}
	resp, err = pl.Do(req)
	if err == nil {
		t.Fatal("unexpected nil error")
	}
	if resp != nil {
		t.Fatal("expected nil response")
	}
	req, err = NewRequest(context.Background(), http.MethodGet, testURL)
	if err != nil {
		t.Fatal(err)
	}
	req.Raw().Header["Inv alid"] = []string{"value"}
	resp, err = pl.Do(req)
	if err == nil {
		t.Fatal("unexpected nil error")
	}
	if resp != nil {
		t.Fatal("expected nil response")
	}
}

type mockTransport struct {
	succeed bool
	both    bool
}

func (m *mockTransport) Do(*http.Request) (*http.Response, error) {
	if m.both {
		return nil, nil
	}
	if m.succeed {
		return &http.Response{StatusCode: http.StatusOK}, nil
	}
	return nil, errors.New("failed")
}

func TestPipelineDo(t *testing.T) {
	req, err := NewRequest(context.Background(), http.MethodGet, testURL)
	if err != nil {
		t.Fatal(err)
	}
	tp := mockTransport{succeed: true}
	pl := New(&tp)
	resp, err := pl.Do(req)
	if err != nil {
		t.Fatal(err)
	}
	if sc := resp.StatusCode; sc != http.StatusOK {
		t.Fatalf("unexpected status code %d", sc)
	}
	tp.succeed = false
	resp, err = pl.Do(req)
	if err == nil {
		t.Fatal("unexpected nil error")
	}
	if resp != nil {
		t.Fatal("expected nil response")
	}
	tp.both = true
	resp, err = pl.Do(req)
	if err == nil {
		t.Fatal("unexpected nil error")
	}
	if resp != nil {
		t.Fatal("expected nil response")
	}
}

func TestNewRequestFromRequest(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	expectedData := bytes.NewReader([]byte{1, 2, 3, 4, 5})

	httpRequest, err := http.NewRequestWithContext(ctx, "POST", "https://example.com", expectedData)
	require.NoError(t, err)

	req, err := NewRequestFromRequest(httpRequest)
	require.NoError(t, err)

	// our stream has been drained - the func has to make a copy of the body so it can be seekable.
	// so our stream should be at end.
	currentPos, err := expectedData.Seek(0, io.SeekCurrent)
	require.NoError(t, err)
	require.Equal(t, int64(5), currentPos)

	actualData, err := io.ReadAll(req.Body())
	require.NoError(t, err)
	require.Equal(t, []byte{1, 2, 3, 4, 5}, actualData)

	// now we change stuff in the policy.Request...
	replacementBuff := bytes.NewReader([]byte{6})
	err = req.SetBody(exported.NopCloser(replacementBuff), "application/coolstuff")
	require.NoError(t, err)

	// and it's automatically reflected in the http.Request, which helps us with interop
	// with other HTTP pipelines.
	require.Equal(t, "application/coolstuff", httpRequest.Header.Get("Content-Type"))
	newBytes, err := io.ReadAll(httpRequest.Body)
	require.NoError(t, err)
	require.Equal(t, []byte{6}, newBytes)
}

func TestNewRequestFromRequest_AvoidExtraCopyIfReadSeekCloser(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	expectedData := exported.NopCloser(bytes.NewReader([]byte{1, 2, 3, 4, 5}))

	httpRequest, err := http.NewRequestWithContext(ctx, "POST", "https://example.com", expectedData)
	require.NoError(t, err)

	req, err := NewRequestFromRequest(httpRequest)
	require.NoError(t, err)

	// our stream should _NOT_ get drained since it was already an io.ReadSeekCloser
	currentPos, err := expectedData.Seek(0, io.SeekCurrent)
	require.NoError(t, err)
	require.Equal(t, int64(0), currentPos)

	actualData, err := io.ReadAll(req.Body())
	require.NoError(t, err)
	require.Equal(t, []byte{1, 2, 3, 4, 5}, actualData)

	// now we change stuff in the policy.Request...
	replacementBuff := bytes.NewReader([]byte{6})
	err = req.SetBody(exported.NopCloser(replacementBuff), "application/coolstuff")
	require.NoError(t, err)

	// and it's automatically reflected in the http.Request, which helps us with interop
	// with other HTTP pipelines.
	require.Equal(t, "application/coolstuff", httpRequest.Header.Get("Content-Type"))
	newBytes, err := io.ReadAll(httpRequest.Body)
	require.NoError(t, err)
	require.Equal(t, []byte{6}, newBytes)
}

func TestNewRequest(t *testing.T) {
	req, err := NewRequest(context.Background(), http.MethodPost, testURL)
	if err != nil {
		t.Fatal(err)
	}
	if m := req.Raw().Method; m != http.MethodPost {
		t.Fatalf("unexpected method %s", m)
	}
	type myValue struct{}
	var mv myValue
	if req.OperationValue(&mv) {
		t.Fatal("expected missing custom operation value")
	}
	req.SetOperationValue(myValue{})
	if !req.OperationValue(&mv) {
		t.Fatal("missing custom operation value")
	}
}

type testPolicy struct{}

func (testPolicy) Do(*Request) (*http.Response, error) {
	return &http.Response{}, nil
}

func TestRequestPolicies(t *testing.T) {
	req, err := NewRequest(context.Background(), http.MethodPost, testURL)
	if err != nil {
		t.Fatal(err)
	}
	resp, err := req.Next()
	if err == nil {
		t.Fatal("unexpected nil error")
	}
	if resp != nil {
		t.Fatal("expected nil response")
	}
	req.policies = []Policy{}
	resp, err = req.Next()
	if err == nil {
		t.Fatal("unexpected nil error")
	}
	if resp != nil {
		t.Fatal("expected nil response")
	}
	req.policies = []Policy{testPolicy{}}
	resp, err = req.Next()
	if err != nil {
		t.Fatal(err)
	}
	if resp == nil {
		t.Fatal("unexpected nil response")
	}
}

func TestRequestBody(t *testing.T) {
	req, err := NewRequest(context.Background(), http.MethodPost, testURL)
	if err != nil {
		t.Fatal(err)
	}
	if err := req.RewindBody(); err != nil {
		t.Fatal(err)
	}
	if err := req.Close(); err != nil {
		t.Fatal(err)
	}
	if req.Body() != nil {
		t.Fatal("expected nil body")
	}
	if req.req.GetBody != nil {
		t.Fatal("expected nil GetBody")
	}
	if err := req.SetBody(exported.NopCloser(strings.NewReader("test")), "application/text"); err != nil {
		t.Fatal(err)
	}
	if req.Body() == nil {
		t.Fatal("unexpected nil body")
	}
	if req.req.GetBody == nil {
		t.Fatal("unexpected nil GetBody")
	}
	body, err := req.req.GetBody()
	if err != nil {
		t.Fatal(err)
	}
	b, err := io.ReadAll(body)
	if err != nil {
		t.Fatal(err)
	}
	if string(b) != "test" {
		t.Fatalf("unexpected body %s", string(b))
	}
	if err := req.RewindBody(); err != nil {
		t.Fatal(err)
	}
	if err := req.Close(); err != nil {
		t.Fatal(err)
	}
}

func TestRequestEmptyBody(t *testing.T) {
	req, err := NewRequest(context.Background(), http.MethodPost, testURL)
	require.NoError(t, err)
	require.NoError(t, req.SetBody(exported.NopCloser(strings.NewReader("")), "application/text"))
	require.Nil(t, req.Body())
	require.NotContains(t, req.Raw().Header, shared.HeaderContentLength)
	require.Equal(t, []string{"application/text"}, req.Raw().Header[shared.HeaderContentType])

	// SetBody should treat a nil ReadSeekCloser the same as one having no content
	req, err = NewRequest(context.Background(), http.MethodPost, testURL)
	require.NoError(t, err)
	require.NoError(t, req.SetBody(nil, ""))
	require.Nil(t, req.Body())
	require.NotContains(t, req.Raw().Header, shared.HeaderContentLength)
	require.NotContains(t, req.Raw().Header, shared.HeaderContentType)

	// SetBody should allow replacing a previously set body with an empty one
	req, err = NewRequest(context.Background(), http.MethodPost, testURL)
	require.NoError(t, err)
	require.NoError(t, req.SetBody(exported.NopCloser(strings.NewReader("content")), "application/text"))
	require.NoError(t, req.SetBody(nil, "application/json"))
	require.Nil(t, req.Body())
	require.NotContains(t, req.Raw().Header, shared.HeaderContentLength)
	require.Equal(t, []string{"application/json"}, req.Raw().Header[shared.HeaderContentType])
}

func TestRequestClone(t *testing.T) {
	req, err := NewRequest(context.Background(), http.MethodPost, testURL)
	if err != nil {
		t.Fatal(err)
	}
	if err := req.SetBody(exported.NopCloser(strings.NewReader("test")), "application/text"); err != nil {
		t.Fatal(err)
	}
	type ensureCloned struct {
		Count int
	}
	source := ensureCloned{Count: 12345}
	req.SetOperationValue(source)
	clone := req.Clone(context.Background())
	var cloned ensureCloned
	if !clone.OperationValue(&cloned) {
		t.Fatal("missing operation value")
	}
	if cloned.Count != source.Count {
		t.Fatal("wrong operation value")
	}
	if clone.body == nil {
		t.Fatal("missing body")
	}
}

func TestNewRequestFail(t *testing.T) {
	req, err := NewRequest(context.Background(), http.MethodOptions, "://test.contoso.com/")
	if err == nil {
		t.Fatal("unexpected nil error")
	}
	if req != nil {
		t.Fatal("unexpected request")
	}
	req, err = NewRequest(context.Background(), http.MethodPatch, "/missing/the/host")
	if err == nil {
		t.Fatal("unexpected nil error")
	}
	if req != nil {
		t.Fatal("unexpected request")
	}
	req, err = NewRequest(context.Background(), http.MethodPatch, "mailto://nobody.contoso.com")
	if err == nil {
		t.Fatal("unexpected nil error")
	}
	if req != nil {
		t.Fatal("unexpected request")
	}
}

func TestRequestWithContext(t *testing.T) {
	type ctxKey1 struct{}
	type ctxKey2 struct{}

	req1, err := NewRequest(context.WithValue(context.Background(), ctxKey1{}, 1), http.MethodPost, testURL)
	require.NoError(t, err)
	require.NotNil(t, req1.Raw().Context().Value(ctxKey1{}))

	req2 := req1.WithContext(context.WithValue(context.Background(), ctxKey2{}, 1))
	require.Nil(t, req2.Raw().Context().Value(ctxKey1{}))
	require.NotNil(t, req2.Raw().Context().Value(ctxKey2{}))

	// shallow copy, so changing req2 affects req1
	req2.Raw().Header.Add("added-req2", "value")
	require.EqualValues(t, "value", req1.Raw().Header.Get("added-req2"))
}

func TestSetBodyWithClobber(t *testing.T) {
	req, err := NewRequest(context.Background(), http.MethodPatch, "https://contoso.com")
	require.NoError(t, err)
	require.NotNil(t, req)
	req.req.Header.Set(shared.HeaderContentType, "clobber-me")
	require.NoError(t, SetBody(req, exported.NopCloser(strings.NewReader(`"json-string"`)), shared.ContentTypeAppJSON, true))
	require.EqualValues(t, shared.ContentTypeAppJSON, req.req.Header.Get(shared.HeaderContentType))
}

func TestSetBodyWithNoClobber(t *testing.T) {
	req, err := NewRequest(context.Background(), http.MethodPatch, "https://contoso.com")
	require.NoError(t, err)
	require.NotNil(t, req)
	const mergePatch = "application/merge-patch+json"
	req.req.Header.Set(shared.HeaderContentType, mergePatch)
	require.NoError(t, SetBody(req, exported.NopCloser(strings.NewReader(`"json-string"`)), shared.ContentTypeAppJSON, false))
	require.EqualValues(t, mergePatch, req.req.Header.Get(shared.HeaderContentType))
}
