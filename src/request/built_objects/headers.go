package built_objects

import (
	"../builders"
)

func NewHeadersObject() map[string]string {
	headersBuilder := &builders.Headers{}
	headers := headersBuilder.Init().SetContentType("application/json").Build()
	return headers
}
