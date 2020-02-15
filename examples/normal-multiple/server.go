package main

import (
	"context"
	"errors"
	"github.com/devplayg/hippo/v2"
	"net/http"
	"sync"
	"time"
)

func main() {
	config := &hippo.Config{
		Name:        "Simple Server",
		Description: "simple server based on Hippo",
		Version:     "2.0",
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
	wg := new(sync.WaitGroup)
	wg.Add(1)
	go func() {
		defer wg.Done()
		if err := s.startServer1(); err != nil {
			s.Log.Error(err)
		}
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		if err := s.startServer2(); err != nil {
			s.Log.Error(err)
		}
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		if err := s.startHttpServer(); err != nil {
			s.Log.Error(err)
		}
	}()

	s.Log.Info("all servers has been started")
	wg.Wait()
	return nil
}

func (s *Server) startServer1() error {
	s.Log.Debug("server-1 has been started")
	for {
		s.Log.Info("server-1 is working on it")
		return errors.New("asdfasd")
		select {
		case <-s.Ctx.Done(): // for gracefully shutdown
			s.Log.Debug("server-1 canceled; no longer works")
			return nil
		case <-time.After(2 * time.Second):
		}
	}
}

func (s *Server) startServer2() error {
	s.Log.Debug("server-2 has been started")
	for {
		s.Log.Info("server-2 is working on it")
		select {
		case <-s.Ctx.Done(): // for gracefully shutdown
			s.Log.Debug("server-2 canceled; no longer works")
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

	s.Log.Info("HTTP server has been started")
	if err := srv.ListenAndServe(); err != http.ErrServerClosed {
		s.Log.Fatalf("HTTP server ListenAndServe: %v", err)
	}
	<-ch
	s.Log.Info("HTTP server has been stopped")
	return nil
}

func (s *Server) Stop() error {
	s.Log.Debug("all server has been stopped")
	return nil
}
