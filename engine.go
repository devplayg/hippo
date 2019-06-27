// Hippo is an easy, fast, lightweight server framework.
package hippo

import (
	"os"
	"path/filepath"
)

// Engine supports engine framework.
type Engine struct {
	WorkingDir  string
	ErrChan     chan error
	processName string
	server      Server
	logFile     string
}

// NewEngine allocates a new server to engine.
func NewEngine(server Server) *Engine {
	e := Engine{
		processName: GetProcessName(),
		server:      server,
		ErrChan:     make(chan error),
	}
	server.SetEngine(&e)

	workingDir, _ := filepath.Abs(os.Args[0])
	e.WorkingDir = filepath.Dir(workingDir)

	err := e.initLogger()
	if err != nil {
		panic(err)
	}

	return &e
}

// Start starts server and opens error channel.
func (e *Engine) Start() error {
	go drainError(e.ErrChan)

	err := e.server.Start()
	if err != nil {
		panic(err)
	}
	return nil
}

// Stop stops engine.
func (e *Engine) Stop() error {
	err := e.server.Stop()
	if err != nil {
		return err
	}
	return nil
}

func (e *Engine) initLogger() error {
	logFile := filepath.Join(e.WorkingDir, filepath.Base(e.processName)+".log")
	return initLogger(logFile, e.server.IsDebug())
}
