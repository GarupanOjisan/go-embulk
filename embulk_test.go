package go_embulk

import (
	"bytes"
	"io"
	"testing"
)

func TestGoEmbulk_Run(t *testing.T) {
	type fields struct {
		Source       io.Reader
		Destinations []io.Writer
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		{
			name: "Success",
			fields: fields{
				Source: bytes.NewBufferString("test"),
				Destinations: []io.Writer{
					func() io.Writer {
						return &bytes.Buffer{}
					}(),
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := &GoEmbulk{
				Source:       tt.fields.Source,
				Destinations: tt.fields.Destinations,
			}
			if err := g.Run(); (err != nil) != tt.wantErr {
				t.Errorf("Run() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
