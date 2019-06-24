package server

import (
	"github.com/devplayg/hippo/classifier"
	"github.com/devplayg/hippo/engine"
)

type Classifier struct {
	engine *engine.Engine
}

func NewClassifier() *Classifier {
	return &Classifier{}
}

func (c *Classifier) Start(engine *engine.Engine) error {
	c.engine = engine
	option := c.engine.Option.(*classifier.Option)
	println(option.Storage)
	println(option.Dir)
	println(c.engine.WorkingDir)
	return nil
}

//func (c *Classifier) SetEngine(e *engine.Engine) {
//	c.engine = e
//}
