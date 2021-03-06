package main

import (
	"github.com/devplayg/hippo/v2"
)

// Simple server
func main() {
	engine := hippo.NewEngine(&Server{}, &hippo.Config{Debug: true})
	if err := engine.Start(); err != nil {
		panic(err)
	}
}

// Debug
//func main() {
//	config := &hippo.Config{
//		Debug: true,
//		//Trace: true,
//	}
//	engine := hippo.NewEngine(&Server{}, config)
//	if err := engine.Start(); err != nil {
//		panic(err)
//	}
//}

// Log to file
//func main() {
//	config := &hippo.Config{
//		Debug:  true,
//		LogDir: ".",
//	}
//	engine := hippo.NewEngine(&Server{}, config)
//	if err := engine.Start(); err != nil {
//		panic(err)
//	}
//}

type Server struct {
	hippo.Launcher // DO NOT REMOVE; Launcher links server and engine each other.
}

func (s *Server) Start() error {
	s.Log.Debug("server has been started")
	return nil
}

func (s *Server) Stop() error {
	s.Log.Debug("server has been stopped")
	return nil
}
