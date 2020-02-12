// Hippo is an easy, fast, lightweight server framework.
package hippo

import (
	"context"
	"fmt"
	"github.com/sirupsen/logrus"
	"os"
	"os/signal"
	"path/filepath"
	"strings"
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
	log         *logrus.Logger
}

// NewEngine allocates a new server to engine.
func NewEngine(server Server, config *Config) *Engine {
	ctx, cancel := context.WithCancel(context.Background())
	return &Engine{
		processName: strings.TrimSuffix(filepath.Base(os.Args[0]), filepath.Ext(os.Args[0])),
		server:      server,
		Config:      config,
		started:     time.Now(),
		ctx:         ctx,
		cancel:      cancel,
		errChan:     make(chan error),
	}
}

func (e *Engine) init() error {
	if err := e.initConfig(); err != nil {
		return err
	}
	if err := e.initDirs(); err != nil {
		return err
	}
	if err := e.initLogger(); err != nil {
		return err
	}
	e.log.Debug("logger has been initialized")
	return nil
}

func (e *Engine) initConfig() error {
	if e.Config == nil {
		e.Config = newDefaultConfig(e.processName)
		return nil
	}

	config := newDefaultConfig(e.processName)
	if len(e.Config.Name) < 1 {
		e.Config.Name = config.Name
	}
	if len(e.Config.DisplayName) < 1 {
		e.Config.DisplayName = config.DisplayName
	}
	if len(e.Config.Description) < 1 {
		e.Config.Description = config.Description
	}

	return nil
}

func (e *Engine) initDirs() error {
	workingDir, err := filepath.Abs(os.Args[0])
	if err != nil {
		return err
	}
	e.workingDir = filepath.Dir(workingDir)
	return nil
}

func (e *Engine) initLogger() error {
	logger := logrus.New()
	defer func() {
		e.log = logger
	}()
	if e.Config.Debug {
		logger.SetLevel(logrus.DebugLevel)
	}
	if e.Config.Trace {
		logger.SetLevel(logrus.TraceLevel)
	}
	if len(e.Config.LogDir) < 1 {
		logger.SetFormatter(&logrus.TextFormatter{
			ForceColors:   true,
			FullTimestamp: true,
		})
		logger.SetOutput(os.Stdout)
		return nil
	}

	logger.SetFormatter(&logrus.TextFormatter{
		DisableColors: true,
	})
	logDir := e.workingDir
	if len(e.Config.LogDir) > 0 {
		dir, err := filepath.Abs(e.Config.LogDir)
		if err != nil {
			return fmt.Errorf("invalid log directory: %w", err)
		}
		if err := ensureDir(logDir); err != nil {
			return fmt.Errorf("failed to create log directory: %w", err)
		}
		logDir = dir
	}
	logPath := filepath.Join(logDir, filepath.Base(e.processName)+".log")
	file, err := os.OpenFile(logPath, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	logger.SetOutput(file)
	return nil
}

// Start starts server and opens error channel.
func (e *Engine) Start() error {
	if err := e.init(); err != nil {
		return fmt.Errorf("failed to initialize hippo engine: %w", err)
	}

	e.log.Trace("engine starting")
	e.server.setEngine(e)
	done := make(chan bool)
	go func() {
		e.log.Debug("engine has been started")
		e.errChan <- e.server.Start()
		done <- true
	}()
	e.waitForSignals()
	<-done

	if err := e.stop(); err != nil {
		return err
	}
	return nil
}

// Stop stops engine.
func (e *Engine) stop() error {
	if err := e.server.Stop(); err != nil {
		e.log.Error("failed to stop %s", e.processName)
		return err
	}
	e.log.Debug("engine has been stopped")
	return nil
}

func (e *Engine) waitForSignals() {
	ch := make(chan os.Signal)
	signal.Notify(ch, os.Interrupt, syscall.SIGTERM)
	for {
		select {
		case err := <-e.errChan:
			if err != nil {
				e.log.Error(fmt.Errorf("server has stopped unintentionally: %w", err))
			}
			return
		case <-ch:
			e.log.Info("received signal, shutting down..")
			e.cancel()
		}
	}
}

func (e *Engine) getContext() context.Context {
	return e.ctx
}

func (e *Engine) GetWorkingDir() string {
	return e.workingDir
}

func ensureDir(dir string) error {
	if _, err := os.Stat(dir); os.IsExist(err) {
		return nil
	}

	return os.MkdirAll(dir, os.ModePerm)
}
