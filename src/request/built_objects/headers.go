package built_objects

import (
	"github.com/Viterzgir666/Examples-API-Go/src/request/builders"
)

func NewHeadersObject() map[string]string {
	headersBuilder := &builders.Headers{}
	headers := headersBuilder.Init().SetContentType("application/json").Build()
	return headers
}
