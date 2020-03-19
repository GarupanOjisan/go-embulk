package main

import (
	"context"
	"fmt"
	"log"

	"github.com/garupanojisan/go-embulk"
	"github.com/garupanojisan/go-embulk/plugins/input/mysql"
	"github.com/garupanojisan/go-embulk/plugins/output/gcs"
)

func main() {
	ctx := context.Background()
	g := &go_embulk.GoEmbulk{
		Input: mysql.NewMySQLInput(fmt.Sprintf("test:test@tcp(localhost)/test"), "SELECT * FROM test"),
		Outputs: []go_embulk.Output{
			&gcs.GCS{
				Ctx:                ctx,
				ProjectID:          "test",
				CredentialFilePath: "", // use GOOGLE_APPLICATION_CREDENTIALS
				BucketName:         "test",
				ObjectName:         "example/test.csv",
			},
		},
	}

	if err := g.Run(); err != nil {
		log.Fatal(err)
	}
}
