package mapperplugins

import (
	"reflect"

	"github.com/identityofsine/fofx-go-gin-api-template/internal/constants/exception"
)

type Int64Mapper struct{}

func (t Int64Mapper) GetDestinationObject() interface{} {
	// Return the type of the destination object as a string
	return int64(0)
}

func (t Int64Mapper) GetDestinationObjectString() string {
	// Return the type of the destination object as a string
	return reflect.TypeOf(t.GetDestinationObject()).Name()
}

func (t Int64Mapper) GetObject() interface{} {

	// Return the type of the object as a int64
	return int64(0)

}

func (t Int64Mapper) GetObjectString() string {
	// Return the type of the object as a string
	return reflect.TypeOf(t.GetObject()).Name()
}

func (t Int64Mapper) Map(obj interface{}) (interface{}, MapperError) {

	var nsObj int64

	// Check if the object is of type int64
	if o, ok := obj.(int64); !ok {
		// If the object is a string, reverse map it to sql.NullString
		if _, ok := obj.(int); ok {
			return t.ReverseMap(obj)
		}
		return nil, NewMapperError("nullstring::Map", "Invalid object type", "Expected sql.NullString", exception.CODE_BAD_REQUEST)
	} else {
		nsObj = o
	}

	return nsObj, nil

}

func (t Int64Mapper) ReverseMap(obj interface{}) (interface{}, MapperError) {

	// Check if the object is of type int64
	if o, ok := obj.(int64); ok {
		return o, nil
	}

	// If the object is a int, convert it to int64
	if num, ok := obj.(int); ok {
		return int64(num), nil
	}

	return nil, NewMapperError("nullstring::ReverseMap", "Invalid object type", "Expected int64 or int", exception.CODE_BAD_REQUEST)

}

func (t Int64Mapper) MapAll(objects []interface{}) ([]interface{}, MapperError) {
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

func (t Int64Mapper) ReverseMapAll(objects []interface{}) ([]interface{}, MapperError) {
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
