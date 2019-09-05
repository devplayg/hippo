package hippo

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"github.com/spf13/pflag"
	"os"
	"os/signal"
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
