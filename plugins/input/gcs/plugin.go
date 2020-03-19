package gcs

import (
	"context"

	"cloud.google.com/go/storage"
	"google.golang.org/api/option"
)

type GCS struct {
	reader *storage.Reader

	Ctx                context.Context
	ProjectID          string
	CredentialFilePath string
	BucketName         string
	ObjectName         string
}

func (g *GCS) Read(p []byte) (n int, err error) {
	return g.reader.Read(p)
}

func (g *GCS) Init() error {
	var client *storage.Client
	var err error
	if g.CredentialFilePath != "" {
		client, err = storage.NewClient(g.Ctx, option.WithCredentialsFile(g.CredentialFilePath))
		if err != nil {
			return err
		}
		return nil
	}

	client, err = storage.NewClient(g.Ctx)
	if err != nil {
		return err
	}

	r, err := client.Bucket(g.BucketName).Object(g.ObjectName).NewReader(g.Ctx)
	if err != nil {
		return err
	}
	g.reader = r

	return nil
}

func (g *GCS) Finalize() error {
	return g.reader.Close()
}
