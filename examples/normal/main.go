package main

import (
	"github.com/devplayg/hippo"
	"time"
)

func main() {
	hippo := hippo.NewHippo(&Server{}, nil)
	if err := hippo.Start(); err != nil {
		panic(err)
	}
}

type Server struct {
	hippo.Launcher // DO NOT REMOVE; Launcher links server and hippo each other.
}

func (s *Server) Start() error {
	s.Log.Print("server has been started")

	for {
		// repetitive work
		s.Log.Print("working on it")

		// Intentional error
		//return errors.New("intentional error")

		select {
		case <-s.Ctx.Done(): // for gracefully shutdown
			s.Log.Print("hippo asked me to stop working")
			return nil
		case <-time.After(2 * time.Second):
		}
	}
}

func (s *Server) Stop() error {
	s.Log.Print("server has been stopped")
	return nil
}
