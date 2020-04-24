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

	// Start USER server
	wg.Add(1)
	go func() {
		if err := s.startUserServer(wg); err != nil {
			s.Cancel() // error? Stop all servers
			return
		}
	}()

	// Start USER server
	wg.Add(1)
	go func() {
		if err := s.startHttpServer(wg); err != nil {
			s.Cancel() // error? Stop all servers
			return
		}
	}()

	// Wait for all servers to stop
	wg.Wait()

	return nil

}

func (s *Server) startUserServer(wg *sync.WaitGroup) error {
	defer wg.Done()
	for {
		// Do your repetitive jobs
		s.Log.Debug("server is working on it")

		select {
		case <-s.Ctx.Done(): // for gracefully shutdown
			return nil
		case <-time.After(2 * time.Second):
		}
	}
}

func (s *Server) startHttpServer(wg *sync.WaitGroup) error {
	defer wg.Done()

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
	s.Log.Debug("server has been stopped")
	return nil
}
