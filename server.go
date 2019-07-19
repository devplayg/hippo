package hippo

type Server interface {
	Start() error
	Stop() error
}
