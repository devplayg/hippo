package classifier

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
	//if len(c.Dir) < 1 {
	//	return
	//}
	return nil
}
