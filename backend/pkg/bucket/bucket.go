package bucket

import (
	"os"
	"time"

	"github.com/identityofsine/fofx-go-gin-api-template/pkg/config"
)

// This bucket package is responsible for retrieving and uploading files to a bucket.

func GetFile(path string) ([]byte, error) {

	bucket := config.GetBucketConfig()
	file, err := os.ReadFile(bucket.BucketPath + "/" + path)
	if err != nil {
		return nil, err
	}

	return file, nil
}

func UploadFile(fileExt string, fileData *[]byte) (string, error) {
	bucket := config.GetBucketConfig()
	fileName := generateFileName(fileExt)
	filePath := bucket.BucketPath + "/" + fileName

	// Ensure the directory exists
	if err := os.MkdirAll(bucket.BucketPath, os.ModePerm); err != nil {
		return "", err
	}

	// Write the file
	if err := os.WriteFile(filePath, *fileData, 0644); err != nil {
		return "", err
	}

	return fileName, nil
}

func generateFileName(fileExt string) string {
	currentTime := time.Now()
	return currentTime.Format("20060102150405") + "." + fileExt
}

// GetFiles checks if multiple files exist and returns a map of filename to file existence
func GetFiles(paths []string) map[string]bool {
	bucket := config.GetBucketConfig()
	fileExists := make(map[string]bool)

	for _, path := range paths {
		filePath := bucket.BucketPath + "/" + path
		if _, err := os.Stat(filePath); err == nil {
			fileExists[path] = true
		} else {
			fileExists[path] = false
		}
	}

	return fileExists
}
