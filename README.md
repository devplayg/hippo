# hippo

The `hippo` is an easy, fast, lightweight server engine which supports gracefully shutdown.

[![Build Status](https://travis-ci.org/devplayg/hippo.svg?branch=master)](https://travis-ci.org/devplayg/hippo)
[![Go Report Card](https://goreportcard.com/badge/github.com/devplayg/hippo)](https://goreportcard.com/report/github.com/devplayg/hippo)

![Hippo](hippo.png)


## Simple server struct

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

## Gracefully shutdown

```go
type NormalServer struct {
	hippo.Launcher // DO NOT REMOVE
}

func (s *NormalServer) Start() error {
	for {
		s.Log.Info("server is working on it")

		// return errors.New("intentional error")

		select {
		case <-s.Done: // for gracefully shutdown
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

    DEBU[2020-02-12T10:23:00+09:00] logger has been initialized                  
    DEBU[2020-02-12T10:23:00+09:00] engine has been started                      
    DEBU[2020-02-12T10:23:00+09:00] server has been started                      
    INFO[2020-02-12T10:23:00+09:00] server is working on it                      
    INFO[2020-02-12T10:23:02+09:00] server is working on it                      
    INFO[2020-02-12T10:23:03+09:00] received signal, shutting down..             
    DEBU[2020-02-12T10:23:03+09:00] server canceled; no longer works             
    DEBU[2020-02-12T10:23:03+09:00] server has been stopped                      
    DEBU[2020-02-12T10:23:03+09:00] engine has been stopped 