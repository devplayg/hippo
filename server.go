package hippo

type Server interface {
	Start() error
	Stop() error
	SetEngine(e *Engine)
	IsDebug() bool
	GetName() string
	GetVersion() string
}
