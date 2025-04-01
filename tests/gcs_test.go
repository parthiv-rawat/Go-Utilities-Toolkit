package tests

import (
	"go_utilities_toolkit/utils/gcs"
	"testing"
)

// NOTE: This test requires GCP credentials to be properly configured in the environment.
func TestUploadToGCS(t *testing.T) {
	bucket := "your-test-bucket"
	object := "test-object.txt"
	data := []byte("Hello from GCS")
	err := gcs.UploadToGCS(bucket, object, data)
	if err != nil {
		t.Errorf("GCS upload failed: %v", err)
	}
}
