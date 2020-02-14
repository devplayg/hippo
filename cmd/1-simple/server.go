package main

import (
	"github.com/devplayg/hippo/v2"
)

// Simple server
func main() {
	engine := hippo.NewEngine(&SimpleServer{}, nil)
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
//	engine := hippo.NewEngine(&SimpleServer{}, config)
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
//	engine := hippo.NewEngine(&SimpleServer{}, config)
//	if err := engine.Start(); err != nil {
//		panic(err)
//	}
//}

type SimpleServer struct {
	hippo.Launcher // DO NOT REMOVE; links servers and engines each other.
}

func (s *SimpleServer) Start() error {
	s.Log.Info("server has been started")
	return nil
}

func (s *SimpleServer) Stop() error {
	s.Log.Info("server has been stopped")
	return nil
}
