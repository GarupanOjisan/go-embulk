package go_embulk

import (
	"fmt"
	"io"
	"io/ioutil"
)

type GoEmbulk struct {
	Source       io.Reader
	Destinations []io.Writer
}

func (g *GoEmbulk) Run() error {
	buffer, err := ioutil.ReadAll(g.Source)
	if err != nil {
		return fmt.Errorf("error reading Source.Read() %v", err.Error())
	}

	for i, dst := range g.Destinations {
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
