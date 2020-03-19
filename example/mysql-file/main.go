package main

import (
	"fmt"
	"github.com/garupanojisan/go-embulk/plugins/output/file"
	"log"

	"github.com/garupanojisan/go-embulk"
	"github.com/garupanojisan/go-embulk/plugins/input/mysql"
)

func main() {
	g := &go_embulk.GoEmbulk{
		Input: mysql.NewMySQLInput(fmt.Sprintf("test:test@tcp(localhost)/test"), "SELECT * FROM test"),
		Outputs: []go_embulk.Output{
			&file.File{FilePath: "out.csv"},
		},
	}

	if err := g.Run(); err != nil {
		log.Fatal(err)
	}
}
