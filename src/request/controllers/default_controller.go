package controllers

import (
	"testing"

	"github.com/gavv/httpexpect"
)

type Request struct{}

func (r *Request) CreateServer(t *testing.T) *httpexpect.Expect {
	return httpexpect.New(t, "https://jsonplaceholder.typicode.com")
}
