package hippo

type Server interface {
	Start() error
	Stop() error
	setEngine(e *Engine)
}

type Launcher struct {
	Engine *Engine
}

func (l *Launcher) setEngine(e *Engine) {
	l.Engine = e
}
