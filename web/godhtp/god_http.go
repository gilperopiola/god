package godhtp

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
)

/* ~-~-~-~ God's HTTP Package -~-~-~-~ */

// Sends an HTTP request. Returns the status code, response body and error.
func Send(httpCli *http.Client, method, url, body string, opts ...RequestOpt) (int, []byte, error) {
	req, err := http.NewRequest(method, url, bytes.NewReader([]byte(body)))
	if err != nil {
		return 0, []byte{}, err
	}

	BuildRequest(req, opts)

	resp, err := httpCli.Do(req)
	if ok, respStatus := ValidateResp(resp, err); !ok {
		return respStatus, []byte{}, err
	}
	defer resp.Body.Close()

	return ReadResp(resp)
}

/* ~-~-~-~ God's HTTP Responses -~-~-~-~ */

// Returns false if either the response is nil or the error is not.
// Otherwise, returns true and the status code.
func ValidateResp(resp *http.Response, err error) (bool, int) {
	if resp == nil {
		return false, 0
	}
	return err == nil, resp.StatusCode
}

// Unsafe: The caller should check for a nil resp before calling this function.
// Remember to close the response body.
func ReadResp(resp *http.Response) (int, []byte, error) {
	body, err := io.ReadAll(resp.Body)
	return resp.StatusCode, body, fmt.Errorf("could not read http resp body: %w", err)
}
