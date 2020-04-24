package main

import (
	"github.com/devplayg/hippo/v2"
	"time"
)

func main() {
	config := &hippo.Config{
		Name:        "Simple Server",
		Description: "simple server based on Hippo engine",
		Version:     "2.1",
		Debug:       true,
		Trace:       false,
	}
	engine := hippo.NewEngine(&Server{}, config)
	if err := engine.Start(); err != nil {
		panic(err)
	}
}

type Server struct {
	hippo.Launcher // DO NOT REMOVE; Launcher links server and engine each other.
}

func (s *Server) Start() error {
	s.Log.Debug("server has been started")

	for {
		// Do your repetitive jobs
		s.Log.Info("server is working on it")

		// Intentional error
		// return errors.New("intentional error")

		select {
		case <-s.Ctx.Done(): // for gracefully shutdown
			s.Log.Debug("server canceled; no longer works")
			return nil
		case <-time.After(2 * time.Second):
		}
	}
}

func (s *Server) Stop() error {
	s.Log.Debug("server has been stopped")
	return nil
}
