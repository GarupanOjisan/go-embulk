package go_embulk

import (
	"io"
	"testing"
)

func TestGoEmbulk_Run(t *testing.T) {
	type fields struct {
		Source       Input
		Destinations []Output
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		{
			name: "Success",
			fields: fields{
				Source: &TestInput{},
				Destinations: []Output{
					&TestOutput{},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := &GoEmbulk{
				Input:   tt.fields.Source,
				Outputs: tt.fields.Destinations,
			}
			if err := g.Run(); (err != nil) != tt.wantErr {
				t.Errorf("Run() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

type TestInput struct{}

func (t *TestInput) Read(p []byte) (int, error) {
	return 0, io.EOF
}

func (t *TestInput) Init() error {
	return nil
}

func (t *TestInput) Finalize() error {
	return nil
}

type TestOutput struct{}

func (t *TestOutput) Write(data []byte) (int, error) {
	return len(data), nil
}

func (t *TestOutput) Init() error {
	return nil
}

func (t *TestOutput) Finalize() error {
	return nil
}
