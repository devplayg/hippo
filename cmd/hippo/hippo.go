package main

import (
	"fmt"
	"github.com/devplayg/hippo/engine"
	"github.com/spf13/pflag"
	"os"
	"runtime"
)

const (
	appName        = "hippo"
	appDescription = "Hippo Server"
	appVersion     = "0.1"
)

var (
	option *engine.Option
)

func main() {

	//engine := engine.NewEngine(option)
	//err := engine.Start(server.NewHippoServer())
	//if err != nil {
	//	panic(err)
	//}
}

func init() {
	// Get flag set
	fs := pflag.NewFlagSet(appName, pflag.ContinueOnError)

	// Options
	//debug := fs.BoolP("debug", "d", false, "Debug")
	cpu := fs.Uint8P("cpu", "c", 0, "CPU Count")

	// Usage
	fs.Usage = func() {
		fmt.Printf("%s v%s\n\n", appDescription, appVersion)
		fs.PrintDefaults()
	}
	_ = fs.Parse(os.Args[1:])

	// Set options

	//option = &engine.Option{
	//	Name:        appName,
	//	Version:     appVersion,
	//	Description: appDescription,
	//	Debug:       *debug,
	//}

	// Number of logical CPUs usable
	runtime.GOMAXPROCS(int(*cpu))
}
