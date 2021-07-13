// Package hippo is an easy, fast, lightweight server framework.
package hippo

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"
	"path/filepath"
	"strings"
	"syscall"
	"time"
)

// Hippo helps servers start up safely and shut down gracefully..
type Hippo struct {
	workingDir  string
	Config      *Config
	processName string
	server      Server
	started     time.Time
	ctx         context.Context
	cancel      context.CancelFunc
	errChan     chan error
	log         StdLogger
}

// NewHippo allocates a new server to hippo.
func NewHippo(server Server, config *Config) *Hippo {
	ctx, cancel := context.WithCancel(context.Background())
	return &Hippo{
		processName: strings.TrimSuffix(filepath.Base(os.Args[0]), filepath.Ext(os.Args[0])),
		server:      server,
		Config:      config,
		started:     time.Now(),
		ctx:         ctx,
		cancel:      cancel,
		errChan:     make(chan error),
	}
}

func (e *Hippo) init() error {
	if err := e.initConfig(); err != nil {
		return err
	}
	if err := e.initDirs(); err != nil {
		return err
	}
	if err := e.initLogger(); err != nil {
		return err
	}
	return nil
}

func (e *Hippo) initConfig() error {
	if e.Config == nil {
		e.Config = newDefaultConfig(e.processName)
		return nil
	}

	if len(strings.TrimSpace(e.Config.Name)) < 1 {
		e.Config.Name = e.processName
	}
	if len(strings.TrimSpace(e.Config.Description)) < 1 {
		e.Config.Description = e.processName
	}

	return nil
}

func (e *Hippo) initDirs() error {
	workingDir, err := filepath.Abs(os.Args[0])
	if err != nil {
		return err
	}
	e.workingDir = filepath.Dir(workingDir)
	return nil
}

func (e *Hippo) initLogger() error {
	if e.Config.Logger == nil {
		e.log = log.New(os.Stdout, "", log.LstdFlags)
		return nil
	}
	return nil
}

// Start starts server and opens error channel.
func (e *Hippo) Start() error {
	if err := e.init(); err != nil {
		return fmt.Errorf("failed to initialize hippo: %w", err)
	}

	if err := e.server.initLauncher(e); err != nil {
		return fmt.Errorf("failed to initialize launcher: %w", err)
	}

	done := make(chan bool)
	go func() {
		e.log.Println("hippo has been started")
		e.errChan <- e.server.Start()
		close(done)
	}()
	e.waitForSignals()
	<-done

	if err := e.stop(); err != nil {
		return err
	}
	return nil
}

// Stop stops hippo.
func (e *Hippo) stop() error {
	if err := e.server.Stop(); err != nil {
		panic(err)
		return err
	}
	// e.log.Println("hippo has been stopped")
	return nil
}

func (e *Hippo) waitForSignals() {
	ch := make(chan os.Signal)
	signal.Notify(ch, os.Interrupt, syscall.SIGTERM)
	for {
		select {
		case err := <-e.errChan:
			if err != nil {
				e.log.Println(fmt.Errorf("an error occurred while the server was running: %w", err))
			}
			return
		case <-ch:
			e.log.Println("received signal, shutting down..")
			e.cancel()
		}
	}
}

func (e *Hippo) context() context.Context {
	return e.ctx
}

// Path returns absolute directory
func (e *Hippo) Path(path string) string {
	if filepath.IsAbs(path) {
		return path
	}

	return filepath.ToSlash(filepath.Join(e.workingDir, path))
}

func (e *Hippo) Debug() bool {
	return e.Config.Debug
}
