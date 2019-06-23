package server

import "github.com/devplayg/hippo/engine"

type Hippo struct {
	engine *engine.Engine
	//setEngine *engine.Engine
}

func NewHippoServer() *Hippo {
	return &Hippo{}
}

//func NewHippoEngine(option *obj.Option) Engine {
//	return &HippoEngine{
//		name:        option.Name,
//		description: option.Description,
//		debug:       option.Debug,
//		version:     option.Version,
//	}
//}
//
//type HippoEngine struct {
//	debug       bool
//	name        string
//	description string
//	version     string
//}
//
func (h *Hippo) Start() error {
	return nil
}

func (h *Hippo) SetEngine(e *engine.Engine) {
	h.engine = e
}

//func (h *HippoEngine) Name() string {
//	return h.name
//}
//
//func (h *HippoEngine) Description() string {
//	return h.description
//}
//
//func (h *HippoEngine) Version() string {
//	return h.version
//}
//
//func (h *HippoEngine) Debug() bool {
//	return h.debug
//}
