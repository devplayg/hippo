package hippo

import (
	"fmt"
	"github.com/spf13/pflag"
	"os"
)

//func init() {
//	logrus.SetFormatter(&logrus.JSONFormatter{})
//}

//func WaitForSignals() {
//	signalCh := make(chan os.Signal, 1)
//	signal.Notify(signalCh, os.Interrupt, syscall.SIGTERM)
//	select {
//	case <-signalCh:
//		logrus.Info("Signal received, shutting down...")
//	}
//}

func Usage(fs *pflag.FlagSet, description, version string) func() {
	return func() {
		fmt.Printf("%s v%s\n", description, version)
		fs.PrintDefaults()
		os.Exit(1)
	}
}

func EnsureDir(dir string) error {
	if _, err := os.Stat(dir); os.IsExist(err) {
		return nil
	}

	return os.MkdirAll(dir, os.ModePerm)
}

func NewDefaultConfig(processName string) *Config {
	return &Config{
		Name:        processName,
		DisplayName: processName,
		Description: processName,
		Version:     "unknown",
		LogDir:      "",
		Debug:       false,
		Trace:       false,
	}
}
