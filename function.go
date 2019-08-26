package hippo

import (
	"fmt"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/pflag"
	"os"
	"os/signal"
	"syscall"
)

func init() {
	log.SetFormatter(&log.JSONFormatter{})
}

//func GetProcessName() string {
//	return strings.TrimSuffix(filepath.Base(os.Args[0]), filepath.Ext(os.Args[0]))
//}

func WaitForSignals() {
	signalCh := make(chan os.Signal, 1)
	signal.Notify(signalCh, os.Interrupt, syscall.SIGTERM)
	select {
	case <-signalCh:
		fmt.Println("Signal received, shutting down...")
	}
}

func Usage(fs *pflag.FlagSet, description, version string) func() {
	return func() {
		fmt.Printf("%s v%s\n", description, version)
		fs.PrintDefaults()
		os.Exit(1)
	}
}
