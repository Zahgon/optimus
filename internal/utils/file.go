package utils

import (
	"io"
	"os"
)

func WriteStringToFileIndexed() func(filePath, data string, writer io.Writer) error {
	_ = "STUB: not implemented"
	return nil
}

// IsPathOccupied checks whether the targeted path is already occupied
func IsPathOccupied(path string) (bool, error) { _ = "STUB: not implemented"; return false, nil }

// IsTerminal checks if file descriptor is terminal or not
func IsTerminal(f *os.File) bool { _ = "STUB: not implemented"; return false }
