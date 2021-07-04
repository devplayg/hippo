package hippo

import (
	"context"
)

type Launcher struct {
	Hippo      *Hippo
	Log        StdLogger
	Ctx        context.Context
	Cancel     context.CancelFunc
	WorkingDir string
}

func (l *Launcher) init(h *Hippo) {
	l.Hippo = h
	l.Log = h.log
	l.Ctx = h.context()
	l.Cancel = h.cancel
	l.WorkingDir = h.workingDir
}
