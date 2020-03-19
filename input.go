package go_embulk

import "io"

type Input interface {
	io.Reader
	Init() error
	Finalize() error
}
