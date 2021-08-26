package mocks

import (
	"bytes"
	"io"
	"net/http"
)

type MockResponseFunc func(req *http.Request) (*http.Response, error)

type MockClient struct {
	DoFunc MockResponseFunc
}

var GetDoFunc MockResponseFunc

func (m *MockClient) Do(req *http.Request) (*http.Response, error) {
	return GetDoFunc(req)
}

func MockHTTPRequest(body string, statusCode int) MockResponseFunc {
	r := io.NopCloser(bytes.NewReader([]byte(body)))

	return func(*http.Request) (*http.Response, error) {
		return &http.Response{
			StatusCode: statusCode,
			Body:       r,
		}, nil
	}
}
