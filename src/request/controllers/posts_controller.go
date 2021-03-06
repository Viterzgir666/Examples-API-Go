package controllers

import (
	"testing"

	"github.com/Viterzgir666/Examples-API-Go/src/generics"
	"github.com/gavv/httpexpect"
)

func (r *Request) AddPost(t *testing.T, headers map[string]string, jsonBody []byte) *httpexpect.Request {
	return r.CreateServer(t).POST(generics.PostsEndpoint).
		WithHeaders(headers).
		WithBytes(jsonBody)
}
