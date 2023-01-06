//go:build !windows

package aiozfs

import (
	"os"
)

func tempFileOnce(dir, pattern string) (*os.File, error) {
	return os.CreateTemp(dir, pattern)
}

func readFileOnce(filename string) ([]byte, error) {
	return os.ReadFile(filename)
}
