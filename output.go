package go_embulk

import "io"

type Output interface {
	io.Writer
	Init() error
	Finalize() error
}
