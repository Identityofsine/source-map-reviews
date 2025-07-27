package mapperplugins

import (
	"database/sql"
	"reflect"

	"github.com/identityofsine/fofx-go-gin-api-template/internal/constants/exception"
)

type NullStringMapper struct{}

func (t NullStringMapper) GetDestinationObject() interface{} {
	// Return the type of the destination object as a string
	return ""
}

func (t NullStringMapper) GetDestinationObjectString() string {
	// Return the type of the destination object as a string
	return reflect.TypeOf(t.GetDestinationObject()).Name()
}

func (t NullStringMapper) GetObject() interface{} {

	// Return the type of the object as a time.Time
	return sql.NullString{}

}

func (t NullStringMapper) GetObjectString() string {
	// Return the type of the object as a string
	return reflect.TypeOf(t.GetObject()).Name()
}

func (t NullStringMapper) Map(obj interface{}) (interface{}, MapperError) {

	var nsObj sql.NullString

	// Check if the object is of type sql.NullString
	if o, ok := obj.(sql.NullString); !ok {
		// If the object is a string, reverse map it to sql.NullString
		if _, ok := obj.(string); ok {
			return t.ReverseMap(obj)
		}
		return nil, NewMapperError("nullstring::Map", "Invalid object type", "Expected sql.NullString", exception.CODE_BAD_REQUEST)
	} else {
		nsObj = o
	}

	var string string
	// Map the sql.NullString object to a string representation
	if !nsObj.Valid {
		string = ""
	} else {
		string = nsObj.String
	}

	return string, nil
}

func (t NullStringMapper) ReverseMap(obj interface{}) (interface{}, MapperError) {

	var sObj string

	// Check if the object is of type string
	if o, ok := obj.(string); !ok {
		// If the object is a sql.NullString, reverse map it to a string
		if _, ok := obj.(sql.NullString); ok {
			return t.Map(obj)
		}
		return nil, NewMapperError("nullstring::ReverseMap", "Invalid object type", "Expected string", exception.CODE_BAD_REQUEST)
	} else {
		sObj = o
	}

	// Reverse map the string representation back to a time.Time object
	parsedTime := sql.NullString{
		String: sObj,
		Valid:  sObj != "",
	}

	return parsedTime, nil
}

func (t NullStringMapper) MapAll(objects []interface{}) ([]interface{}, MapperError) {
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

func (t NullStringMapper) ReverseMapAll(objects []interface{}) ([]interface{}, MapperError) {
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
