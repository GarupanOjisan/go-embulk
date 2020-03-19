# Go-Embulk

Implementation [Embulk](https://github.com/embulk/embulk) in Go.

## Usage



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
type ExampleSource struct {
}

func (s *ExampleSource) Read(p []byte) error {
    // do something (e.g. initialization, authentication, etc...)
    
    // set data into p

    // do something
}
```