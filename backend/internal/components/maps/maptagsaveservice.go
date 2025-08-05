package maps

import (
	"fmt"

	"github.com/identityofsine/fofx-go-gin-api-template/internal/components/user"
	"github.com/identityofsine/fofx-go-gin-api-template/internal/constants/exception"
	"github.com/identityofsine/fofx-go-gin-api-template/internal/repository"
	"github.com/identityofsine/fofx-go-gin-api-template/internal/types/routeexception"
	"github.com/identityofsine/fofx-go-gin-api-template/pkg/cookies"
	"github.com/identityofsine/fofx-go-gin-api-template/pkg/db/dbmapper"
	"github.com/identityofsine/fofx-go-gin-api-template/pkg/storedlogs"
)

// SaveTagLk handles both creating new tag lookups and updating existing ones
func SaveTagLk(tagLk TagLk, cookies cookies.Cookies) (*TagLk, routeexception.RouteError) {
	currentUser, err := user.GetUserByCookies(&cookies)
	if err != nil || currentUser == nil {
		return nil, exception.BadRequest
	}

	// Validate required fields
	if tagLk.TagLk == "" {
		return nil, routeexception.NewRouteError(nil, "Tag lookup key is required", "tag-lk-required", exception.CODE_BAD_REQUEST)
	}

	if tagLk.TagDescription == "" {
		return nil, routeexception.NewRouteError(nil, "Tag description is required", "tag-description-required", exception.CODE_BAD_REQUEST)
	}

	// Map to database model and save
	mapped := dbmapper.MapDbFields[TagLk, repository.LkTagDB](tagLk)
	savedTagLk, dbErr := repository.SaveLkTagDb(*mapped)
	if dbErr != nil {
		return nil, routeexception.NewRouteError(dbErr, "Failed to save tag lookup", "save-tag-lk-failed", dbErr.Code)
	}

	// Map back to domain model
	result := dbmapper.MapDbFields[repository.LkTagDB, TagLk](*savedTagLk)
	if result == nil {
		return nil, routeexception.NewRouteError(nil, "Failed to map saved tag lookup", "map-saved-tag-lk-failed", exception.CODE_INTERNAL_SERVER_ERROR)
	}

	return result, nil
}

// DeleteTagLk handles deleting a tag lookup
func DeleteTagLk(tagLk string, cookies cookies.Cookies) routeexception.RouteError {
	currentUser, err := user.GetUserByCookies(&cookies)
	if err != nil || currentUser == nil {
		return exception.BadRequest
	}

	if tagLk == "" {
		return routeexception.NewRouteError(nil, "Tag lookup key is required", "tag-lk-required", exception.CODE_BAD_REQUEST)
	}

	// Check if this tag is being used by any maps
	tags, dbErr := repository.GetLkTagsByLkTags([]string{tagLk})
	if dbErr != nil {
		return routeexception.NewRouteError(dbErr, "Failed to check tag usage", "check-tag-usage-failed", dbErr.Code)
	}

	if tags != nil && len(*tags) > 0 {
		// Check if any maps are using this tag
		allMapTags, dbErr := repository.GetMapTags()
		if dbErr == nil && allMapTags != nil {
			for _, mapTag := range *allMapTags {
				if mapTag.LkTag == tagLk {
					return routeexception.NewRouteError(nil, "Cannot delete tag that is in use by maps", "tag-in-use", exception.CODE_BAD_REQUEST)
				}
			}
		}
	}

	// Delete the tag lookup
	dbErr = repository.DeleteLkTagDb(tagLk)
	if dbErr != nil {
		return routeexception.NewRouteError(dbErr, "Failed to delete tag lookup", "delete-tag-lk-failed", dbErr.Code)
	}

	return nil
}

// AddTagToMap adds a tag to a map
func AddTagToMap(mapName string, tagLk string, cookies cookies.Cookies) (*MapTag, routeexception.RouteError) {
	currentUser, err := user.GetUserByCookies(&cookies)
	if err != nil || currentUser == nil {
		return nil, exception.BadRequest
	}

	// Validate required fields
	if mapName == "" {
		return nil, routeexception.NewRouteError(nil, "Map name is required", "map-name-required", exception.CODE_BAD_REQUEST)
	}

	if tagLk == "" {
		return nil, routeexception.NewRouteError(nil, "Tag lookup key is required", "tag-lk-required", exception.CODE_BAD_REQUEST)
	}

	// Verify the tag lookup exists
	existingTagLk, dbErr := repository.GetLkTagByLkTag(tagLk)
	if dbErr != nil {
		return nil, routeexception.NewRouteError(dbErr, "Failed to verify tag lookup", "verify-tag-lk-failed", dbErr.Code)
	}

	if existingTagLk == nil {
		return nil, routeexception.NewRouteError(nil, "Tag lookup does not exist", "tag-lk-not-found", exception.CODE_RESOURCE_NOT_FOUND)
	}

	// Create the map-tag relationship
	mapTagDB := repository.MapTagDB{
		LkTag:   tagLk,
		MapName: mapName,
	}

	savedMapTag, dbErr := repository.SaveMapTagDb(mapTagDB)
	if dbErr != nil {
		return nil, routeexception.NewRouteError(dbErr, "Failed to add tag to map", "add-tag-to-map-failed", dbErr.Code)
	}

	// Map to domain model
	result := dbmapper.MapDbFields[repository.MapTagDB, MapTag](*savedMapTag)
	if result == nil {
		return nil, routeexception.NewRouteError(nil, "Failed to map saved map tag", "map-saved-map-tag-failed", exception.CODE_INTERNAL_SERVER_ERROR)
	}

	// Populate with tag details
	result.TagDescription = existingTagLk.Description.String
	result.TagDescriptionShort = existingTagLk.ShortDescription.String

	return result, nil
}

// RemoveTagFromMap removes a tag from a map
func RemoveTagFromMap(mapName string, tagLk string, cookies cookies.Cookies) routeexception.RouteError {
	currentUser, err := user.GetUserByCookies(&cookies)
	if err != nil || currentUser == nil {
		return exception.BadRequest
	}

	// Validate required fields
	if mapName == "" {
		return routeexception.NewRouteError(nil, "Map name is required", "map-name-required", exception.CODE_BAD_REQUEST)
	}

	if tagLk == "" {
		return routeexception.NewRouteError(nil, "Tag lookup key is required", "tag-lk-required", exception.CODE_BAD_REQUEST)
	}

	// Delete the map-tag relationship
	dbErr := repository.DeleteMapTagDb(mapName, tagLk)
	if dbErr != nil {
		return routeexception.NewRouteError(dbErr, "Failed to remove tag from map", "remove-tag-from-map-failed", dbErr.Code)
	}

	return nil
}

// UpdateMapTags replaces all tags for a map with the provided list
func UpdateMapTags(mapName string, tags []string, cookies cookies.Cookies) ([]MapTag, routeexception.RouteError) {
	currentUser, err := user.GetUserByCookies(&cookies)
	if err != nil || currentUser == nil {
		return nil, exception.BadRequest
	}

	// Validate required fields
	if mapName == "" {
		return nil, routeexception.NewRouteError(nil, "Map name is required", "map-name-required", exception.CODE_BAD_REQUEST)
	}

	// Verify all tags exist
	if len(tags) > 0 {
		existingTags, dbErr := repository.GetLkTagsByLkTags(tags)
		if dbErr != nil {
			return nil, routeexception.NewRouteError(dbErr, "Failed to verify tags", "verify-tags-failed", dbErr.Code)
		}

		if existingTags == nil || len(*existingTags) != len(tags) {
			return nil, routeexception.NewRouteError(nil, "One or more tags do not exist", "tags-not-found", exception.CODE_BAD_REQUEST)
		}
	}

	// Delete all existing tags for this map
	dbErr := repository.DeleteMapTagsByMapName(mapName)
	if dbErr != nil {
		storedlogs.LogWarn(fmt.Sprintf("Warning: Failed to delete existing tags for map %s: %v", mapName, dbErr))
	}

	// Add new tags
	var result []MapTag
	result = make([]MapTag, 0, len(tags))

	for _, tagLk := range tags {
		addedTag, err := AddTagToMap(mapName, tagLk, cookies)
		if err != nil {
			storedlogs.LogError(fmt.Sprintf("Failed to add tag %s to map %s: %v", tagLk, mapName, err), err)
			continue
		}
		result = append(result, *addedTag)
	}

	return result, nil
}
