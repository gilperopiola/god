package web

import (
	"bytes"
	"io"
	"net/http"
)

// TODO - Make client a param
func SendHTTP(client http.Client, method, url, body string, opts ...HTTPOption) (int, []byte, error) {
	req, err := http.NewRequest(method, url, bytes.NewReader([]byte(body)))
	if err != nil {
		return 0, []byte{}, err
	}

	for _, opt := range opts {
		opt(req)
	}

	resp, err := client.Do(req)
	if resp == nil {
		return 0, []byte{}, err
	}
	if err != nil {
		return resp.StatusCode, []byte{}, err
	}
	defer resp.Body.Close()

	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return resp.StatusCode, respBody, err
	}

	return resp.StatusCode, respBody, nil
}

/* -~-~-~- */

type HTTPOption func(*http.Request)

func WithHeader(key, value string) HTTPOption {
	return func(req *http.Request) {
		if req != nil && req.Header != nil {
			req.Header.Set(key, value)
		}
	}
}

func WithTokenAuthHeader(token string) HTTPOption {
	return WithHeader("Authorization", "Bearer "+token)
}
