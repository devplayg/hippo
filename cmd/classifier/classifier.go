package main

import (
	"fmt"
	"github.com/devplayg/hippo"
	"github.com/devplayg/hippo/classifier"
	"github.com/spf13/pflag"
	"os"
	"path/filepath"
	"runtime"
)

const (
	appName        = "classifier"
	appDescription = "Data Classifier"
	appVersion     = "1.0"
)

var (
	option *classifier.Option
)

func main() {
	engine := hippo.NewEngine(option)
	err := engine.Start(classifier.NewClassifier())
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
	cpu := fs.IntP("cpu", "c", 0, "CPU Count")
	dir := fs.StringP("dir", "d", "", "Source directory (required)")
	storage := fs.StringP("storage", "s", "/storage", "Storage")

	// Usage
	fs.Usage = func() {
		fmt.Printf("%s v%s\n", appDescription, appVersion)
		fs.PrintDefaults()
	}
	_ = fs.Parse(os.Args[1:])

	if len(*dir) < 1 {
		fs.Usage()
		os.Exit(1)
	}

	abs, _ := filepath.Abs(*dir)
	println(abs)

	// Set options
	option = classifier.NewOption(appName, appDescription, appVersion, *debug)
	option.Dir = *dir
	option.Storage = *storage

	err := option.Validate()
	if err != nil {
		panic(err)
	}

	// Number of logical CPUs usable
	runtime.GOMAXPROCS(*cpu)
}
