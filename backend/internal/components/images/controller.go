package images

import (
	"io"
	"mime"
	"net/http"
	"os"
	"path/filepath"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/identityofsine/fofx-go-gin-api-template/internal/constants/exception"
	"github.com/identityofsine/fofx-go-gin-api-template/pkg/config"
)

func GetImageRoute(c *gin.Context) {
	// Get the path from the request parameters
	path := c.Param("path")

	// Build the full file path
	bucket := config.GetBucketConfig()
	fullPath := filepath.Join(bucket.BucketPath, path)

	// Open file for streaming (this also checks if file exists)
	file, err := os.Open(fullPath)
	if err != nil {
		if os.IsNotExist(err) {
			c.AbortWithStatusJSON(404, exception.ResourceNotFound)
			return
		}
		c.AbortWithStatusJSON(500, gin.H{
			"error":   "Failed to open file",
			"message": err.Error(),
		})
		return
	}
	defer file.Close()

	// Get file info for content length
	fileInfo, err := file.Stat()
	if err != nil {
		c.AbortWithStatusJSON(500, gin.H{
			"error":   "Failed to get file info",
			"message": err.Error(),
		})
		return
	}

	// Detect content type from file extension
	contentType := mime.TypeByExtension(filepath.Ext(fullPath))
	if contentType == "" {
		contentType = "application/octet-stream"
	}

	// Set headers before streaming
	c.Header("Content-Type", contentType)
	c.Header("Content-Length", strconv.FormatInt(fileInfo.Size(), 10))
	c.Header("Cache-Control", "public, max-age=31536000") // Cache for 1 year

	// Stream the file directly
	c.Status(http.StatusOK)
	_, err = io.Copy(c.Writer, file)
	if err != nil {
		// Log error but don't try to send JSON response as headers are already sent
		c.Error(err)
	}
}
