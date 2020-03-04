package utils

import (
	"io/ioutil"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/require"
)

func LoadSchema(t *testing.T, loc string) string {
	path := filepath.Join("../src/response/schemas", loc+".json")

	b, err := ioutil.ReadFile(path)
	require.NoError(t, err, "Failed to load JSON schema")

	return string(b)
}
