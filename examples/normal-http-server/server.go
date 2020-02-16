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
		Version:     "2.1",
		Debug:       true,
		Trace:       false,
	}
	engine := hippo.NewEngine(&Server{}, config)
	if err := engine.Start(); err != nil {
		panic(err)
	}
}

type Server struct {
	hippo.Launcher // DO NOT REMOVE; Launcher links servers and engines each other.
}

func (s *Server) Start() error {
	s.Log.Debug("server has been started")
	ch := make(chan struct{})
	go func() {
		if err := s.startHttpServer(); err != nil {
			s.Log.Error(err)
		}
		close(ch)
	}()

	defer func() {
		<-ch
	}()

	for {
		// Do your repetitive jobs
		s.Log.Info("server is working on it")

		// Intentional error
		// s.Cancel() // send cancel signal to engine
		// return errors.New("intentional error")

		select {
		case <-s.Ctx.Done(): // for gracefully shutdown
			s.Log.Debug("server canceled; no longer works")
			return nil
		case <-time.After(2 * time.Second):
		}
	}
}

func (s *Server) startHttpServer() error {
	var srv http.Server

	ch := make(chan struct{})
	go func() {
		<-s.Ctx.Done()
		if err := srv.Shutdown(context.Background()); err != nil {
			s.Log.Error(err)
		}
		close(ch)
	}()

	s.Log.Debug("HTTP server has been started")
	if err := srv.ListenAndServe(); err != http.ErrServerClosed {
		s.Log.Fatalf("HTTP server ListenAndServe: %v", err)
	}
	<-ch
	s.Log.Debug("HTTP server has been stopped")
	return nil
}

func (s *Server) Stop() error {
	s.Log.Debug("server has been stopped")
	return nil
}
