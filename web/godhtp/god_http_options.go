package godhtp

import (
	"net/http"
)

/* -~-~-~ Build Request Options -~-~-~ */

type RequestOpt func(*http.Request)

func BuildRequest(req *http.Request, opts []RequestOpt) {
	for _, opt := range opts {
		opt(req)
	}
}

func WithHeader(key, val string) RequestOpt {
	return func(req *http.Request) {
		if req != nil && req.Header != nil {
			req.Header.Set(key, val)
		}
	}
}

func WithJSON() RequestOpt {
	return WithHeader("Content-Type", "application/json")
}

func WithBearer(token string) RequestOpt {
	return WithHeader("Authorization", "Bearer "+token)
}

func WithXRequestID(id string) RequestOpt {
	return WithHeader("X-Request-Id", id)
}
