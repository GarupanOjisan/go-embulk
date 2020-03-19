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
		Source: bytes.NewBufferString("hello go-embulk"),
		Destinations: []io.Writer{
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
type ExampleSource struct {
}

func (s *ExampleSource) Read(p []byte) error {
    // do something (e.g. initialization, authentication, etc...)
    
    // set data into p

    // do something
}
```

## Output 

Output struct should implements io.Writer.

```golang
type ExampleDestination struct {
}

func (s *ExampleDestination) Write(data []byte) (int, error) {
    // do something (e.g. initialization, authentication, etc...)
    
    // write data

    // do something
}
```