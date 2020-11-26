package xos

import "os"

// Exists checks for the existence of a file or directory.
// The second return value is true if it's a directory.
func Exists(path string) (bool, bool) {
	stat, err := os.Stat(path)
	if err != nil {
		return false, false
	}

	return true, stat.IsDir()
}
