package images

import (
	"fmt"

	"github.com/identityofsine/fofx-go-gin-api-template/internal/constants/exception"
	"github.com/identityofsine/fofx-go-gin-api-template/internal/repository"
	"github.com/identityofsine/fofx-go-gin-api-template/internal/types/routeexception"
	"github.com/identityofsine/fofx-go-gin-api-template/pkg/bucket"
	"github.com/identityofsine/fofx-go-gin-api-template/pkg/db/dbmapper"
)

func SaveImage(request *SaveImageRequest, fileData []byte) (*Image, routeexception.RouteError) {

	if request == nil {
		return nil, exception.BadRequest
	}

	if ok, err := verifyIsImage(fileData); !ok || err != nil {

		if err != nil {
			return nil, routeexception.NewRouteError(
				err,
				"Failed to verify image",
				"verify-image-failed",
				500,
			)
		}

		return nil, routeexception.NewRouteError(
			nil,
			"Invalid image save request",
			"invalid-image-save-request",
			exception.CODE_BAD_REQUEST,
		)

	}

	sysFile, err := bucket.UploadFile(
		request.FileExt,
		&fileData,
	)
	if err != nil {
		return nil, routeexception.NewRouteError(
			err,
			"Failed to upload file",
			"upload-file-failed",
			exception.CODE_INTERNAL_SERVER_ERROR,
		)
	}

	// Here you would implement the logic to save the image.
	// This is a placeholder implementation.
	imageTemp := Image{
		ImagePath: sysFile,
		Caption:   request.Caption,
	}
	imageDb := dbmapper.MapDbFields[Image, repository.ImageDB](imageTemp)
	if imageDb == nil {
		return nil, routeexception.NewRouteError(
			nil,
			"Failed to map image fields",
			"map-image-fields-failed",
			exception.CODE_INTERNAL_SERVER_ERROR,
		)
	}

	img, derr := repository.SaveImageDb(*imageDb)
	if derr != nil {
		fmt.Println("Error saving image to database:", derr)
		return nil, routeexception.NewRouteError(
			derr,
			"Failed to save image to database",
			"save-image-db-failed",
			derr.Code,
		)
	}

	// Simulate saving the image and returning the saved image object.
	return dbmapper.MapDbFields[repository.ImageDB, Image](*img), nil

}

func verifyIsImage(fileData []byte) (bool, error) {
	// Here you would implement the logic to verify if the file data is a valid image.
	// This is a placeholder implementation.
	if len(fileData) == 0 {
		return false, nil // Empty data is not a valid image
	}

	// You can add more checks here to verify the image format, size, etc.
	return true, nil
}
