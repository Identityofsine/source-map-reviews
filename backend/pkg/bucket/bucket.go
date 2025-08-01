package bucket

import (
	"os"

	"github.com/identityofsine/fofx-go-gin-api-template/pkg/config"
)

// This bucket package is responsible for retrieving and uploading files to a bucket.

func GetFile(path string) ([]byte, error) {

	bucket := config.GetBucketConfig()
	file, err := os.ReadFile(bucket.BucketPath + "/" + path)

	return file, err
}
