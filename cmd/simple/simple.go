package main

import (
	"github.com/devplayg/hippo"
	log "github.com/sirupsen/logrus"
	"time"
)

func main() {
	config := &hippo.Config{
		Name:        "server",
		DisplayName: "simple server",
		Description: "simple server based on Hippo",
		Version:     "1.0.1",
		Debug:       true,
		IsService:   true,
	}
	engine := hippo.NewEngine(&SimpleServer{}, config)
	if err := engine.Start(); err != nil {
		log.Fatal(err)
	}
}

type SimpleServer struct {
	engine *hippo.Engine
}

func (s *SimpleServer) Start() error {
	go func() {
		// Do work here
		for {
			log.Infof("Hello %s\n", s.engine.Config.Name)
			time.Sleep(10 * time.Second)
		}
	}()
	return nil
}

func (s *SimpleServer) Stop() error {
	return nil
}

func (s *SimpleServer) SetEngine(e *hippo.Engine) {
	s.engine = e
}
