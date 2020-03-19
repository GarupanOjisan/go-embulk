# Go-Embulk

Implementation [Embulk](https://github.com/embulk/embulk) in Go.

## Usage

```golang
package main

import (
	"bytes"
	"io"
	"log"
	"os"

	"github.com/garupanojisan/go-embulk"
)

func main() {
	g := &go_embulk.GoEmbulk{
		Input: bytes.NewBufferString("hello go-embulk"),
		Outputs: []io.Writer{
			os.Stdout,
		},
	}

	if err := g.Run(); err != nil {
		log.Fatal(err)
	}
}
```

## Input

Input struct should implements io.Reader.

For example,
 
```golang
type ExampleInput struct {
}

func (s *ExampleInput) Read(p []byte) error {
    // do something (e.g. initialization, authentication, etc...)
    
    // set data into p

    // do something
}
```

## Output 

Output struct should implements io.Writer.

```golang
type ExampleOutput struct {
}

func (s *ExampleOutput) Write(data []byte) (int, error) {
    // do something (e.g. initialization, authentication, etc...)
    
    // write data

    // do something
}
```