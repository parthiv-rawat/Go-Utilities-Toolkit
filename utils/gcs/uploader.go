package gcs

import (
	"context"

	"cloud.google.com/go/storage"
)

func UploadToGCS(bucketName, objectName string, data []byte) error {
	ctx := context.Background()
	client, err := storage.NewClient(ctx)
	if err != nil {
		return err
	}
	defer client.Close()

	wc := client.Bucket(bucketName).Object(objectName).NewWriter(ctx)
	defer wc.Close()

	_, err = wc.Write(data)
	return err
}
