package hippo

import (
	"os"
	"path/filepath"
)

type Engine struct {
	name        string
	description string
	version     string
	debug       bool
	WorkingDir  string
	processName string

	Option Option
	server Server
}

func NewEngine(server Server, option Option) *Engine {
	e := Engine{
		Option:      option,
		processName: GetProcessName(),
		server:      server,
	}
	server.SetEngine(&e)
	//e.LogPrefix = "[" + e.processName + "] "
	workingDir, _ := filepath.Abs(os.Args[0])
	e.WorkingDir = filepath.Dir(workingDir)
	return &e
}

func (e *Engine) Start() error {
	err := e.Option.Validate()
	if err != nil {
		panic(err)
	}
	err = e.server.Start()
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
