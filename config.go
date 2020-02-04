package hippo

// Hippo configuration struct
type Config struct {
	Name        string
	DisplayName string
	Description string
	Version     string
	LogDir      string
	Debug       bool
	Trace       bool
}
