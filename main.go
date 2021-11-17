package main

import (
	"path/filepath"
	"runtime"
)

func getPathToCurrentDir() string {
	_, b, _, _ := runtime.Caller(0)

	return filepath.Dir(b)
}
