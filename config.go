package hippo

import (
	"log"
)

// Hippo configuration struct
type Config struct {
	Name        string
	DisplayName string
	Description string
	Version     string
	LogDir      string
	Debug       bool
	Trace       bool
	Logger      *log.Logger
}

func newDefaultConfig(processName string) *Config {
	return &Config{
		Name:        processName,
		DisplayName: processName,
		Description: processName,
		Version:     "unknown",
		LogDir:      "",
		Debug:       false,
		Trace:       false,
	}
}
