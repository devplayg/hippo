package main

import (
	"github.com/devplayg/hippo"
)

//// Simple server
//func main() {
//	engine := hippo.NewEngine(&SimpleServer{}, nil)
//	if err := engine.Start(); err != nil {
//		panic(err)
//	}
//}

// Log to stdout
//func main() {
//	config := &hippo.Config{
//		Debug: true,
//		Trace: false,
//	}
//	engine := hippo.NewEngine(&SimpleServer{}, config)
//	if err := engine.Start(); err != nil {
//		panic(err)
//	}
//}

// Log to file
func main() {
	config := &hippo.Config{
		Debug:  true,
		LogDir: ".",
	}
	engine := hippo.NewEngine(&SimpleServer{}, config)
	if err := engine.Start(); err != nil {
		panic(err)
	}
}

/*
time="2020-02-04T15:40:09+09:00" level=info msg="server has been started"
time="2020-02-04T15:40:09+09:00" level=info msg="server has been stopped"
*/

type SimpleServer struct {
	hippo.Launcher // DO NOT REMOVE; links servers and engines each other.
}

func (s *SimpleServer) Start() error {
	s.Engine.Log.Info("server has been started")
	return nil
}

func (s *SimpleServer) Stop() error {
	s.Engine.Log.Info("server has been stopped")
	return nil
}
