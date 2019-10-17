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
	e := Engine{
		processName: config.Name,
		server:      server,
		Config:      config,
		started:     time.Now(),
	}

	workingDir, err := filepath.Abs(os.Args[0])
	if err != nil {
		panic(err)
	}
	e.WorkingDir = filepath.Dir(workingDir)
	return &e
}

// Start starts server and opens error channel.
func (e *Engine) Start() error {
	err := e.server.Start()
	if err != nil {
		return err
	}
	defer e.Stop()

	logrus.Infof("%s has been started", e.Config.DisplayName)
	if e.Config.IsService {
		WaitForSignals()
	}

	return nil
}

// Stop stops engine.
func (e *Engine) Stop() error {
	err := e.server.Stop()
	if err != nil {
		logrus.Error("failed to stop %s", e.Config.DisplayName)
		return err
	}
	logrus.Infof("%s has been stopped (running time: %s)", e.Config.DisplayName, time.Since(e.started))
	return nil
}
