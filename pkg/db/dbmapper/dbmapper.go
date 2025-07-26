package dbmapper

import (
	"reflect"
)

// MapAllDbFields maps fields from a slice of input structs to a slice of output structs using reflection.
func MapAllDbFields[in any, out any](input []in) *[]out {
	outs := make([]out, 0, len(input))
	for _, inp := range input {
		if outValue := MapDbFields[in, out](inp); outValue != nil {
			outs = append(outs, *outValue)
		}
	}

	return &outs
}

// MapToDb maps fields from an input struct to an output struct using reflection.
func MapDbFields[in any, out any](input in) *out {

	// Use reflection to map fields from input to output
	mapFields := getReflectionValues[in](input, "db")
	if mapFields == nil {
		return nil // Return nil if input is not a struct or has no db tags
	}

	// Ensure input is a struct
	output, outputValue := createReflectValue[out]()

	//change the output type fields that are matched by db tag

	if output == nil || outputValue == nil {
		return nil // Return nil if output is not a pointer or is Nil
	}
	// Set the fields in the output struct using the reflection map
	if outputValue.Kind() != reflect.Struct {
		return nil // Return nil if output is not a Struct
	}

	setReflectValues(output, mapFields)

	return output
}

// MapDbFullFields maps all fields from a stream of input structs, the following objects will be subject to nested mapping
func MapDbFullFields[in any, out any](input in, object ...any) *out {
	// Map the basic fields
	output := MapDbFields[in, out](input)
	if output == nil {
		return nil
	}

	// If no nested objects, return the mapped struct
	if len(object) == 0 {
		return output
	}

	// Get the output struct as a reflect.Value
	outputValue := reflect.ValueOf(output).Elem()
	if outputValue.Kind() != reflect.Struct {
		return nil
	}

	// For each field in the output struct, check for dbobj tag
	for i := 0; i < outputValue.NumField(); i++ {
		field := outputValue.Type().Field(i)
		dbobjTag := field.Tag.Get("dbobj")
		if dbobjTag == "" {
			continue
		}

		// Find the matching object by type
		for _, obj := range object {
			objValue := reflect.ValueOf(obj)

			// Accept both pointer and non-pointer values
			if objValue.Kind() == reflect.Ptr && !objValue.IsNil() {
				objValue = objValue.Elem()
			}

			// Set if assignable to the field
			if objValue.Type().AssignableTo(field.Type) {
				outputValue.Field(i).Set(objValue)
				break
			}
		}
	}

	return output
}

func getReflectionValues[in any](obj in, tag string) map[string]reflect.Value {

	// Use reflection to map fields from input to output
	inputValue := reflect.ValueOf(obj)
	if inputValue.Kind() != reflect.Struct {
		return nil
	}

	mapFields := make(map[string]reflect.Value)

	for i := 0; i < inputValue.NumField(); i++ {
		field := inputValue.Type().Field(i)
		dbTag := field.Tag.Get(tag)
		if dbTag == "" {
			continue // Skip fields without a db tag
		}
		mapFields[dbTag] = inputValue.Field(i)
	}

	return mapFields

}

func setReflectValues[in any](output in, reflectionMap map[string]reflect.Value) *in {
	// Use reflection to set fields in the output struct
	outputValue := reflect.ValueOf(output).Elem()
	if outputValue.Kind() != reflect.Struct {
		return nil // Return nil if output is not a struct
	}

	for i := 0; i < outputValue.NumField(); i++ {
		field := outputValue.Type().Field(i)
		dbTag := field.Tag.Get("db")
		if dbTag == "" {
			continue // Skip fields without a db tag
		}
		if value, ok := reflectionMap[dbTag]; ok {
			if outputValue.FieldByName(field.Name).CanSet() {
				outputValue.FieldByName(field.Name).Set(value)
			} else {
				return nil // Return nil if the field cannot be set
			}
		}
	}

	return &output
}

func createReflectValue[IN any]() (*IN, *reflect.Value) {
	// Ensure input is a struct
	output := new(IN) // Create a new instance of the output type
	if reflect.TypeOf(output).Kind() != reflect.Ptr || reflect.ValueOf(output).IsNil() {
		return nil, nil // Return nil if output is not a pointer or is nil
	}
	outputValue := reflect.ValueOf(output).Elem()

	return output, &outputValue

}
