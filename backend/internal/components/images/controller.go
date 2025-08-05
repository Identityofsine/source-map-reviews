package images

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"mime"
	"net/http"
	"os"
	"path/filepath"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/identityofsine/fofx-go-gin-api-template/internal/constants/exception"
	"github.com/identityofsine/fofx-go-gin-api-template/internal/types/routeexception"
	"github.com/identityofsine/fofx-go-gin-api-template/pkg/config"
	"github.com/identityofsine/fofx-go-gin-api-template/pkg/storedlogs"
)

func GetImageRoute(c *gin.Context) {
	// Get the path from the request parameters
	path := c.Param("path")

	// Build the full file path
	bucket := config.GetBucketConfig()
	fullPath := filepath.Join(bucket.BucketPath, path)

	// Check if file exists and get file info
	fileInfo, err := os.Stat(fullPath)
	if err != nil {
		if os.IsNotExist(err) {
			c.AbortWithStatusJSON(404, exception.ResourceNotFound)
			return
		}
		c.AbortWithStatusJSON(500, gin.H{
			"error":   "Failed to access file",
			"message": err.Error(),
		})
		return
	}

	// Open file for streaming
	file, err := os.Open(fullPath)
	if err != nil {
		c.AbortWithStatusJSON(500, gin.H{
			"error":   "Failed to open file",
			"message": err.Error(),
		})
		return
	}
	defer file.Close()

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

func SaveImageRoute(c *gin.Context) {

	var form SaveImageForm
	// Bind the form data
	if err := c.ShouldBind(&form); err != nil {
		c.AbortWithStatusJSON(exception.CODE_BAD_REQUEST, gin.H{"error": "Invalid form data"})
		return
	}

	if form.JSON == nil {
		c.AbortWithStatusJSON(exception.CODE_BAD_REQUEST, gin.H{"error": "JSON data is required"})
		return
	}

	var saveImageRequest SaveImageRequest
	saveImageRequest = SaveImageRequest{}

	if err := json.Unmarshal([]byte(*form.JSON), &saveImageRequest); err != nil {
		c.AbortWithStatusJSON(exception.CODE_BAD_REQUEST,
			routeexception.NewRouteError(err, "Invalid JSON data", "invalid-json-data", exception.CODE_BAD_REQUEST))
		return
	}

	if form.File == nil {
		c.AbortWithStatusJSON(exception.CODE_BAD_REQUEST,
			routeexception.NewRouteError(nil, "File is required", "file-required", exception.CODE_BAD_REQUEST),
		)
		return
	}

	var images []Image
	images = make([]Image, 0, len(form.File))

	files := form.File
	for _, fileHeader := range files {
		file, err := fileHeader.Open()
		defer file.Close()
		if err != nil {
			c.AbortWithStatusJSON(exception.CODE_BAD_REQUEST,
				routeexception.NewRouteError(err, "Failed to open file", "file-open-failed", exception.CODE_BAD_REQUEST),
			)
			return
		}
		buf := bytes.NewBuffer(nil)
		if _, err := io.Copy(buf, file); err != nil {
			c.AbortWithStatusJSON(exception.CODE_BAD_REQUEST,
				routeexception.NewRouteError(err, "Failed to read file", "file-read-failed", exception.CODE_BAD_REQUEST),
			)
			return
		}

		img, ierr := SaveImage(&saveImageRequest, buf.Bytes())
		if ierr != nil {
			c.AbortWithStatusJSON(exception.CODE_INTERNAL_SERVER_ERROR,
				routeexception.NewRouteError(err, "Failed to save image", "image-save-failed", exception.CODE_INTERNAL_SERVER_ERROR),
			)
			return
		}

		images = append(images, *img)
		storedlogs.LogInfo(fmt.Sprintf("File %s stored successfully", fileHeader.Filename))
	}

	c.JSON(200, gin.H{
		"message": "File(s) uploaded successfully",
		"images":  images,
	})
}
