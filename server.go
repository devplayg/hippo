package hippo

// Server is the interface that configures the server
type Server interface {
	Start() error
	Stop() error
	initLauncher(e *Hippo) error
}
