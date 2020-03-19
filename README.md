# Go-Embulk

Implementation [Embulk](https://github.com/embulk/embulk) in Go.

## Usage

```golang
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
```

## Input

Input struct should implements Input interface.

For example,
 
```golang
type ExampleInput struct {
}

func (s *ExampleInput) Read(p []byte) error {
    // do something (e.g. initialization, authentication, etc...)
    
    // set data into p

    // do something
}

func (s *ExampleInput) Init() error {
    return nil
}

func (s *ExampleInput) Finalize() error {
    return nil
}
```

## Output 

Output struct should implements Output interface.

```golang
type ExampleOutput struct {
}

func (s *ExampleOutput) Write(data []byte) (int, error) {
    // do something (e.g. initialization, authentication, etc...)
    
    // write data

    // do something
}

func (s *ExampleOutput) Init() error {
    return nil
}

func (s *ExampleOutput) Finalize() error {
    return nil
}
```