package engine

import (
	"github.com/devplayg/hippo/lib"
	"github.com/devplayg/hippo/obj"
	log "github.com/sirupsen/logrus"
	"os"
	"path/filepath"
)

func init() {
	println("init from engine")
}

type Engine struct {
	//Config      map[string]string
	appName    string
	appVersion string
	debug      bool
	//cpuCount    int
	processName string
	ProcessDir  string
	//TempDir     string
	//logOutput   int // 0: STDOUT, 1: File
	LogPrefix string
	//DB          *sql.DB
	//TimeZone    *time.Location
}

func NewEngine(option *obj.Option) *Engine {
	e := Engine{
		appName:     option.AppName,
		appVersion:  option.AppVersion,
		debug:       option.Debug,
		processName: lib.GetProcessName(),
	}
	e.LogPrefix = "[" + e.processName + "] "
	abs, _ := filepath.Abs(os.Args[0])
	e.ProcessDir = filepath.Dir(abs)
	return &e
}

func (e *Engine) Start() error {
	log.Infof("%s started", e.appName)
	return nil
}
