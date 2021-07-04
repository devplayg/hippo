package hippo

// Config is hippo configuration
type Config struct {
	Name        string
	Description string
	Version     string
	Logger      StdLogger
}

func newDefaultConfig(processName string) *Config {
	return &Config{
		Name:        processName,
		Description: processName,
		Version:     "0.0.1",
	}
}
