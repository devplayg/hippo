package server

import "io"

type VirtualReader struct {
	io.Reader
}

func NewVirtualReader() *VirtualReader {
	return &VirtualReader{}
}

func (r *VirtualReader) Start() error {
	return nil
}

func (r *VirtualReader) Read(p []byte) (n int, err error) {
	println(string(p))
	return 0, nil
}
