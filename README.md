# Hippo V2

Hippo is an easy, fast, lightweight server engine which supports gracefully shutdown.

- powered by Go, Logrus

![Hippo](doc/hippo.png)


## Struct

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
    LogDir: ".",
}
engine := hippo.NewEngine(&SimpleServer{}, config)
if err := engine.Start(); err != nil {
    panic(err)
}
```