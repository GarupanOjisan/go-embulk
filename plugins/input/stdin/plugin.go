package stdin

import "os"

type Stdin struct {
	in *os.File
}

func (s *Stdin) Read(p []byte) (n int, err error) {
	return s.in.Read(p)
}

func (s *Stdin) Init() error {
	s.in = os.Stdin
	return nil
}

func (s *Stdin) Finalize() error {
	return nil
}
