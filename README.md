# hippo

The `hippo` is an easy, fast, lightweight server engine which supports gracefully shutdown.

[![Build Status](https://travis-ci.org/devplayg/hippo.svg?branch=master)](https://travis-ci.org/devplayg/hippo)
[![Go Report Card](https://goreportcard.com/badge/github.com/devplayg/hippo)](https://goreportcard.com/report/github.com/devplayg/hippo)

![Hippo](hippo-v2.png)

Import it in your program as:

```go
import "github.com/devplayg/hippo/v2"
```

## Simple server 

```go
type SimpleServer struct {
    hippo.Launcher // DO NOT REMOVE; links servers and engines each other.
}

func (s *SimpleServer) Start() error {
    return nil
}

func (s *SimpleServer) Stop() error {
    return nil
}
```

## Run

### Simple

```go
engine := hippo.NewEngine(&SimpleServer{}, nil)
if err := engine.Start(); err != nil {
    panic(err)
}
```

### Debug

```go
config := &hippo.Config{
    Debug: true,
}
engine := hippo.NewEngine(&SimpleServer{}, config)
if err := engine.Start(); err != nil {
    panic(err)
}
```

### Log to the file

```go
config := &hippo.Config{
    Debug:  true,
    LogDir: "/var/log/",
}
engine := hippo.NewEngine(&SimpleServer{}, config)
if err := engine.Start(); err != nil {
    panic(err)
}
```

## Normal server

Shutting down the server gracefully 

```go
type NormalServer struct {
    hippo.Launcher
}

func (s *NormalServer) Start() error {
    for {
        s.Log.Info("server is working on it")

        // return errors.New("intentional error")

        select {
        case <-s.Ctx.Done(): // for gracefully shutdown
            return nil
        case <-time.After(2 * time.Second):
        }
    }
}

func (s *NormalServer) Stop() error {
    return nil
}
```

console output

    DEBU[2020-02-15T19:19:17+09:00] engine has been started                      
    DEBU[2020-02-15T19:19:17+09:00] server has been started                      
    INFO[2020-02-15T19:19:17+09:00] server is working on it                      
    INFO[2020-02-15T19:19:19+09:00] server is working on it                      
    INFO[2020-02-15T19:19:20+09:00] received signal, shutting down..             
    DEBU[2020-02-15T19:19:20+09:00] server canceled; no longer works             
    DEBU[2020-02-15T19:19:20+09:00] server has been stopped                      
    DEBU[2020-02-15T19:19:20+09:00] engine has been stopped 
    
## Normal multiple servers

Shutting down multiple servers gracefully

```go
type Server struct {
    hippo.Launcher 
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
    go func(){
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
    for {
        select {
        case <-s.Ctx.Done(): // for gracefully shutdown
            return nil
        case <-time.After(2 * time.Second):
        }
    }
}

func (s *Server) startServer2() error {
    for {
        select {
        case <-s.Ctx.Done(): // for gracefully shutdown
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

    if err := srv.ListenAndServe(); err != http.ErrServerClosed {
        s.Log.Fatalf("HTTP server ListenAndServe: %v", err)
    }
    <-ch
    return nil
}

func (s *Server) Stop() error {
    return nil
}
```