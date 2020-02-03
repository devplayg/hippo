package hippo

// Hippo configuration struct
type Config struct {
	Name        string
	DisplayName string
	Description string
	Version     string
	IsService   bool
	LogDir      string
	Debug       bool
	Trace       bool
	Verbose     bool
}
