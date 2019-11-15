package hippo

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"github.com/spf13/pflag"
	"os"
	"os/signal"
	"path/filepath"
	"syscall"
)

func init() {
	logrus.SetFormatter(&logrus.JSONFormatter{})
}

func WaitForSignals() {
	signalCh := make(chan os.Signal, 1)
	signal.Notify(signalCh, os.Interrupt, syscall.SIGTERM)
	select {
	case <-signalCh:
		logrus.Info("Signal received, shutting down...")
	}
}

func Usage(fs *pflag.FlagSet, description, version string) func() {
	return func() {
		fmt.Printf("%s v%s\n", description, version)
		fs.PrintDefaults()
		os.Exit(1)
	}
}

func InitLogger(dir, name string, debug, verbose bool) error {
	logrus.SetFormatter(&logrus.TextFormatter{
		ForceColors:   true,
		DisableColors: true,
	})

	if debug {
		logrus.SetLevel(logrus.DebugLevel)
	}

	if verbose {
		logrus.SetOutput(os.Stdout)
		return nil
	}

	workingDir, err := filepath.Abs(os.Args[0])
	if err != nil {
		return err
	}

	logDir := filepath.Dir(workingDir)
	if len(dir) > 0 {
		logDir = dir
	}

	if err := EnsureDir(logDir); err != nil {
		return err
	}

	logFile := filepath.Join(logDir, filepath.Base(name)+".log")
	file, err := os.OpenFile(logFile, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}

	logrus.SetOutput(file)
	return nil
}

func EnsureDir(dir string) error {
	if _, err := os.Stat(dir); os.IsExist(err) {
		return nil
	}

	return os.MkdirAll(dir, os.ModePerm)
}
