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

func (c *Classifier) Start() error {
	option := c.engine.Option.(*classifier.Option)
	println(option.Storage)
	return nil
}

func (c *Classifier) SetEngine(e *engine.Engine) {
	c.engine = e
}
