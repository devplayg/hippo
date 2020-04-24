package main

import (
	"context"
	"github.com/devplayg/hippo/v2"
	"net/http"
	"sync"
	"time"
)

func main() {
	config := &hippo.Config{
		Name:        "Simple Server",
		Description: "simple server based on Hippo engine",
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
	hippo.Launcher // DO NOT REMOVE; Launcher links server and engine each other.
}

func (s *Server) Start() error {
	wg := new(sync.WaitGroup)

	// Server 1
	wg.Add(1)
	go func() {
		defer wg.Done()
		if err := s.startServer1(); err != nil {
			s.Log.Error(err)
			s.Cancel()
			return
		}
	}()

	// Server 2
	wg.Add(1)
	go func() {
		defer wg.Done()
		if err := s.startServer2(); err != nil {
			s.Log.Error(err)
			s.Cancel()
			return
		}
	}()

	// Server 3
	wg.Add(1)
	go func() {
		defer wg.Done()
		if err := s.startHttpServer(); err != nil {
			s.Log.Error(err)
			s.Cancel()
			return
		}
	}()

	s.Log.Debug("all servers has been started")
	wg.Wait()
	return nil
}

func (s *Server) startServer1() error {
	s.Log.Debug("server-1 has been started")
	defer s.Log.Debug("server-1 has been stopped")
	for {
		s.Log.Debug("server-1 is working on it")
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
	defer s.Log.Debug("server-2 has been stopped")
	for {
		s.Log.Debug("server-2 is working on it")

		// return errors.New("intentional error on server-2; no longer works")

		select {
		case <-s.Ctx.Done(): // for gracefully shutdown
			s.Log.Debug("server-2 canceled; no longer works")
			return nil
		case <-time.After(2 * time.Second):
		}
	}
}

func (s *Server) startHttpServer() error {
	s.Log.Debug("HTTP server has been started")
	defer s.Log.Debug("HTTP server has been stopped")

	// Start HTTP server
	var srv http.Server
	ctx, cancel := context.WithCancel(context.Background())
	go func() {
		if err := srv.ListenAndServe(); err != http.ErrServerClosed {
			ctx = context.WithValue(ctx, "err", err)
			cancel()
		}
	}()

	select {
	case <-ctx.Done(): // from local context
		return ctx.Value("err").(error)

	case <-s.Ctx.Done(): // from receiver context
		if err := srv.Shutdown(context.Background()); err != http.ErrServerClosed {
			return err
		}
		return nil
	}
}

func (s *Server) Stop() error {
	s.Log.Debug("all server has been stopped")
	return nil
}
