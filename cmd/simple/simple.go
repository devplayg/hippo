package main

import (
	"github.com/devplayg/hippo"
	"log"
)

const (
	appName        = "devplayg"
	appDisplayName = "Simple server"
	appDescription = "Simple server based on Hippo "
	appVersion     = "1.0.0"
)

func main() {
	config := &hippo.Config{
		Name:        appName,
		DisplayName: appDisplayName,
		Description: appDescription,
		Version:     appVersion,
		Debug:       true,
		Verbose:     false,
	}
	simpleServer := &SimpleServer{}
	engine := hippo.NewEngine(simpleServer, config)
	err := engine.Start()
	if err != nil {
		log.Fatal(err)
	}
}

type SimpleServer struct{}

func (s *SimpleServer) Start() error {

	return nil
}

func (s *SimpleServer) Stop() error {
	return nil
}
