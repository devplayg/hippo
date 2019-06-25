package hippo

import (
	"os"
	"path/filepath"
)

type Engine struct {
	WorkingDir  string
	ErrChan     chan error
	processName string
	server      Server
}

func NewEngine(server Server) *Engine {
	e := Engine{
		processName: GetProcessName(),
		server:      server,
		ErrChan:     make(chan error),
	}
	server.SetEngine(&e)
	workingDir, _ := filepath.Abs(os.Args[0])
	e.WorkingDir = filepath.Dir(workingDir)
	return &e
}

func (e *Engine) Start() error {
	go drainError(e.ErrChan)

	err := e.server.Start()
	if err != nil {
		panic(err)
	}
	return nil
}

func (e *Engine) Stop() error {
	err := e.server.Stop()
	if err != nil {
		return err
	}
	return nil
}
