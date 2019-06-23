package engine

type Option interface {
	Debug() bool
	Name() string
	Version() string
	Description() string
}

type Server interface {
	Start() error
	SetEngine(*Engine)
}

type Engine struct {
	name        string
	description string
	version     string
	debug       bool
	Option      Option
	server      Server

	//	//Config      map[string]string
	//	debug      bool
	//	//cpuCount    int
	//	//processName string
	//	workingDir  string
	//	//TempDir     string
	//	//logOutput   int // 0: STDOUT, 1: File
	//	//LogPrefix string
	//	//DB          *sql.DB
	//	//TimeZone    *time.Location
}

//func NewEngine(option obj.Option) *Engine {
//}
//
//
func NewEngine(option Option) *Engine {
	e := Engine{
		Option: option,
		//name:        option.Name(),
		//description: option.Version(),
		//version:     option.Version(),
		//debug:       option.Debug(),
	}
	//e.LogPrefix = "[" + e.processName + "] "
	//abs, _ := filepath.Abs(os.Args[0])
	//e.ProcessDir = filepath.Dir(abs)
	return &e
}

func (e *Engine) Start(server Server) error {
	server.SetEngine(e)
	err := server.Start()
	if err != nil {
		panic(err)
	}

	return nil
}
