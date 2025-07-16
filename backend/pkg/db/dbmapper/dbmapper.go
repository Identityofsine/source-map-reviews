package dbmapper

import "reflect"

// MapToDb maps fields from an input struct to an output struct using reflection.
func MapDbFields[in any, out any](input in) *out {

	// Use reflection to map fields from input to output
	inputValue := reflect.ValueOf(input)
	if inputValue.Kind() != reflect.Struct {
		return nil
	}

	// Ensure input is a struct
	output, outputValue := createReflectValue[out]()

	mapFields := make(map[string]reflect.Value)

	for i := 0; i < inputValue.NumField(); i++ {
		field := inputValue.Type().Field(i)
		dbTag := field.Tag.Get("db")
		if dbTag == "" {
			continue // Skip fields without a db tag
		}
		mapFields[dbTag] = inputValue.Field(i)
	}

	//change the output type fields that are matched by db tag
	for i := 0; i < outputValue.NumField(); i++ {
		oType := reflect.TypeOf(output).Elem()
		field := oType.Field(i)
		dbTag := field.Tag.Get("db")
		if dbTag == "" {
			continue // Skip fields without a db Tag
		}
		if value, ok := mapFields[dbTag]; ok {
			if outputValue.FieldByName(field.Name).CanSet() {
				outputValue.FieldByName(field.Name).Set(value)
			} else {
				return nil // Return nil if the field cannot be set
			}
		}
	}

	return output
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
