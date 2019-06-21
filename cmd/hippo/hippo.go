package main

import (
	"fmt"
	"github.com/devplayg/hippo/engine"
	"github.com/devplayg/hippo/lib"
	"github.com/devplayg/hippo/obj"
	"github.com/devplayg/hippo/server"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/pflag"
	"os"
)

const (
	appName    = "hippo"
	appVersion = "0.1"
)

var (
	option *obj.Option
)

func main() {
	hippoEngine := engine.NewEngine(option)
	err := hippoEngine.Start()
	if err != nil {
		panic(err)
	}

	hippoReader := server.NewVirtualReader()
	err = hippoReader.Start()
	if err != nil {
		log.Error(err)
	}
	hippoReader.Read([]byte("asdfasdf"))

	lib.WaitForSignals()
}

func init() {
	// Get flag set
	fs := pflag.NewFlagSet("dff", pflag.ContinueOnError)

	// Get arguments
	debug := fs.BoolP("debug", "d", false, "Debug")
	fs.Usage = func() {
		fmt.Printf("Duplicate file finder v%s\n\n", appVersion)
		fs.PrintDefaults()
	}
	_ = fs.Parse(os.Args[1:])

	// Set options
	option = &obj.Option{
		AppName:    appName,
		AppVersion: appVersion,
		Debug:      *debug,
	}

	// CPU 설정
	//runtime.GOMAXPROCS(runtime.NumCPU())

	//dff.InitLogger(*verbose)

	//hippoEngine = engine.NewEngine(appName, *debug)
	//duplicateFileFinder = dff.NewDuplicateFileFinder(*dirs, *minNumOfFilesInFileGroup, *minFileSize, *sortBy, *format)
}

//func printHelp() {
//	fmt.Printf("%s v%s\n\n", appDescription, appVersion)
//	fs.PrintDefaults()
//}
