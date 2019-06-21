package lib

import (
	"fmt"
	"os"
	"os/signal"
	"path/filepath"
	"strings"
	"syscall"
)

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
