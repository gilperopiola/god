package web

import (
	"bytes"
	"io"
	"net/http"
)

func SendHTTP(client http.Client, method, url, body string, opts ...HTTPOpt) (int, []byte, error) {
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
	return resp.StatusCode, respBody, err
}

/* -~-~-~- */

type HTTPOpt func(*http.Request)

func WithHeader(key, val string) HTTPOpt {
	return func(req *http.Request) {
		if req != nil && req.Header != nil {
			req.Header.Set(key, val)
		}
	}
}

func WithBearerToken(token string) HTTPOpt {
	return WithHeader("Authorization", "Bearer "+token)
}

func WithXRequestID(id string) HTTPOpt {
	return WithHeader("X-Request-Id", id)
}
