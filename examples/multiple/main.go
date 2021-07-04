package main

import (
	"context"
	"github.com/devplayg/hippo"
	"net/http"
	"sync"
	"time"
)

func main() {
	hippo := hippo.NewHippo(&Server{}, nil)
	if err := hippo.Start(); err != nil {
		panic(err)
	}
}

type Server struct {
	hippo.Launcher // DO NOT REMOVE; Launcher links server and hippo each other.
}

func (s *Server) Start() error {
	wg := new(sync.WaitGroup)

	// Server 1
	wg.Add(1)
	go func() {
		defer func() {
			s.Log.Print("server-1 has been stopped")
			wg.Done()
		}()
		s.Log.Print("server-1 has been started")
		if err := s.startServer1(); err != nil {
			s.Log.Print(err)
			s.Cancel()
			return
		}
	}()

	// Server 2
	wg.Add(1)
	go func() {
		defer func() {
			s.Log.Print("server-2 has been stopped")
			wg.Done()
		}()

		s.Log.Print("server-2 has been started")
		if err := s.startServer2(); err != nil {
			s.Log.Print(err)
			s.Cancel()
			return
		}
	}()

	// Server 3
	wg.Add(1)
	go func() {
		defer func() {
			s.Log.Print("http server has been stopped")
			wg.Done()
		}()
		s.Log.Print("http server has been started")

		if err := s.startHttpServer(); err != nil {
			s.Log.Print(err)
			s.Cancel()
			return
		}
	}()

	s.Log.Print("all servers has been started")
	wg.Wait()
	return nil
}

func (s *Server) startServer1() error {
	for {
		s.Log.Print("server-1 is working on it")
		select {
		case <-s.Ctx.Done(): // for gracefully shutdown
			s.Log.Print("server-1 canceled; no longer works")
			return nil
		case <-time.After(2 * time.Second):
		}
	}
}

func (s *Server) startServer2() error {
	for {
		s.Log.Print("server-2 is working on it")

		// return errors.New("intentional error on server-2; no longer works")

		select {
		case <-s.Ctx.Done(): // for gracefully shutdown
			s.Log.Print("server-2 canceled; no longer works")
			return nil
		case <-time.After(2 * time.Second):
		}
	}
}

func (s *Server) startHttpServer() error {
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
	s.Log.Print("all server has been stopped")
	return nil
}
