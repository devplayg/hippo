package hippo

type Server interface {
	Start(engine *Engine) error
}
