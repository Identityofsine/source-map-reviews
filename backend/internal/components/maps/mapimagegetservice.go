package maps

import (
	"fmt"

	"github.com/identityofsine/fofx-go-gin-api-template/internal/components/images"
	"github.com/identityofsine/fofx-go-gin-api-template/internal/components/reviews"
	"github.com/identityofsine/fofx-go-gin-api-template/internal/constants/exception"
	"github.com/identityofsine/fofx-go-gin-api-template/internal/types/routeexception"
	"github.com/identityofsine/fofx-go-gin-api-template/pkg/bucket"
	"github.com/identityofsine/fofx-go-gin-api-template/util"
)

func GetMapImagesByMapName(mapNames []string) (map[string]*[]images.Image, routeexception.RouteError) {

	if mapNames == nil || len(mapNames) == 0 {
		return nil, routeexception.NewRouteError(
			nil,
			"Map names cannot be empty",
			"map-names-empty",
			exception.CODE_BAD_REQUEST,
		)
	}

	// Use the new bulk function
	return GetMapImages(mapNames)
}

// GetMapImages efficiently processes multiple map names to get their images
func GetMapImages(mapNames []string) (map[string]*[]images.Image, routeexception.RouteError) {

	if mapNames == nil || len(mapNames) == 0 {
		return nil, routeexception.NewRouteError(
			nil,
			"Map names cannot be empty",
			"map-names-empty",
			exception.CODE_BAD_REQUEST,
		)
	}

	// Check which system images exist in bulk
	systemImagePaths := make([]string, len(mapNames))
	for i, mapName := range mapNames {
		systemImagePaths[i] = fmt.Sprintf("%s.jpg", mapName)
	}
	systemImageExists := bucket.GetFiles(systemImagePaths)

	// Get all reviews for all maps in bulk
	reviewsByMap, err := reviews.GetReviewsByMapNames(mapNames)
	if err != nil {
		return nil, routeexception.NewRouteError(
			err,
			"Failed to get reviews for maps",
			"get-map-reviews-failed",
			err.Code,
		)
	}

	mapImages := make(map[string]*[]images.Image, len(mapNames))

	for _, mapName := range mapNames {
		systemImagePath := fmt.Sprintf("%s.jpg", mapName)
		hasSystemImage := systemImageExists[systemImagePath]
		mapReviews := reviewsByMap[mapName]

		if len(mapReviews) == 0 {
			if !hasSystemImage {
				mapImages[mapName] = nil // No images available
				continue
			} else {
				// Only system image available
				mapImages[mapName] = &[]images.Image{{
					ImagePath: "/api/images/" + mapName + ".jpg",
					Caption:   "System image for map",
				}}
				continue
			}
		}

		// Calculate total capacity for the image slice
		startingCapacity := 0
		if hasSystemImage {
			startingCapacity = 1
		}

		var imagesList []images.Image
		imagesList = make([]images.Image, 0, util.Reduce(
			mapReviews,
			startingCapacity,
			func(acc int, review reviews.MapReview) int {
				return acc + len(review.Images)
			},
		))

		// Add system image first if it exists
		if hasSystemImage {
			imagesList = append(imagesList, images.Image{
				ImagePath: "/api/images/" + mapName + ".jpg",
				Caption:   "System image for map",
			})
		}

		// Add images from reviews
		imagesFromReviews := util.FlatList(
			util.Map(
				mapReviews,
				func(review reviews.MapReview) []images.Image {
					return util.Map(review.Images, func(img reviews.MapReviewImage) images.Image {
						return img.Image
					})
				},
			))
		imagesList = append(imagesList, imagesFromReviews...)

		mapImages[mapName] = &imagesList
	}

	return mapImages, nil
}

// GetMapImagesByMapName is kept for backward compatibility but now uses the bulk function
func GetMapImagesSingle(mapName string) (*[]images.Image, routeexception.RouteError) {
	if mapName == "" {
		return nil, routeexception.NewRouteError(
			nil,
			"Map name cannot be empty",
			"map-name-empty",
			exception.CODE_BAD_REQUEST,
		)
	}

	// Use the bulk function for consistency
	result, err := GetMapImages([]string{mapName})
	if err != nil {
		return nil, err
	}

	images := result[mapName]
	if images == nil {
		return nil, routeexception.NewRouteError(
			nil,
			"No images found for map",
			"no-images-found",
			exception.CODE_RESOURCE_NOT_FOUND,
		)
	}

	return images, nil
}
