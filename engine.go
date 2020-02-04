// Hippo is an easy, fast, lightweight server framework.
package hippo

import (
	"context"
	"fmt"
	"github.com/sirupsen/logrus"
	"os"
	"os/signal"
	"path/filepath"
	"syscall"
	"time"
)

// Engine supports engine framework.
type Engine struct {
	workingDir  string
	Config      *Config
	processName string
	server      Server
	started     time.Time
	ctx         context.Context
	cancel      context.CancelFunc
	errChan     chan error
}

// NewEngine allocates a new server to engine.
func NewEngine(server Server, config *Config) *Engine {
	ctx, cancel := context.WithCancel(context.Background())
	return &Engine{
		processName: config.Name,
		server:      server,
		Config:      config,
		started:     time.Now(),
		ctx:         ctx,
		cancel:      cancel,
		errChan:     make(chan error),
	}
}

func (e *Engine) init() error {
	if err := e.initDirs(); err != nil {
		return err
	}
	if err := e.initLogger(); err != nil {
		return err
	}
	return nil
}

func (e *Engine) initDirs() error {
	workingDir, err := filepath.Abs(os.Args[0])
	e.workingDir = filepath.Dir(workingDir)
	if err != nil {
		return err
	}
	return nil
}

func (e *Engine) initLogger() error {
	logrus.SetFormatter(&logrus.TextFormatter{
		ForceColors:   true,
		DisableColors: true,
	})

	if e.Config.Debug {
		logrus.SetLevel(logrus.DebugLevel)
	}

	if e.Config.Trace {
		logrus.SetLevel(logrus.TraceLevel)
	}

	if e.Config.Verbose {
		logrus.SetOutput(os.Stdout)
		return nil
	}

	logDir := e.workingDir
	if len(e.Config.LogDir) > 0 {
		dir, err := filepath.Abs(e.Config.LogDir)
		if err != nil {
			return fmt.Errorf("invalid log directory: %w", err)
		}
		if err := EnsureDir(logDir); err != nil {
			return fmt.Errorf("failed to create log directory: %w", err)
		}
		logDir = dir
	}
	logPath := filepath.Join(logDir, filepath.Base(e.processName)+".log")
	file, err := os.OpenFile(logPath, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	logrus.SetOutput(file)
	return nil
}

// Start starts server and opens error channel.
func (e *Engine) Start() error {
	if err := e.init(); err != nil {
		return fmt.Errorf("failed to initialize hippo engine: %w", err)
	}

	e.server.SetEngine(e)
	done := make(chan bool)
	go func() {
		defer func() {
			done <- true
		}()
		logrus.Debug("1) engine has been started")
		if err := e.server.Start(); err != nil {
			e.errChan <- err
			return
		}
		logrus.Debug("4) verified server no longer work")
		e.errChan <- nil
	}()
	e.waitForSignals()
	//logrus.Debug("waiting for done signal")
	<-done

	if err := e.server.Stop(); err != nil {
		return err
	}

	logrus.Debug("6) engine has been stopped")
	return nil
}

// Stop stops engine.
func (e *Engine) Stop() error {
	if err := e.server.Stop(); err != nil {
		logrus.Error("failed to stop %s", e.processName)
		return err
	}
	logrus.Debug("8) engine has been stopped")
	return nil
}

func (e *Engine) waitForSignals() {
	ch := make(chan os.Signal)
	signal.Notify(ch, os.Interrupt, syscall.SIGTERM)
	for {
		select {
		case err := <-e.errChan:
			if err != nil {
				logrus.Error(fmt.Errorf("server has stopped unintentionally: %w", err))
			}
			return
		case <-ch:
			logrus.Debug("received signal, shutting down..")
			e.cancel()
		}
	}
}

func (e *Engine) GetContext() context.Context {
	return e.ctx
}

func (e *Engine) GetWorkingDir() string {
	return e.workingDir
}
