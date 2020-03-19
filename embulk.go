package go_embulk

import (
	"fmt"
	"io"
	"io/ioutil"
)

type GoEmbulk struct {
	Input   io.Reader
	Outputs []io.Writer
}

func (g *GoEmbulk) Run() error {
	buffer, err := ioutil.ReadAll(g.Input)
	if err != nil {
		return fmt.Errorf("error reading Input.Read() %v", err.Error())
	}

	for i, dst := range g.Outputs {
		n, err := dst.Write(buffer)
		if err != nil {
			return fmt.Errorf("error writing destination[%d] %v", i, err.Error())
		}
		if n != len(buffer) {
			return fmt.Errorf("error wrote %d bytes, but expects %d bytes", n, len(buffer))
		}
	}
	return nil
}
