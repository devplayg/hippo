// Hippo is an easy, fast, lightweight server framework.
package hippo

import (
	"github.com/sirupsen/logrus"
	"os"
	"path/filepath"
)

// Engine supports engine framework.
type Engine struct {
	WorkingDir  string
	processName string
	server      Server
	Config      *Config
}

// NewEngine allocates a new server to engine.
func NewEngine(server Server, config *Config) *Engine {
	e := Engine{
		processName: config.Name,
		server:      server,
		Config:      config,
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

	logrus.Infof("%s is started.", e.Config.Name)
	WaitForSignals()

	return nil
}

// Stop stops engine.
func (e *Engine) Stop() error {
	err := e.server.Stop()
	if err != nil {
		logrus.Error("failed to stop %s", e.Config.Name)
		return err
	}
	logrus.Infof("%s is stopped.", e.Config.Name)
	return nil
}
