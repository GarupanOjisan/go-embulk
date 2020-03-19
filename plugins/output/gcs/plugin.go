package gcs

import (
	"bytes"
	"context"
	"io"

	"cloud.google.com/go/storage"
	"google.golang.org/api/option"
)

type GCS struct {
	client *storage.Client

	Ctx                context.Context
	ProjectID          string
	CredentialFilePath string
	BucketName         string
	ObjectName         string
}

func (g *GCS) Write(p []byte) (n int, err error) {
	bucket := g.client.Bucket(g.BucketName)
	w := bucket.Object(g.ObjectName).NewWriter(g.Ctx)
	defer w.Close()

	src := bytes.NewBuffer(p)
	written, err := io.Copy(w, src)
	if err != nil {
		return 0, err
	}
	return int(written), nil
}

func (g *GCS) Init() error {
	if g.CredentialFilePath != "" {
		client, err := storage.NewClient(g.Ctx, option.WithCredentialsFile(g.CredentialFilePath))
		if err != nil {
			return err
		}
		g.client = client
		return nil
	}

	client, err := storage.NewClient(g.Ctx)
	if err != nil {
		return err
	}
	g.client = client

	return nil
}

func (g *GCS) Finalize() error {
	return nil
}
