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
		//defer func() {
		//}()
		interval := 5 * time.Second
		// Do work here
		for {
			log.Infof("Hello %s", s.engine.Config.Name)

			select {
			case <-s.engine.Ctx.Done():
				s.engine.Done <- true
				return

			case <-time.After(interval):
			}
		}
	}()
	return nil
}

func (s *SimpleServer) Stop() error {
	//log.Infof("%s stopped", s.engine.Config.Name)
	//if s.engine.Config.IsService {
	//	s.engine.Done<-true
	//}
	return nil
}

func (s *SimpleServer) SetEngine(e *hippo.Engine) {
	s.engine = e
}
