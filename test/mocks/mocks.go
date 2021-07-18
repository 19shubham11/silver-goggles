package mocks

import (
	"bytes"
	"io"
	"net/http"
)

type MockClient struct {
	DoFunc func(req *http.Request) (*http.Response, error)
}

var GetDoFunc func(req *http.Request) (*http.Response, error)

func (m *MockClient) Do(req *http.Request) (*http.Response, error) {
	return GetDoFunc(req)
}

func MockHTTPRequest(body string, statusCode int) func(*http.Request) (*http.Response, error) {
	r := io.NopCloser(bytes.NewReader([]byte(body)))
	return func(*http.Request) (*http.Response, error) {
		return &http.Response{
			StatusCode: statusCode,
			Body:       r,
		}, nil
	}
}
