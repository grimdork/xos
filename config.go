package xos

import (
	"strings"
)

// ConfigPath holds OS-specific paths for data.
type ConfigPath struct {
	base string
	name string
	strings.Builder
	dirty bool
}

// NewConfig returns a ConfigPath for CLI/desktop apps.
func NewConfig(name string) (*ConfigPath, error) {
	cp := &ConfigPath{
		name:  name,
		dirty: true,
	}

	b, err := basePath()
	if err != nil {
		return nil, err
	}

	cp.base = b
	return cp, nil
}

// NewConfig returns a ConfigPath for server apps.
// On Darwin this is just a wrapper for NewConfig().
// On Linux or any BSD it's likely to build paths with /etc/<appname>/.
func NewServerConfig(name string) (*ConfigPath, error) {
	cp := &ConfigPath{
		name:  name,
		dirty: true,
	}

	b, err := baseServerPath()
	if err != nil {
		return nil, err
	}

	cp.base = b
	return cp, nil
}

// SetBase to something different than the default.
func (cp *ConfigPath) SetBase(s string) {
	cp.base = s
	cp.dirty = true
}
