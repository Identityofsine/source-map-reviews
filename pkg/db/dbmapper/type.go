package dbmapper

import (
	"fmt"
	"reflect"

	. "github.com/identityofsine/fofx-go-gin-api-template/pkg/db/dbmapper/mapperplugins"
)

// attemptTypeCastThroughPlugins tries to find a plugin that can convert the input type to the output type.
func attemptTypeCastThroughPlugins(inputValue reflect.Value, outputType reflect.Type) (*reflect.Value, bool) {

	inputType := inputValue.Type()

	// Check if the input type is assignable to the output type
	// If they are the same type, return the output type
	if isTypeMatch(inputType, outputType) {
		return &inputValue, true
	}

	plugin := GetMapperPlugin(inputType.Name(), outputType.Name())
	if plugin == nil {
		return nil, false
	}

	// Get the actual value from the reflect.Value
	if inputValue.Kind() == reflect.Ptr {
		if inputValue.IsNil() {
			return nil, false
		}
		inputValue = inputValue.Elem()
	}

	out, err := plugin.Map(inputValue.Interface())
	if err != nil {
		fmt.Println("Error in plugin mapping:", err)
		return nil, false
	} else if out == nil {
		return nil, false
	}

	newOut := reflect.ValueOf(out)

	return &newOut, true
}

func isTypeMatch(inputType reflect.Type, outputType reflect.Type) bool {
	// Check if the input type is assignable to the output type
	if inputType.Kind() == reflect.Ptr && outputType.Kind() == reflect.Ptr {
		return inputType.Elem().AssignableTo(outputType.Elem())
	}
	return inputType.AssignableTo(outputType)
}
