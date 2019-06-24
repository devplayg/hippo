package main

import (
	"fmt"
	"github.com/devplayg/hippo/classifier"
	"github.com/devplayg/hippo/engine"
	"github.com/devplayg/hippo/server"
	"github.com/spf13/pflag"
	"os"
	"runtime"
)

const (
	appName        = "hippo"
	appDescription = "Hippo Server"
	appVersion     = "1.0"
)

var (
	option *classifier.Option
)

func main() {
	engine := engine.NewEngine(option)
	err := engine.Start(server.NewClassifier())
	if err != nil {
		panic(err)
	}
}

func init() {
	runtime.GOMAXPROCS(runtime.NumCPU())

	// Get flag set
	fs := pflag.NewFlagSet(appName, pflag.ContinueOnError)

	// Options
	debug := fs.Bool("debug", false, "Debug")
	cpu := fs.Uint8P("cpu", "c", 0, "CPU Count")

	// Usage
	fs.Usage = func() {
		fmt.Printf("%s v%s\n", appDescription, appVersion)
		fs.PrintDefaults()
	}
	_ = fs.Parse(os.Args[1:])

	// Set options
	option = classifier.NewOption(appName, appDescription, appVersion, *debug)

	// Number of logical CPUs usable
	runtime.GOMAXPROCS(int(*cpu))
}
