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
    hippo.Launcher // DO NOT REMOVE; Launcher links servers and engines each other.
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
engine := hippo.NewEngine(&SimpleServer{}, &hippo.Config{Debug:true})
if err := engine.Start(); err != nil {
    panic(err)
}
```

### Log to the file

powered by https://github.com/sirupsen/logrus

```go
config := &hippo.Config{
    Debug:  true,
    LogDir: "/var/log/",
}
engine := hippo.NewEngine(&SimpleServer{}, config)
if err := engine.Start(); err != nil {
    s.Log.Error(err)
}
```

## Normal server

Shutting down the server gracefully 

```go
type NormalServer struct {
    hippo.Launcher // DO NOT REMOVE; Launcher links servers and engines each other.
}

func (s *NormalServer) Start() error {
    s.Log.Debug("server has been started")

    for {
        // Do your job
        s.Log.Info("server is working on it")

        // return errors.New("intentional error")

        select {
        case <-s.Ctx.Done(): // for gracefully shutdown
            s.Log.Debug("server canceled; no longer works")
            return nil
        case <-time.After(2 * time.Second):
        }
    }
}

func (s *NormalServer) Stop() error {
    s.Log.Debug("server has been stopped")
    return nil
}
```

console output

    DEBU[2020-02-16T07:36:35+09:00] engine has been started                      
    DEBU[2020-02-16T07:36:35+09:00] server has been started                      
    INFO[2020-02-16T07:36:35+09:00] server is working on it                      
    INFO[2020-02-16T07:36:37+09:00] server is working on it                      
    INFO[2020-02-16T07:36:39+09:00] server is working on it                      
    INFO[2020-02-16T07:36:40+09:00] received signal, shutting down..             
    DEBU[2020-02-16T07:36:40+09:00] server canceled; no longer works             
    DEBU[2020-02-16T07:36:40+09:00] server has been stopped                      
    DEBU[2020-02-16T07:36:40+09:00] engine has been stopped  
    
    
## Server including HTTP Server

Shutting down the server including HTTP server

```go
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
        // Do your job
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
```

Console output

    DEBU[2020-02-16T07:45:35+09:00] engine has been started                      
    DEBU[2020-02-16T07:45:35+09:00] server has been started                      
    INFO[2020-02-16T07:45:35+09:00] server is working on it                      
    DEBU[2020-02-16T07:45:35+09:00] HTTP server has been started                 
    INFO[2020-02-16T07:45:37+09:00] server is working on it                      
    INFO[2020-02-16T07:45:39+09:00] server is working on it                      
    INFO[2020-02-16T07:45:40+09:00] received signal, shutting down..             
    DEBU[2020-02-16T07:45:40+09:00] server canceled; no longer works             
    DEBU[2020-02-16T07:45:40+09:00] HTTP server has been stopped                 
    DEBU[2020-02-16T07:45:40+09:00] server has been stopped                      
    DEBU[2020-02-16T07:45:40+09:00] engine has been stopped    

## Multiple servers

Shutting down multiple servers gracefully

```go
type Server struct {
    hippo.Launcher // DO NOT REMOVE; Launcher links servers and engines each other.
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

    s.Log.Debug("all servers has been started")
    wg.Wait()
    return nil
}

func (s *Server) startServer1() error {
    s.Log.Debug("server-1 has been started")
    for {
        s.Log.Info("server-1 is working on it")
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

        // s.Cancel() // if you want to stop all server, uncomment this line
        return errors.New("intentional error on server-2; no longer works")

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

    s.Log.Debug("HTTP server has been started")
    if err := srv.ListenAndServe(); err != http.ErrServerClosed {
        s.Log.Fatalf("HTTP server ListenAndServe: %v", err)
    }
    <-ch
    s.Log.Debug("HTTP server has been stopped")
    return nil
}

func (s *Server) Stop() error {
    s.Log.Debug("all server has been stopped")
    return nil
}
```

Console output

    DEBU[2020-02-16T07:50:00+09:00] engine has been started                      
    DEBU[2020-02-16T07:50:00+09:00] all servers has been started                 
    DEBU[2020-02-16T07:50:00+09:00] HTTP server has been started                 
    DEBU[2020-02-16T07:50:00+09:00] server-1 has been started                    
    INFO[2020-02-16T07:50:00+09:00] server-1 is working on it                    
    DEBU[2020-02-16T07:50:00+09:00] server-2 has been started                    
    INFO[2020-02-16T07:50:00+09:00] server-2 is working on it                    
    ERRO[2020-02-16T07:50:00+09:00] intentional error on server-2; no longer works 
    INFO[2020-02-16T07:50:02+09:00] server-1 is working on it                    
    INFO[2020-02-16T07:50:03+09:00] received signal, shutting down..             
    DEBU[2020-02-16T07:50:03+09:00] server-1 canceled; no longer works           
    DEBU[2020-02-16T07:50:03+09:00] HTTP server has been stopped                 
    DEBU[2020-02-16T07:50:03+09:00] all server has been stopped                  
    DEBU[2020-02-16T07:50:03+09:00] engine has been stopped
