package main

import (
	"github.com/devplayg/hippo"
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
	return nil
}

func (s *Server) Stop() error {
	s.Log.Print("server has been stopped")
	return nil
}
