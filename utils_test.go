package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFileExists(t *testing.T) {
	exists := fileExists("./tests/example.mbtiles")
	assert.Equal(t, exists, true)
}

func TestParseKeyValue(t *testing.T) {
	kv, err := parseKeyValue("myKey=myValue")
	assert.NoError(t, err)
	assert.Equal(t, kv.key, "myKey")
	assert.Equal(t, kv.value, "myValue")
}

