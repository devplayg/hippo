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
	WorkingDir string
	Config     *Config

	processName string
	server      Server
	started     time.Time
	Ctx         context.Context
	cancel      context.CancelFunc
	Done        chan bool
}

// NewEngine allocates a new server to engine.
func NewEngine(server Server, config *Config) *Engine {
	return &Engine{
		processName: config.Name,
		server:      server,
		Config:      config,
		started:     time.Now(),
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
	e.WorkingDir = filepath.Dir(workingDir)
	if err != nil {
		return err
	}
	return nil
}

func (e *Engine) initLogger() error {
	if e.Config.Debug {
		logrus.SetLevel(logrus.DebugLevel)
	}

	if e.Config.Trace {
		logrus.SetLevel(logrus.DebugLevel)
	}

	logrus.SetFormatter(&logrus.TextFormatter{
		ForceColors:   true,
		DisableColors: true,
	})

	if e.Config.Verbose {
		logrus.SetOutput(os.Stdout)
		return nil
	}

	if len(e.Config.LogDir) > 0 {
		logDir, err := filepath.Abs(e.Config.LogDir)
		if err != nil {
			return fmt.Errorf("invalid log directory: %w", err)
		}
		if err := EnsureDir(logDir); err != nil {
			return fmt.Errorf("failed to create log directory: %w", err)
		}
	} else {

	}

	//workingDir, err := filepath.Abs(os.Args[0])
	//if err != nil {
	//	return err
	//}
	//
	//logDir := filepath.Dir(workingDir)
	//if len(dir) > 0 {
	//	logDir = dir
	//}
	//
	//if err := EnsureDir(logDir); err != nil {
	//	return err
	//}
	//
	//logFile := filepath.Join(logDir, filepath.Base(name)+".log")
	//file, err := os.OpenFile(logFile, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
	//if err != nil {
	//	return err
	//}
	//
	//logrus.SetOutput(file)
}

// Start starts server and opens error channel.
func (e *Engine) Start() error {
	if err := e.init(); err != nil {
		return fmt.Errorf("failed to initialize hippo engine: %w", err)
	}

	logrus.Debug("hippo engine has been started")
	e.server.SetEngine(e)

	if e.Config.IsService {
		e.Ctx, e.cancel = context.WithCancel(context.Background())
		e.Done = make(chan bool)
	}

	if err := e.server.Start(); err != nil {
		return err
	}
	defer e.Stop()

	logrus.Infof("%s has been started", e.Config.Name)
	if e.Config.IsService {
		e.waitForSignals()
	}

	return nil
}

// Stop stops engine.
func (e *Engine) Stop() error {
	if e.Config.IsService {
		e.cancel()
		<-e.Done
		logrus.Debugf("%s has been shutted down gracefully", e.Config.Name)
	}
	if err := e.server.Stop(); err != nil {
		logrus.Error("failed to stop %s", e.Config.Name)
		return err
	}
	logrus.Infof("%s has been stopped", e.Config.Name)
	logrus.Debugf("hippo engine has been stopped (running time: %s)", time.Since(e.started))
	return nil
}

func (e *Engine) waitForSignals() {
	signalCh := make(chan os.Signal, 1)
	signal.Notify(signalCh, os.Interrupt, syscall.SIGTERM)
	select {
	case <-signalCh:
		logrus.Info("Signal received, shutting down...")
	}
}
