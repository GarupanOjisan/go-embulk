package main

import (
	"fmt"
	"log"

	"github.com/garupanojisan/go-embulk"
	"github.com/garupanojisan/go-embulk/plugins/input/mysql"
	"github.com/garupanojisan/go-embulk/plugins/output/stdout"
)

func main() {
	g := &go_embulk.GoEmbulk{
		Input: mysql.NewMySQLInput(fmt.Sprintf("test:test@tcp(localhost)/test"), "SELECT * FROM test"),
		Outputs: []go_embulk.Output{
			stdout.NewStdout(),
		},
	}

	if err := g.Run(); err != nil {
		log.Fatal(err)
	}
}
