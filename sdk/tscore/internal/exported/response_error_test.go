//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License.

package exported

import (
	"io"
	"net/http"
	"net/url"
	"strings"
	"testing"

	"github.com/Azure/azure-sdk-for-go/sdk/tscore/internal/log"
	"github.com/stretchr/testify/require"
)

func TestNewResponseErrorNoBody(t *testing.T) {
	fakeURL, err := url.Parse("https://fakeurl.com/the/path?qp=removed")
	if err != nil {
		t.Fatal(err)
	}
	err = NewResponseError(&http.Response{
		Status:     "the system is down",
		StatusCode: http.StatusInternalServerError,
		Body:       http.NoBody,
		Request: &http.Request{
			Method: http.MethodGet,
			URL:    fakeURL,
		},
	})
	re, ok := err.(*ResponseError)
	if !ok {
		t.Fatalf("unexpected error type %T", err)
	}
	if c := re.StatusCode; c != http.StatusInternalServerError {
		t.Fatalf("unexpected status code %d", c)
	}
	const want = `GET https://fakeurl.com/the/path
--------------------------------------------------------------------------------
RESPONSE 500: the system is down
--------------------------------------------------------------------------------
Response contained no body
--------------------------------------------------------------------------------
`
	if got := re.Error(); got != want {
		t.Fatalf("\ngot:\n%s\nwant:\n%s\n", got, want)
	}
}

func TestNewResponseError(t *testing.T) {
	fakeURL, err := url.Parse("https://fakeurl.com/the/path?qp=removed")
	if err != nil {
		t.Fatal(err)
	}
	err = NewResponseError(&http.Response{
		Status:     "the system is down",
		StatusCode: http.StatusInternalServerError,
		Body:       io.NopCloser(strings.NewReader(`{ "code": "ErrorItsBroken", "message": "it's not working" }`)),
		Request: &http.Request{
			Method: http.MethodGet,
			URL:    fakeURL,
		},
	})
	re, ok := err.(*ResponseError)
	if !ok {
		t.Fatalf("unexpected error type %T", err)
	}
	if c := re.StatusCode; c != http.StatusInternalServerError {
		t.Fatalf("unexpected status code %d", c)
	}
	const want = `GET https://fakeurl.com/the/path
--------------------------------------------------------------------------------
RESPONSE 500: the system is down
--------------------------------------------------------------------------------
{
  "code": "ErrorItsBroken",
  "message": "it's not working"
}
--------------------------------------------------------------------------------
`
	if got := re.Error(); got != want {
		t.Fatalf("\ngot:\n%s\nwant:\n%s\n", got, want)
	}
}

func TestNewResponseErrorInvalidBody(t *testing.T) {
	fakeURL, err := url.Parse("https://fakeurl.com/the/path?qp=removed")
	if err != nil {
		t.Fatal(err)
	}
	err = NewResponseError(&http.Response{
		Status:     "the system is down",
		StatusCode: http.StatusInternalServerError,
		Body:       io.NopCloser(strings.NewReader("JSON error string")),
		Request: &http.Request{
			Method: http.MethodGet,
			URL:    fakeURL,
		},
	})
	re, ok := err.(*ResponseError)
	if !ok {
		t.Fatalf("unexpected error type %T", err)
	}
	if c := re.StatusCode; c != http.StatusInternalServerError {
		t.Fatalf("unexpected status code %d", c)
	}
	const want = `GET https://fakeurl.com/the/path
--------------------------------------------------------------------------------
RESPONSE 500: the system is down
--------------------------------------------------------------------------------
JSON error string
--------------------------------------------------------------------------------
`
	if got := re.Error(); got != want {
		t.Fatalf("\ngot:\n%s\nwant:\n%s\n", got, want)
	}
}

func TestNewResponseErrorXML(t *testing.T) {
	fakeURL, err := url.Parse("https://fakeurl.com/the/path?qp=removed")
	if err != nil {
		t.Fatal(err)
	}
	err = NewResponseError(&http.Response{
		Status:     "the system is down",
		StatusCode: http.StatusInternalServerError,
		Body:       io.NopCloser(strings.NewReader(`<?xml version="1.0" encoding="UTF-8" standalone="yes"?><Error><Code>ContainerAlreadyExists</Code><Message>The specified container already exists.\nRequestId:73b2473b-c1c8-4162-97bb-dc171bff61c9\nTime:2021-12-13T19:45:40.679Z</Message></Error>`)),
		Request: &http.Request{
			Method: http.MethodGet,
			URL:    fakeURL,
		},
	})
	re, ok := err.(*ResponseError)
	if !ok {
		t.Fatalf("unexpected error type %T", err)
	}
	if c := re.StatusCode; c != http.StatusInternalServerError {
		t.Fatalf("unexpected status code %d", c)
	}
	const want = `GET https://fakeurl.com/the/path
--------------------------------------------------------------------------------
RESPONSE 500: the system is down
--------------------------------------------------------------------------------
<?xml version="1.0" encoding="UTF-8" standalone="yes"?><Error><Code>ContainerAlreadyExists</Code><Message>The specified container already exists.\nRequestId:73b2473b-c1c8-4162-97bb-dc171bff61c9\nTime:2021-12-13T19:45:40.679Z</Message></Error>
--------------------------------------------------------------------------------
`
	if got := re.Error(); got != want {
		t.Fatalf("\ngot:\n%s\nwant:\n%s\n", got, want)
	}
}

func TestNilRawResponse(t *testing.T) {
	const expected = "Missing RawResponse\n--------------------------------------------------------------------------------\n--------------------------------------------------------------------------------\n"
	require.EqualValues(t, expected, (&ResponseError{}).Error())
}

func TestNilRequestInRawResponse(t *testing.T) {
	const expected = "Request information not available\n--------------------------------------------------------------------------------\nRESPONSE 400: status\n--------------------------------------------------------------------------------\nResponse contained no body\n--------------------------------------------------------------------------------\n"
	respErr := &ResponseError{
		RawResponse: &http.Response{
			Body:       http.NoBody,
			Status:     "status",
			StatusCode: http.StatusBadRequest,
		},
	}
	require.EqualValues(t, expected, respErr.Error())
}

func TestNilResponseBody(t *testing.T) {
	const expected = "Request information not available\n--------------------------------------------------------------------------------\nRESPONSE 0: \n--------------------------------------------------------------------------------\nResponse contained no body\n--------------------------------------------------------------------------------\n"
	require.EqualValues(t, expected, (&ResponseError{RawResponse: &http.Response{}}).Error())
}

func TestLogResponseError(t *testing.T) {
	fakeURL, err := url.Parse("https://fakeurl.com/the/path?qp=removed")
	require.NoError(t, err)
	rawlog := map[log.Event][]string{}
	log.SetListener(func(cls log.Event, s string) {
		rawlog[cls] = append(rawlog[cls], s)
	})
	defer log.SetListener(nil)
	_ = NewResponseError(&http.Response{
		Status:     "the system is down",
		StatusCode: http.StatusInternalServerError,
		Body:       io.NopCloser(strings.NewReader(`{ "error": { "code": "ErrorItsBroken", "message": "it's not working" } }`)),
		Request: &http.Request{
			Method: http.MethodGet,
			URL:    fakeURL,
		},
	})
	const want = `GET https://fakeurl.com/the/path
--------------------------------------------------------------------------------
RESPONSE 500: the system is down
--------------------------------------------------------------------------------
{
  "error": {
    "code": "ErrorItsBroken",
    "message": "it's not working"
  }
}
--------------------------------------------------------------------------------
`
	msg, ok := rawlog[log.EventResponseError]
	require.True(t, ok)
	require.Len(t, msg, 1)
	require.EqualValues(t, want, msg[0])
}
