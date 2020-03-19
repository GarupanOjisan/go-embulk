package main

import (
	"log"

	"github.com/garupanojisan/go-embulk"
	"github.com/garupanojisan/go-embulk/plugins/input/stdin"
	"github.com/garupanojisan/go-embulk/plugins/output/stdout"
)

func main() {
	g := &go_embulk.GoEmbulk{
		Input: &stdin.Stdin{},
		Outputs: []go_embulk.Output{
			stdout.NewStdout(),
		},
	}

	if err := g.Run(); err != nil {
		log.Fatal(err)
	}
}
