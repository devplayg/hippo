# hippo

The `hippo` is an easy, fast, lightweight engine which supports gracefully shutting down the servers.

[![Build Status](https://travis-ci.org/devplayg/hippo.svg?branch=master)](https://travis-ci.org/devplayg/hippo)
[![Go Report Card](https://goreportcard.com/badge/github.com/devplayg/hippo)](https://goreportcard.com/report/github.com/devplayg/hippo)

![Hippo](hippo.png)

Import it in your program as:

```go
import "github.com/devplayg/hippo/v2"
```

## 1. Simple server 

```go
type SimpleServer struct {
    hippo.Launcher // DO NOT REMOVE; Launcher links servers and engine each other.
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

## 2. Normal server

Shutting down the server gracefully 

```go
type NormalServer struct {
    hippo.Launcher // DO NOT REMOVE; Launcher links servers and engine each other.
}

func (s *NormalServer) Start() error {
    s.Log.Debug("server has been started")

    for {
        // Do your repetitive jobs
        s.Log.Info("server is working on it")

        // Intentional error
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
    
    
## 3. Server including HTTP Server

Shutting down the server including HTTP server

https://github.com/devplayg/hippo/blob/master/examples/normal-http-server/server.go

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
    
## 4. Server with multiprocessing

Shutting down server with multiprocessing

https://github.com/devplayg/hippo/blob/master/examples/server-with-multiprocessing/server.go

Console output

    time="2020-03-04T15:16:55+09:00" level=debug msg="engine has been started"
    time="2020-03-04T15:16:55+09:00" level=debug msg="server has been started"
    time="2020-03-04T15:16:55+09:00" level=debug msg="8 workers are ready"
    time="2020-03-04T15:16:55+09:00" level=debug msg="got 3 tasks"
    time="2020-03-04T15:16:55+09:00" level=debug msg="working task-0"
    time="2020-03-04T15:16:55+09:00" level=debug msg="working task-1"
    time="2020-03-04T15:16:55+09:00" level=debug msg="working task-2"
    time="2020-03-04T15:16:56+09:00" level=debug msg="got 6 tasks"
    time="2020-03-04T15:16:56+09:00" level=debug msg="working task-3"
    time="2020-03-04T15:16:56+09:00" level=debug msg="working task-4"
    time="2020-03-04T15:16:56+09:00" level=debug msg="done task-4"
    time="2020-03-04T15:16:56+09:00" level=debug msg="working task-5"
    time="2020-03-04T15:16:56+09:00" level=debug msg="working task-6"
    time="2020-03-04T15:16:56+09:00" level=debug msg="working task-7"
    time="2020-03-04T15:16:56+09:00" level=debug msg="working task-8"
    time="2020-03-04T15:16:57+09:00" level=debug msg="done task-3"
    time="2020-03-04T15:16:57+09:00" level=debug msg="done task-5"
    time="2020-03-04T15:16:57+09:00" level=debug msg="got 6 tasks"
    time="2020-03-04T15:16:57+09:00" level=debug msg="working task-9"
    time="2020-03-04T15:16:57+09:00" level=debug msg="working task-10"
    time="2020-03-04T15:16:58+09:00" level=debug msg="done task-10"
    time="2020-03-04T15:16:58+09:00" level=debug msg="working task-11"
    time="2020-03-04T15:16:59+09:00" level=info msg="received signal, shutting down.."
    time="2020-03-04T15:16:59+09:00" level=debug msg="done task-1"
    time="2020-03-04T15:16:59+09:00" level=debug msg="waiting for working 7 workers"
    time="2020-03-04T15:16:59+09:00" level=debug msg="working task-12"
    time="2020-03-04T15:16:59+09:00" level=debug msg="done task-2"
    time="2020-03-04T15:17:02+09:00" level=debug msg="done task-0"
    time="2020-03-04T15:17:03+09:00" level=debug msg="done task-7"
    time="2020-03-04T15:17:04+09:00" level=debug msg="done task-9"
    time="2020-03-04T15:17:05+09:00" level=debug msg="done task-6"
    time="2020-03-04T15:17:05+09:00" level=debug msg="done task-8"
    time="2020-03-04T15:17:07+09:00" level=debug msg="done task-12"
    time="2020-03-04T15:17:07+09:00" level=debug msg="done task-11"
    time="2020-03-04T15:17:07+09:00" level=debug msg="works of all workers are over"
    time="2020-03-04T15:17:07+09:00" level=debug msg="server has been stopped"
    time="2020-03-04T15:17:07+09:00" level=debug msg="engine has been stopped"    

## 5. Multiple servers

Shutting down multiple servers gracefully

https://github.com/devplayg/hippo/blob/master/examples/normal-multiple/server.go

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
