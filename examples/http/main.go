package main

import (
	"context"
	"fmt"
	"github.com/devplayg/hippo"
	"log"
	"net/http"
	"sync"
	"time"
)

func main() {
	hippo := hippo.NewHippo(&ServerWithHttp{}, nil)
	if err := hippo.Start(); err != nil {
		panic(err)
	}
}

type ServerWithHttp struct {
	hippo.Launcher // DO NOT REMOVE; Launcher links server and hippo each other.
}

func (s *ServerWithHttp) Start() error {
	wg := new(sync.WaitGroup)

	// Start USER server
	wg.Add(1)
	go func() {
		defer wg.Done()
		if err := s.startRepetitiveWork(); err != nil {
			log.Print(fmt.Errorf("repetitive work error; %w", err))
			s.Cancel() // error? Stop all servers
			return
		}
	}()

	// Start USER server
	wg.Add(1)
	go func() {
		defer wg.Done()
		if err := s.startHttpServer(); err != nil {
			log.Print(fmt.Errorf("http server error; %w", err))
			s.Cancel() // error? Stop all servers
			return
		}
	}()

	// Wait for all servers to stop
	s.Log.Print("waiting..")
	wg.Wait()

	return nil

}

func (s *ServerWithHttp) startRepetitiveWork() error {
	for {
		s.Log.Print("repetitive work")

		select {
		case <-s.Ctx.Done(): // for gracefully shutdown
			s.Log.Print("hippo asked to stop repetitive tasks")
			return nil
		case <-time.After(2 * time.Second):
		}
	}
}

func (s *ServerWithHttp) startHttpServer() error {
	var srv http.Server
	ctx, cancel := context.WithCancel(context.Background())
	go func() {
		s.Log.Print("http server has been started")
		if err := srv.ListenAndServe(); err != http.ErrServerClosed {
			ctx = context.WithValue(ctx, "err", err)
			cancel()
		}
	}()

	select {
	case <-ctx.Done(): // from local context
		return ctx.Value("err").(error)

	case <-s.Ctx.Done(): // from receiver context
		s.Log.Print("hippo asked to stop http server")
		if err := srv.Shutdown(context.Background()); err != http.ErrServerClosed {
			return err
		}
		return nil
	}
}

func (s *ServerWithHttp) Stop() error {
	s.Log.Print("server has been stopped")
	return nil
}
