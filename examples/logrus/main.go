//package main

/**
package main

import (
	"github.com/devplayg/hippo/v3"
	"github.com/sirupsen/logrus"
	"os"
)

func main() {
	config := hippo.Config{
		Logger: createLogger("server.log", true, true),
	}

	hippo := hippo.NewHippo(&Server{}, &config)
	if err := hippo.Start(); err != nil {
		panic(err)
	}
}

type Server struct {
	hippo.Launcher // DO NOT REMOVE; Launcher links server and hippo each other.
	Log            *logrus.Logger
}

func (s *Server) Start() error {
	s.Log = s.Hippo.Config.Logger.(*logrus.Logger)
	s.Log.Info("server has been started")
	return nil
}

func (s *Server) Stop() error {
	s.Log.Info("server has been stopped")
	return nil
}

func createLogFile(path string) *os.File {
	file, err := os.OpenFile(path, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		panic(err)
	}
	return file
}

func createLogger(path string, debug, verbose bool) *logrus.Logger {
	file := createLogFile(path)

	logger := logrus.New()
	logger.SetFormatter(&logrus.TextFormatter{
		ForceColors:   true,
		FullTimestamp: true,
	})

	if debug {
		logger.SetLevel(logrus.DebugLevel)
	}

	if verbose {
		logger.SetOutput(os.Stdout)
		return logger
	}

	logger.SetFormatter(&logrus.TextFormatter{
		DisableColors: true,
	})

	logger.SetOutput(file)
	return logger
}
*/
