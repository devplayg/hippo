package main

import (
	"github.com/devplayg/hippo"
	"time"
)

func main() {
	config := &hippo.Config{
		Name:        "server",
		DisplayName: "simple server",
		Description: "simple server based on Hippo",
		Version:     "2.0.0",
		Debug:       true,
		Trace:       false,
	}
	engine := hippo.NewEngine(&SimpleServer{}, config)
	if err := engine.Start(); err != nil {
		panic(err)
	}

}

type SimpleServer struct {
	// Launcher links servers and engines together.
	hippo.Launcher // DO NOT REMOVE
}

func (s *SimpleServer) Start() error {
	s.Engine.Log.Debug("server has been started")

	// Do your job
	for {
		s.Engine.Log.Info("server is working on it")

		select {
		case <-s.Engine.GetContext().Done(): // for gracefully shutdown
			s.Engine.Log.Debug("server canceled; no longer works")
			return nil
		case <-time.After(2 * time.Second):
		}
	}
}

func (s *SimpleServer) Stop() error {
	s.Engine.Log.Debug("server has been stopped")
	return nil
}
