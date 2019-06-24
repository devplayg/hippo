package classifier

import (
	"github.com/devplayg/hippo"
	"path/filepath"
)

type Option struct {
	name        string
	description string
	version     string
	debug       bool

	Dir     string
	Storage string
}

func NewOption(name, description, version string, debug bool) *Option {
	return &Option{
		name:        name,
		description: description,
		version:     version,
		debug:       debug,
	}
}

func (c *Option) Name() string {
	return c.name
}

func (c *Option) Description() string {
	return c.description
}

func (c *Option) Version() string {
	return c.version
}

func (c *Option) Debug() bool {
	return c.debug
}

func (c *Option) Validate() error {
	if len(c.Dir) < 1 {
		return hippo.ErrorRequiredOption
	}

	abs, err := filepath.Abs(c.Dir)
	if err != nil {
		return hippo.ErrorInvalidDirectory
	}
	c.Dir = abs

	if len(c.Storage) < 1 {
		return hippo.ErrorRequiredOption
	}

	abs, err = filepath.Abs(c.Storage)
	if err != nil {
		return hippo.ErrorInvalidDirectory
	}
	c.Storage = abs
	return nil
}
