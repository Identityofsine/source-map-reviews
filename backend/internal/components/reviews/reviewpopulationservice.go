package reviews

import (
	"fmt"

	"github.com/identityofsine/fofx-go-gin-api-template/internal/components/images"
	"github.com/identityofsine/fofx-go-gin-api-template/internal/constants/exception"
	"github.com/identityofsine/fofx-go-gin-api-template/internal/repository"
	"github.com/identityofsine/fofx-go-gin-api-template/internal/types/routeexception"
	"github.com/identityofsine/fofx-go-gin-api-template/pkg/db/dbmapper"
	"github.com/identityofsine/fofx-go-gin-api-template/pkg/storedlogs"
)

// PopulateReviewsWithImages efficiently populates multiple reviews with their images using bulk queries
func PopulateReviewsWithImages(reviews []MapReview) routeexception.RouteError {
	if len(reviews) == 0 {
		return nil
	}

	// Extract review IDs
	reviewIds := make([]int64, len(reviews))
	for i, review := range reviews {
		reviewIds[i] = review.MapReviewID
	}

	// Get all review images in bulk
	reviewImageMap, err := getReviewImagesGroupedByReviewId(reviewIds)
	if err != nil {
		return err
	}

	// If no images found, that's not an error - just return the reviews as-is
	if len(reviewImageMap) == 0 {
		storedlogs.LogInfo("No images found for any reviews")
		return nil
	}

	// Extract all unique image IDs
	var allImageIds []int64
	for _, reviewImages := range reviewImageMap {
		for _, reviewImage := range reviewImages {
			allImageIds = append(allImageIds, reviewImage.ImageId)
		}
	}

	// Remove duplicates
	allImageIds = uniqueInt64Slice(allImageIds)

	// Get all images in bulk
	imageMap, err := getImagesMapById(allImageIds)
	if err != nil {
		return err
	}

	// Populate each review with its images
	for i := range reviews {
		reviewId := reviews[i].MapReviewID
		reviewImages, exists := reviewImageMap[reviewId]

		if !exists || len(reviewImages) == 0 {
			reviews[i].Images = []MapReviewImage{}
			continue
		}

		// Populate each review image with the actual image data
		populatedImages := make([]MapReviewImage, 0, len(reviewImages))
		for _, reviewImage := range reviewImages {
			if image, imageExists := imageMap[reviewImage.ImageId]; imageExists {
				reviewImage.Image = image
				populatedImages = append(populatedImages, reviewImage)
			} else {
				storedlogs.LogWarn(fmt.Sprintf("Image ID %d not found for review %d", reviewImage.ImageId, reviewId))
			}
		}

		reviews[i].Images = populatedImages
	}

	return nil
}

// PopulateReviewWithImages populates a single review with its images
func PopulateReviewWithImages(review *MapReview) routeexception.RouteError {
	if review == nil {
		return nil
	}

	reviews := []MapReview{*review}
	err := PopulateReviewsWithImages(reviews)
	if err != nil {
		return err
	}

	*review = reviews[0]
	return nil
}

// getReviewImagesGroupedByReviewId gets review images grouped by review ID
func getReviewImagesGroupedByReviewId(reviewIds []int64) (map[int64][]MapReviewImage, routeexception.RouteError) {
	if len(reviewIds) == 0 {
		return make(map[int64][]MapReviewImage), nil
	}

	// Get review images from database
	dbReviewImages, err := repository.GetMapReviewImagesByReviewIds(reviewIds)
	if err != nil {
		if err.Code == exception.CODE_RESOURCE_NOT_FOUND {
			return make(map[int64][]MapReviewImage), nil
		}
		return nil, routeexception.NewRouteError(
			err,
			"Failed to get review images by review IDs",
			"get-review-images-failed",
			err.Code,
		)
	}

	// Map to domain objects
	reviewImages := dbmapper.MapAllDbFields[repository.MapReviewImageDB, MapReviewImage](dbReviewImages)
	if reviewImages == nil {
		return make(map[int64][]MapReviewImage), nil
	}

	// Group by review ID
	reviewImageMap := make(map[int64][]MapReviewImage)
	for _, reviewImage := range *reviewImages {
		reviewImageMap[reviewImage.MapReviewId] = append(reviewImageMap[reviewImage.MapReviewId], reviewImage)
	}

	return reviewImageMap, nil
}

// getImagesMapById gets images mapped by their ID
func getImagesMapById(imageIds []int64) (map[int64]images.Image, routeexception.RouteError) {
	if len(imageIds) == 0 {
		return make(map[int64]images.Image), nil
	}

	// Get images from database
	dbImages, err := repository.GetImagesByIDs(imageIds)
	if err != nil {
		return nil, routeexception.NewRouteError(
			err,
			"Failed to get images by IDs",
			"get-images-by-ids-failed",
			err.Code,
		)
	}

	// Map to domain objects
	imgs := dbmapper.MapAllDbFields[repository.ImageDB, images.Image](dbImages)
	if imgs == nil {
		return make(map[int64]images.Image), nil
	}

	// Create map for quick lookup
	imageMap := make(map[int64]images.Image)
	for _, img := range *imgs {
		imageMap[img.ImageID] = img
	}

	return imageMap, nil
}

// uniqueInt64Slice removes duplicates from a slice of int64
func uniqueInt64Slice(slice []int64) []int64 {
	if len(slice) == 0 {
		return slice
	}

	seen := make(map[int64]bool)
	result := make([]int64, 0, len(slice))

	for _, item := range slice {
		if !seen[item] {
			seen[item] = true
			result = append(result, item)
		}
	}

	return result
}
