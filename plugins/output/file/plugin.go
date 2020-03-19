package file

import (
	"os"
)

type File struct {
	out      *os.File
	FilePath string
}

func (f *File) Write(p []byte) (n int, err error) {
	return f.out.Write(p)
}

func (f *File) Init() error {
	file, err := os.Create(f.FilePath)
	if err != nil {
		return err
	}
	f.out = file
	return nil
}

func (f *File) Finalize() error {
	f.out.Close()
	return nil
}
