package main

import (
	"github.com/devplayg/hippo/v2"
	"time"
)

func main() {
	config := &hippo.Config{
		Name:        "Simple Server",
		Description: "simple server based on Hippo",
		Version:     "2.0",
		Debug:       true,
		Trace:       false,
	}
	engine := hippo.NewEngine(&NormalServer{}, config)
	if err := engine.Start(); err != nil {
		panic(err)
	}
}

type NormalServer struct {
	hippo.Launcher // DO NOT REMOVE; Launcher links servers and engines each other.
}

func (s *NormalServer) Start() error {
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

func (s *NormalServer) Stop() error {
	s.Log.Debug("server has been stopped")
	return nil
}
