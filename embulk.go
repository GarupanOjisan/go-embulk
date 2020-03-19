package go_embulk

import (
	"fmt"
	"io/ioutil"
)

const MinRead int = 512

type GoEmbulk struct {
	Input   Input
	Outputs []Output
}

func (g *GoEmbulk) Run() error {
	if err := g.Input.Init(); err != nil {
		return err
	}

	buffer, err := ioutil.ReadAll(g.Input)
	if err != nil {
		return fmt.Errorf("error reading Input.Read() %v", err.Error())
	}
	if err := g.Input.Finalize(); err != nil {
		return err
	}

	for i, out := range g.Outputs {
		if err := out.Init(); err != nil {
			return err
		}
		n, err := out.Write(buffer)
		if err != nil {
			return fmt.Errorf("error writing Output[%d] %v", i, err.Error())
		}
		if n != len(buffer) {
			return fmt.Errorf("error wrote %d bytes, but expects %d bytes", n, len(buffer))
		}
		if err := out.Finalize(); err != nil {
			return err
		}
	}
	return nil
}
