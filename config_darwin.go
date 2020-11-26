// +build darwin

package xos

import (
	"os/user"
	"path/filepath"
)

func basePath() (string, error) {
	u, err := user.Current()
	if err != nil {
		return "", err
	}

	dir := filepath.Join(u.HomeDir, "Library", "Application Support")
	return dir, nil
}

func baseServerPath() (string, error) {
	return basePath()
}

// Path to application-specific configuration directory.
func (cp *ConfigPath) Path() string {
	if cp.dirty {
		cp.Reset()
		dir := filepath.Join(cp.base, cp.name)
		cp.WriteString(dir)
		cp.dirty = false
	}

	return cp.String()
}
