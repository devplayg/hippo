# hippo V2

The `hippo` is an easy, fast, lightweight server engine which supports gracefully shutdown.

[![Build Status](https://travis-ci.org/devplayg/hippo.svg?branch=context)](https://travis-ci.org/devplayg/hippo)
[![Go Report Card](https://goreportcard.com/badge/github.com/devplayg/hippo)](https://goreportcard.com/report/github.com/devplayg/hippo)

![Hippo](hippo.png)


## Server struct

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

### Log to STDOUT

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
type SimpleServer struct {
	// Launcher links servers and engines together.
	hippo.Launcher // DO NOT REMOVE
}

func (s *SimpleServer) Start() error {
	s.Engine.Log.Debug("server has been started")

	for {
		// Do your job
		s.Engine.Log.Info("server is working on it")

		// intentional error
		// return errors.New("intentional error")

		select {
		case <-s.Engine.GetContext().Done(): // for gracefully shutdown
			s.Engine.Log.Debug("server canceled; no longer works")
			return nil
		case <-time.After(2 * time.Second):
		}
	}
}

func (s *SimpleServer) Stop() error {
	s.Engine.Log.Debug("server has been stopped")
	return nil
}
```