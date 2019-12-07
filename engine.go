// Hippo is an easy, fast, lightweight server framework.
package hippo

import (
	"github.com/sirupsen/logrus"
	"os"
	"path/filepath"
	"time"
)

// Engine supports engine framework.
type Engine struct {
	WorkingDir  string
	processName string
	server      Server
	Config      *Config
	started     time.Time
}

// NewEngine allocates a new server to engine.
func NewEngine(server Server, config *Config) *Engine {
	return &Engine{
		processName: config.Name,
		server:      server,
		Config:      config,
		started:     time.Now(),
	}
}

func (e *Engine) init() error {
	workingDir, err := filepath.Abs(os.Args[0])
	if err != nil {
		return err
	}
	e.WorkingDir = filepath.Dir(workingDir)
	return nil
}

// Start starts server and opens error channel.
func (e *Engine) Start() error {
	e.server.SetEngine(e)

	if err := e.server.Start(); err != nil {
		return err
	}
	defer e.Stop()

	logrus.Infof("[%s] has been started", e.Config.DisplayName)
	if e.Config.IsService {
		WaitForSignals()
	}

	return nil
}

// Stop stops engine.
func (e *Engine) Stop() error {
	if err := e.server.Stop(); err != nil {
		logrus.Error("failed to stop %s", e.Config.DisplayName)
		return err
	}
	logrus.Infof("[%s] has been stopped (running time: %s)", e.Config.DisplayName, time.Since(e.started))
	return nil
}
