package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func testCreateConn(t *testing.T) {
	_, err := createConn("./tests/example.mbtiles")
	assert.NoError(t, err)
}

func testFind(t *testing.T) {
	conn, err := createConn("./tests/example.mbtiles")	
	assert.NoError(t, err)
	results, err := conn.find(&kv{key: "ROUTE_NO", value: "137"})
	assert.NoError(t, err)
	assert.Equal(t, len(results), 1)	
} 

func testUpdate(t *testing.T) {
	// Not implemented yet
}
