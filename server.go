package hippo

import (
	"github.com/sirupsen/logrus"
)

// Server is the interface that configures the server
type Server interface {
	Start() error
	Stop() error
	setEngine(e *Engine)
}

type Launcher struct {
	Engine *Engine
	Log    *logrus.Logger
	Done   <-chan struct{}
}

func (l *Launcher) setEngine(e *Engine) {
	l.Engine = e
	l.Log = e.log
	l.Done = e.getContext().Done()
}
