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
		Verbose:     true,
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
	log.Debug("2) server has been started")
	// Do your job
	interval := 2 * time.Second
	for {
		log.Debug("server is working on it")
		time.Sleep(3 * time.Second)
		//return errors.New("intentional error")

		select {
		case <-s.engine.GetContext().Done():
			log.Debug("3) canceled. server no longer works")
			return nil
		case <-time.After(interval):
		}
	}
	return nil
}

func (s *SimpleServer) Stop() error {
	log.Debug("5) server has been stopped")
	return nil
}

func (s *SimpleServer) SetEngine(e *hippo.Engine) {
	s.engine = e
}
