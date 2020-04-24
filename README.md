# hippo

The `hippo` is an easy, fast, lightweight engine which supports gracefully shutting down the servers.

[![Build Status](https://travis-ci.org/devplayg/hippo.svg?branch=master)](https://travis-ci.org/devplayg/hippo)
[![Go Report Card](https://goreportcard.com/badge/github.com/devplayg/hippo)](https://goreportcard.com/report/github.com/devplayg/hippo)

![Hippo](hippo.png)

Import it in your program as:

```go
import "github.com/devplayg/hippo/v2"
```

(Would you stop GRACEFULLY?)

![Image of Yaktocat](would-you-stop.png)

## 1. Simple server 

Simple server;
[Example](https://github.com/devplayg/hippo/blob/master/examples/simple/main.go)

```go
type Server struct {
	hippo.Launcher // DO NOT REMOVE; Launcher links server and engine each other.
}

func (s *Server) Start() error {
	return nil
}

func (s *Server) Stop() error {
	return nil
}
```

#### Run

```go
engine := hippo.NewEngine(&SimpleServer{}, nil)
if err := engine.Start(); err != nil {
    panic(err)
}
```

#### Debug

```go
engine := hippo.NewEngine(&SimpleServer{}, &hippo.Config{Debug:true})
if err := engine.Start(); err != nil {
    panic(err)
}
```

#### Log to a file (*powered by [logrus](https://github.com/sirupsen/logrus)*)

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

Output structure

    engine has been started
        server has been started
        server has been stopped
    engine has been stopped


## 2. Normal server

Shutting down the server gracefully;
[Example](https://github.com/devplayg/hippo/blob/master/examples/normal/main.go) 

```go
type Server struct {
    hippo.Launcher // DO NOT REMOVE; Launcher links server and engine each other.
}

func (s *Server) Start() error {
    for {
        // Do your repetitive jobs

        // return errors.New("intentional error")

        select {
        case <-s.Ctx.Done(): // for gracefully shutdown
            s.Log.Debug("server canceled; no longer works")
            return nil
        case <-time.After(2 * time.Second):
        }
    }
}

func (s *Server) Stop() error {
    return nil
}
```

Output structure

    engine has been started                      
        server has been started                      
            server is working on it                      
            server is working on it                      
            server is working on it                      
            received signal, shutting down..             
            server canceled; no longer works             
        server has been stopped                      
    engine has been stopped  
    
    
## 3. Server working with HTTP Server

Shutting down the server including HTTP server; 
[Example](https://github.com/devplayg/hippo/blob/master/examples/http/main.go)

Output structure

    engine has been started                      
        server has been started                      
            HTTP server has been started
            USER server has been started                 
                server is working on it                      
                server is working on it                      
                received signal, shutting down..
                HTTP server received signal; no longer works
                USER server received signal; no longer works             
            HTTP server has been stopped                 
            USER server has been stopped                 
        server has been stopped                      
    engine has been stopped
    
    
## 4. Multiple servers

Shutting down multiple servers gracefully;
[Example](https://github.com/devplayg/hippo/blob/master/examples/multiple/main.go)

Output structure

    engine has been started
        HTTP server has been started
        server-1 has been started
        server-2 has been started
            server-1 is working on it
            server-2 is working on it
                received signal, shutting down..
            server-2 canceled; no longer works
            server-1 canceled; no longer works
        server-1 has been stopped
        server-2 has been stopped
        HTTP server has been stopped
    engine has been stopped
