package hippo

import (
	"fmt"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/pflag"
	"os"
	"os/signal"
	"path/filepath"
	"strings"
	"syscall"
)

func init() {
	log.SetFormatter(&log.JSONFormatter{})
}

func GetProcessName() string {
	return strings.TrimSuffix(filepath.Base(os.Args[0]), filepath.Ext(os.Args[0]))
}

func WaitForSignals() {
	signalCh := make(chan os.Signal, 1)
	signal.Notify(signalCh, os.Interrupt, syscall.SIGTERM)
	select {
	case <-signalCh:
		fmt.Println("Signal received, shutting down...")
	}
}

func drainError(errChan <-chan error) {
	for {
		select {
		case err := <-errChan:
			if err != nil {
				log.Error(err)
			}
		}
	}
}

func Usage(fs *pflag.FlagSet, description, version string) func() {
	return func() {
		fmt.Printf("%s v%s\n", description, version)
		fs.PrintDefaults()
	}
}

func initLogger(path string, debug bool) error {
	file, err := os.OpenFile(path, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	log.SetOutput(file)

	if debug {
		log.SetLevel(log.DebugLevel)
		log.Debug("debug mode")
	}
	return nil
}
