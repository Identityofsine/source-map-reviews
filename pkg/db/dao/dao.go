package dao

import (
	"fmt"
	"reflect"
	"strings"

	"github.com/identityofsine/fofx-go-gin-api-template/internal/constants/exception"
	"github.com/identityofsine/fofx-go-gin-api-template/pkg/db"
	"github.com/identityofsine/fofx-go-gin-api-template/util"
)

// Take in a Struct that ends in DB, and write some predefined methods

// InsertIntoDatabaseByStruct takes a struct that represents a database table and inserts it into the database.
func InsertIntoDatabaseByStruct(dbStruct interface{}) db.DatabaseError {

	tableName, err := getDbModelNameFromStruct(dbStruct)
	if err != nil {
		return err
	}

	columnsMap, err := getStructFields(dbStruct, false)
	if err != nil {
		return db.NewDatabaseError("InsertIntoDatabaseByStruct", err.Message, err.Err, err.Code)
	}

	columns := make([]string, 0, len(columnsMap))
	fields := util.GetMapKeys(columnsMap)
	for _, field := range fields {
		columns = append(columns, columnsMap[field])
	}

	statement := fmt.Sprintf("INSERT INTO %s (%s) VALUES (%s)", tableName, strings.Join(columns, ", "), db.Placeholders(len(columns)))

	fieldValues, err := getStructFieldsValues(dbStruct, fields)

	_, err = db.Query[interface{}](statement, fieldValues...)

	return err
}

func SelectFromDatabaseByStruct[T interface{}](dbStruct T, whereClause string, args ...interface{}) ([]T, db.DatabaseError) {

	tableName, err := getDbModelNameFromStruct(dbStruct)
	if err != nil {
		return nil, err
	}

	fields, err := getStructFieldsArray(dbStruct, true)
	columnsMap, err := getStructFields(dbStruct, true)
	if err != nil {
		return nil, db.NewDatabaseError("InsertIntoDatabaseByStruct", err.Message, err.Err, err.Code)
	}

	columns := make([]string, 0, len(columnsMap))
	for idx := range fields {
		field := fields[idx]
		fmt.Println("Field:", field)
		columns = append(columns, columnsMap[field])
	}

	statement := fmt.Sprintf("SELECT %s FROM %s", strings.Join(columns, ", "), tableName)
	if whereClause != "" {
		statement += " WHERE " + whereClause
	}

	rows, err := db.Query[T](statement, args...)

	if err != nil {
		return nil, err
	}

	return *rows, nil
}

func getDbModelNameFromStruct(dbStruct interface{}) (string, db.DatabaseError) {
	structName := getStructName(dbStruct)
	if structName == "" {
		return "", db.NewDatabaseError("InsertIntoDatabaseByStruct", "Invalid struct type", "invalid-struct", exception.CODE_BAD_REQUEST)
	}
	if !doesStringEndWithDB(structName) {
		return "", db.NewDatabaseError("InsertIntoDatabaseByStruct", "Struct name must end with 'DB'", "struct-name-must-end-with-DB", exception.CODE_BAD_REQUEST)
	}
	// Remove the "DB" suffix to get the table name, and pluralize it
	tableName := fmt.Sprintf("%ss", strings.ToLower(structName)[:len(structName)-2])
	return tableName, nil
}

func getStructFieldsValues(dbStruct interface{}, fields []string) ([]interface{}, db.DatabaseError) {
	if dbStruct == nil {
		return nil, db.NewDatabaseError("getStructFieldsValues", "Invalid struct type", "invalid-struct", exception.CODE_BAD_REQUEST)
	}

	structType := reflect.TypeOf(dbStruct)
	if structType.Kind() != reflect.Struct {
		return nil, db.NewDatabaseError("getStructFieldsValues", "Provided type is not a struct", "not-a-struct", exception.CODE_BAD_REQUEST)
	}

	var values []interface{}
	for _, fieldName := range fields {
		field, ok := structType.FieldByName(fieldName)
		if !ok {
			return nil, db.NewDatabaseError("getStructFieldsValues", fmt.Sprintf("Field %s not found in struct", fieldName), "field-not-found", exception.CODE_BAD_REQUEST)
		}
		value := reflect.ValueOf(dbStruct).FieldByName(field.Name).Interface()
		values = append(values, value)
	}

	return values, nil
}

func getStructFieldsArray(dbStruct interface{}, canOmit bool) ([]string, db.DatabaseError) {
	if dbStruct == nil {
		return nil, db.NewDatabaseError("getStructFieldsArray", "Invalid struct type", "invalid-struct", exception.CODE_BAD_REQUEST)
	}

	structType := reflect.TypeOf(dbStruct)
	if structType.Kind() != reflect.Struct {
		return nil, db.NewDatabaseError("getStructFieldsArray", "Provided type is not a struct", "not-a-struct", exception.CODE_BAD_REQUEST)
	}

	var fields []string
	for i := 0; i < structType.NumField(); i++ {
		field := structType.Field(i)
		fields = append(fields, field.Name)
	}

	return fields, nil
}

// getStructFields retrieves the fields of a struct and returns them as a map where the key is the field name
func getStructFields(dbStruct interface{}, canOmit bool) (map[string]string, db.DatabaseError) {
	if dbStruct == nil {
		return nil, db.NewDatabaseError("getStructFields", "Invalid struct type", "invalid-struct", exception.CODE_BAD_REQUEST)
	}

	structType := reflect.TypeOf(dbStruct)
	if structType.Kind() != reflect.Struct {
		return nil, db.NewDatabaseError("getStructFields", "Provided type is not a struct", "not-a-struct", exception.CODE_BAD_REQUEST)
	}

	var fields = make(map[string]string)
	for i := 0; i < structType.NumField(); i++ {
		field := structType.Field(i)
		dbFieldName := field.Name
		if field.Tag.Get("dao") == "omit" && !canOmit {
			continue
		}
		if field.Tag.Get("db") != "" {
			// use the db tag if it exists
			dbFieldName = field.Tag.Get("db")
		} else {
			// do some extra logic here to mutate if possible
			// for example, convert CamelCase to snake_case
		}
		fields[field.Name] = dbFieldName
	}

	return fields, nil
}

func getStructName(dbStruct interface{}) string {
	dbStructType := reflect.TypeOf(dbStruct).Name()
	return dbStructType
}

func doesStringEndWithDB(s string) bool {
	if len(s) < 2 {
		return false
	}
	return s[len(s)-2:] == "DB"
}
