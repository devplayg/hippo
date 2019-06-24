package classifier

import (
	"github.com/devplayg/hippo"
)

type Classifier struct {
	engine *hippo.Engine
}

func NewClassifier() *Classifier {
	return &Classifier{}
}

func (c *Classifier) Start(engine *hippo.Engine) error {
	c.engine = engine
	option := c.engine.Option.(*Option)
	println(option.Storage)
	println(option.Dir)
	println(c.engine.WorkingDir)
	return nil
}

//func (c *Classifier) SetEngine(e *engine.Engine) {
//	c.engine = e
//}
