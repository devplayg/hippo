package hippo

import (
	"context"
	"github.com/sirupsen/logrus"
)

type Launcher struct {
	Engine     *Engine
	Log        *logrus.Logger
	Ctx        context.Context
	Cancel     context.CancelFunc
	WorkingDir string
}

func (l *Launcher) setEngine(e *Engine) {
	l.Engine = e
	l.Log = e.log
	l.Ctx = e.context()
	l.Cancel = e.cancel
	l.WorkingDir = e.workingDir
}
