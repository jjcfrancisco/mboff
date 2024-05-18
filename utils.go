package main

import (
	"fmt"
	"os"
	"strings"
)

func fileExists(filename string) bool {
	info, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return false
	}
	return !info.IsDir()
}

func parseKeyValue(str string) (*kv, error) {
	parts := strings.SplitN(str, "=", 2)
	if len(parts) != 2 {
		return nil, fmt.Errorf("invalid format: expected key=value")
	}
	return &kv{key: parts[0], value: parts[1]}, nil
}
