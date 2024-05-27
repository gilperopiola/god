package god

import (
	"bytes"
	"io"
	"net/http"
)

func SendHTTP(method, url, body string, opts ...HTTPOption) (int, []byte) {
	req, err := http.NewRequest(method, url, bytes.NewReader([]byte(body)))
	if err != nil {
		return http.StatusBadRequest, []byte(err.Error())
	}

	for _, opt := range opts {
		opt(req)
	}

	client := &http.Client{}

	resp, err := client.Do(req)
	if err != nil {
		return resp.StatusCode, []byte(err.Error())
	}
	defer resp.Body.Close()

	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return resp.StatusCode, []byte(err.Error())
	}

	return resp.StatusCode, bodyBytes
}

/* -~-~-~- */

type HTTPOption func(*http.Request)

func WithHeader(key, value string) HTTPOption {
	return func(req *http.Request) {
		req.Header.Set(key, value)
	}
}

func WithToken(token string) HTTPOption {
	return WithHeader("Authorization", "Bearer "+token)
}
