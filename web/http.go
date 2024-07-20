package web

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
)

/* ~-~-~-~ God's HTTP Package -~-~-~-~ */

func SendHTTP(httpCli *http.Client, method, url, body string, opts ...HttpOpt) (int, []byte, error) {
	req, err := http.NewRequest(method, url, bytes.NewReader([]byte(body)))
	if err != nil {
		return 0, []byte{}, err
	}

	applyHttpOpts(req, opts)

	resp, err := httpCli.Do(req)
	if ok, httpStatus := isHttpRespValid(resp, err); !ok {
		return httpStatus, []byte{}, err
	}
	defer resp.Body.Close()

	return ReadHTTPResponse(resp)
}

// Unsafe: This assumes you've already checked for a nil resp.
// There are reasons.
func ReadHTTPResponse(resp *http.Response) (int, []byte, error) {
	body, err := io.ReadAll(resp.Body)
	return resp.StatusCode, body, fmt.Errorf("could not read http resp body: %w", err)
}

/* -~-~-~ HTTP Options -~-~-~ */

type HttpOpt func(*http.Request)

func applyHttpOpts(req *http.Request, opts []HttpOpt) {
	for _, opt := range opts {
		opt(req)
	}
}

func WithHeader(key, val string) HttpOpt {
	return func(req *http.Request) {
		if req != nil && req.Header != nil {
			req.Header.Set(key, val)
		}
	}
}

func WithBearerToken(token string) HttpOpt {
	return WithHeader("Authorization", "Bearer "+token)
}

func WithXRequestID(id string) HttpOpt {
	return WithHeader("X-Request-Id", id)
}

/* -~-~-~ Helpers / Unsure -~-~-~ */

type HttpResp http.Response

func isHttpRespValid(resp *http.Response, err error) (bool, int) {
	if resp == nil {
		return false, 0
	}
	return err == nil, resp.StatusCode
}
