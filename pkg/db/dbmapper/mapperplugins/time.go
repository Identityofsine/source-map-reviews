package mapperplugins

import (
	"fmt"
	"reflect"
	"time"

	"github.com/identityofsine/fofx-go-gin-api-template/internal/constants/exception"
)

type TimeMapper struct{}

func (t TimeMapper) GetDestinationObject() interface{} {
	// Return the type of the destination object as a string
	return ""
}

func (t TimeMapper) GetDestinationObjectString() string {
	// Return the type of the destination object as a string
	return reflect.TypeOf(t.GetDestinationObject()).Name()
}

func (t TimeMapper) GetObject() interface{} {

	// Return the type of the object as a time.Time
	return time.Time{}

}

func (t TimeMapper) GetObjectString() string {
	// Return the type of the object as a string
	return reflect.TypeOf(t.GetObject()).Name()
}

func (t TimeMapper) Map(obj interface{}) (interface{}, MapperError) {

	var tObj time.Time

	fmt.Println("Mapping time.Time to string", obj)

	// Check if the object is of type time.Time
	if o, ok := obj.(time.Time); !ok {
		if _, ok := obj.(string); ok {
			// If the object is a string, reverse map it to time.Time
			return t.ReverseMap(obj)
		}
		return nil, NewMapperError("time::Map", "Invalid object type", "Expected time.Time", exception.CODE_BAD_REQUEST)
	} else {
		tObj = o
	}

	// Map the time.Time object to a string representation
	string := tObj.Format(time.RFC3339)

	return string, nil
}

func (t TimeMapper) ReverseMap(obj interface{}) (interface{}, MapperError) {

	var sObj string

	// Check if the object is of type string
	if o, ok := obj.(string); !ok {
		if _, ok := obj.(time.Time); ok {
			// If the object is a time.Time, reverse map it to string
			return t.Map(obj)
		}
		return nil, NewMapperError("time::ReverseMap", "Invalid object type", "Expected string", exception.CODE_BAD_REQUEST)
	} else {
		sObj = o
	}

	// Reverse map the string representation back to a time.Time object
	parsedTime, err := time.Parse(time.RFC3339, sObj)
	if err != nil {
		return time.Time{}, NewMapperError("time::ReverseMap", "Failed to parse time", err.Error(), exception.CODE_BAD_REQUEST)
	}

	return parsedTime, nil
}

func (t TimeMapper) MapAll(objects []interface{}) ([]interface{}, MapperError) {
	// Map all time.Time objects to their string representations
	var result []interface{}
	for _, obj := range objects {
		mapped, err := t.Map(obj)
		if err != nil {
			return nil, err
		}
		result = append(result, mapped)
	}

	return result, nil
}

func (t TimeMapper) ReverseMapAll(objects []interface{}) ([]interface{}, MapperError) {
	// Reverse map all string representations back to time.Time objects
	var result []interface{}
	for _, obj := range objects {
		reversed, err := t.ReverseMap(obj)
		if err != nil {
			return nil, err
		}
		result = append(result, reversed)
	}

	return result, nil
}
