package hippo

// Hippo configuration struct
type Config struct {
	Name        string
	Description string
	Version     string
	LogDir      string
	Debug       bool
	Trace       bool
	CertFile    string
	KeyFile     string
	Insecure    bool
}

func newDefaultConfig(processName string) *Config {
	return &Config{
		Name:        processName,
		Description: processName,
		Version:     "unknown",
		LogDir:      "",
		Debug:       false,
		Trace:       false,
		Insecure:    false,
		CertFile:    "",
		KeyFile:     "",
	}
}
