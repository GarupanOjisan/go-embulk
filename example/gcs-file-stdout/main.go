package main

import (
	"context"
	in "github.com/garupanojisan/go-embulk/plugins/input/gcs"
	"github.com/garupanojisan/go-embulk/plugins/output/file"
	"log"

	"github.com/garupanojisan/go-embulk"

	"github.com/garupanojisan/go-embulk/plugins/output/stdout"
)

func main() {
	ctx := context.Background()
	g := &go_embulk.GoEmbulk{
		Input: &in.GCS{
			Ctx:                ctx,
			ProjectID:          "test",
			CredentialFilePath: "", // use GOOGLE_APPLICATION_CREDENTIALS
			BucketName:         "test",
			ObjectName:         "example/test.csv",
		},
		Outputs: []go_embulk.Output{
			&file.File{
				FilePath: "out.csv",
			},
			stdout.NewStdout(),
		},
	}

	if err := g.Run(); err != nil {
		log.Fatal(err)
	}
}
