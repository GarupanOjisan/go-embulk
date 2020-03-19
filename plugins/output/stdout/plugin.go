package stdout

import "os"

type Stdout struct {
	out *os.File
}

func NewStdout() *Stdout {
	return &Stdout{}
}

func (s *Stdout) Write(p []byte) (n int, err error) {
	return s.out.Write(p)
}

func (s *Stdout) Init() error {
	s.out = os.Stdout
	return nil
}

func (s *Stdout) Finalize() error {
	return nil
}
