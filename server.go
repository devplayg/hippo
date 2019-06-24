package hippo

type Server interface {
	Start() error
	Stop() error
	SetEngine(e *Engine)
}

type Option interface {
	Debug() bool
	Name() string
	Version() string
	Description() string
	Validate() error
}
