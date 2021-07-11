package main

import (
	"context"
	"fmt"
	"github.com/devplayg/hippo/v3"
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

	// Start server
	wg.Add(1)
	go func() {
		defer func() {
			log.Print("server has been stopped")
			wg.Done()
		}()
		log.Print("server has been started")
		if err := s.startRepetitiveWork(); err != nil {
			log.Print(fmt.Errorf("repetitive work error; %w", err))
			s.Cancel() // error? Stop all servers
			return
		}
	}()

	// Start http server
	wg.Add(1)
	go func() {
		defer func() {
			log.Print("http server has been stopped")
			wg.Done()
		}()
		log.Print("http server has been started")
		if err := s.startHttpServer(); err != nil {
			log.Print(fmt.Errorf("http server error; %w", err))
			s.Cancel() // error? Stop all servers
			return
		}
	}()

	// Wait for all servers to stop
	wg.Wait()

	return nil

}

func (s *ServerWithHttp) startRepetitiveWork() error {
	for {
		s.Log.Print("repetitive work")

		select {
		case <-s.Ctx.Done(): // for gracefully shutdown
			return nil
		case <-time.After(2 * time.Second):
		}
	}
}

func (s *ServerWithHttp) startHttpServer() error {
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

func (s *ServerWithHttp) Stop() error {
	return nil
}
