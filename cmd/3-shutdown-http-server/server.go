package main

import (
	"context"
	"github.com/devplayg/hippo/v2"
	"net/http"
	"time"
)

func main() {
	config := &hippo.Config{
		Name:        "Simple Server",
		Description: "simple server based on Hippo",
		Version:     "2.0.0",
		Debug:       true,
		Trace:       false,
	}
	engine := hippo.NewEngine(&Server{}, config)
	if err := engine.Start(); err != nil {
		panic(err)
	}
}

type Server struct {
	// Launcher links servers and engines together.
	hippo.Launcher // DO NOT REMOVE
}

func (s *Server) Start() error {
	s.Log.Debug("server has been started")
	ch := make(chan struct{})
	go func() {
		if err := s.startWebServer(); err != nil {
			s.Log.Error(err)
		}
		close(ch)
	}()

	defer func() {
		<-ch
	}()

	for {
		// Do your job
		s.Log.Info("server is working on it")

		// return errors.New("intentional error")

		select {
		case <-s.Done: // for gracefully shutdown
			s.Log.Debug("server canceled; no longer works")
			return nil
		case <-time.After(2 * time.Second):
		}
	}

	return nil

}

func (s *Server) Stop() error {
	s.Log.Debug("server has been stopped")
	return nil
}

func (s *Server) startWebServer() error {
	var srv http.Server

	ch := make(chan struct{})
	go func() {
		<-s.Done
		//s.Log.Info("done ???")
		if err := srv.Shutdown(context.Background()); err != nil {
			s.Log.Error(err)
		}
		close(ch)
	}()

	s.Log.Info("HTTP server is about to start")
	if err := srv.ListenAndServe(); err != http.ErrServerClosed {
		s.Log.Fatalf("HTTP server ListenAndServe: %v", err)
	}
	s.Log.Info("HTTP server has been stopped")
	<-ch
	return nil
}
