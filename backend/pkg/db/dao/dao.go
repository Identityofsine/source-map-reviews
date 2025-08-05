package dao

import (
	"fmt"
	"reflect"
	"strings"

	"github.com/fatih/camelcase"
	"github.com/identityofsine/fofx-go-gin-api-template/internal/constants/exception"
	"github.com/identityofsine/fofx-go-gin-api-template/pkg/db"
	"github.com/identityofsine/fofx-go-gin-api-template/util"
)

// Take in a Struct that ends in DB, and write some predefined methods

// InsertIntoDatabaseByStruct takes a struct that represents a database table and inserts it into the database.
// It returns the updated struct with any auto-generated fields (like IDs) populated.
func InsertIntoDatabaseByStruct[T interface{}](dbStruct T) (*T, db.DatabaseError) {
	tableName, err := getDbModelNameFromStruct(dbStruct)
	if err != nil {
		return nil, err
	}

	// Get fields suitable for insertion (handles primary keys and omitted fields automatically)
	insertFields, err := GetInsertFields(dbStruct)
	if err != nil {
		return nil, err
	}

	if len(insertFields) == 0 {
		return nil, db.NewDatabaseError("InsertIntoDatabaseByStruct", "No fields to insert", "no-fields-to-insert", exception.CODE_BAD_REQUEST)
	}

	// Build column names and values for insertion
	columns := make([]string, 0, len(insertFields))
	values := make([]interface{}, 0, len(insertFields))

	for _, field := range insertFields {
		columns = append(columns, field.DbColumnName)
		values = append(values, field.Value)
	}

	statement := fmt.Sprintf("INSERT INTO %s (%s) VALUES (%s)", tableName, strings.Join(columns, ", "), db.Placeholders(len(columns)))

	_, err = db.Query[interface{}](statement, values...)
	if err != nil {
		return nil, err
	}

	// Try to get the updated record with auto-generated fields
	// This assumes that primary key fields have been auto-generated or updated
	// We'll attempt to find the record using any available unique identifiers
	pkFields, pkErr := GetPrimaryKeyFields(dbStruct)
	if pkErr == nil && len(pkFields) > 0 {
		// Check if any primary key has a valid value to query with
		hasValidPK := false
		var foundPKField DaoFieldInfo
		for _, pkField := range pkFields {
			if pkField.IsPrimaryKey {
				foundPKField = pkField
				hasValidPK = true
				break
			}
		}
		if hasValidPK {

			obj, err := SelectFromDatabaseByStruct(dbStruct, "1 = 1 ORDER BY $1 ASC", foundPKField.Value)

			if err != nil {
				return nil, db.NewDatabaseError("InsertIntoDatabaseByStruct", "Failed to retrieve inserted record", "insert-retrieve-error", exception.CODE_INTERNAL_SERVER_ERROR)
			}
			if len(obj) > 0 {
				return &obj[len(obj)-1], nil // Return the last inserted record
			}

		}
	}

	// If we can't retrieve the updated record, return the original
	return &dbStruct, nil
}

// UpdateIntoDatabaseByStruct updates a database record using the struct's primary key.
// It only updates fields that have been changed from their zero values.
func UpdateIntoDatabaseByStruct(dbStruct interface{}) db.DatabaseError {
	// Validate that the primary key exists and is valid
	err := ValidatePrimaryKey(dbStruct)
	if err != nil {
		return err
	}

	// Check if the record exists in the database
	exists, err := CheckIfPrimaryKeyExistsInDatabase(dbStruct)
	if err != nil {
		return err
	}
	if !exists {
		return db.NewDatabaseError("UpdateIntoDatabaseByStruct", "Record with the given primary key does not exist", "record-not-found", exception.CODE_RESOURCE_NOT_FOUND)
	}

	tableName, err := getDbModelNameFromStruct(dbStruct)
	if err != nil {
		return err
	}

	// Get only changed fields (non-zero values, excluding primary keys and omitted fields)
	changedFields, err := GetChangedFields(dbStruct)
	if err != nil {
		return err
	}

	if len(changedFields) == 0 {
		return db.NewDatabaseError("UpdateIntoDatabaseByStruct", "No fields to update", "no-fields-to-update", exception.CODE_BAD_REQUEST)
	}

	// Build SET clause with changed fields
	columns := make([]string, 0, len(changedFields))
	values := make([]interface{}, 0, len(changedFields))

	// Create placeholders for SET clause
	setPlaceholders := db.Placeholders(len(changedFields))
	placeholderParts := strings.Split(setPlaceholders, ", ")

	for i, field := range changedFields {
		columns = append(columns, fmt.Sprintf("%s = %s", field.DbColumnName, placeholderParts[i]))
		values = append(values, field.Value)
	}

	// Build WHERE clause using primary key
	pkFields, err := GetPrimaryKeyFields(dbStruct)
	if err != nil {
		return err
	}

	var whereConditions []string
	var pkValues []interface{}

	// Create placeholders for WHERE clause continuing from SET clause
	whereStartIndex := len(values) + 1
	for i, pkField := range pkFields {
		whereConditions = append(whereConditions, fmt.Sprintf("%s = $%d", pkField.DbColumnName, whereStartIndex+i))
		pkValues = append(pkValues, pkField.Value)
	}

	whereClause := strings.Join(whereConditions, " AND ")
	statement := fmt.Sprintf("UPDATE %s SET %s WHERE %s", tableName, strings.Join(columns, ", "), whereClause)

	// Combine field values and primary key values
	allValues := append(values, pkValues...)

	_, err = db.Query[interface{}](statement, allValues...)

	return err
}

// UpdateIntoDatabaseByStructWithWhere provides the old behavior with custom WHERE clause
// Deprecated: Use UpdateIntoDatabaseByStruct for primary key-based updates
func UpdateIntoDatabaseByStructWithWhere(dbStruct interface{}, whereClause string, args ...interface{}) db.DatabaseError {
	tableName, err := getDbModelNameFromStruct(dbStruct)
	if err != nil {
		return err
	}

	columnsMap, err := getStructFields(dbStruct, false)
	if err != nil {
		return db.NewDatabaseError("UpdateIntoDatabaseByStructWithWhere", err.Message, err.Err, err.Code)
	}

	columns := make([]string, 0, len(columnsMap))
	fields := util.GetMapKeys(columnsMap)
	for _, field := range fields {
		columns = append(columns, fmt.Sprintf("%s = ?", columnsMap[field]))
	}

	statement := fmt.Sprintf("UPDATE %s SET %s", tableName, strings.Join(columns, ", "))
	if whereClause != "" {
		statement += " WHERE " + whereClause
	}

	fieldValues, err := getStructFieldsValues(dbStruct, fields)
	if err != nil {
		return db.NewDatabaseError("UpdateIntoDatabaseByStructWithWhere", err.Message, err.Err, err.Code)
	}

	_, err = db.Query[interface{}](statement, append(fieldValues, args...)...)

	return err
}

func SelectFromDatabaseByStruct[T interface{}](dbStruct T, whereClause string, args ...interface{}) ([]T, db.DatabaseError) {

	tableName, err := getDbModelNameFromStruct(dbStruct)
	if err != nil {
		return nil, err
	}

	fields, zerr := getStructFieldsArray(dbStruct)
	columnsMap, err := getStructFields(dbStruct, true)
	if err != nil {
		return nil, db.NewDatabaseError("InsertIntoDatabaseByStruct", err.Message, err.Err, err.Code)
	}
	if zerr != nil {
		return nil, db.NewDatabaseError("InsertIntoDatabaseByStruct", zerr.Message, zerr.Err, zerr.Code)
	}

	columns := make([]string, 0, len(columnsMap))
	for idx := range fields {
		field := fields[idx]
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

	tableNameSlice := camelcase.Split(structName[:len(structName)-2]) // Remove the "DB" suffix
	tableName := strings.Join((tableNameSlice), "_")

	// Remove the "DB" suffix to get the table name, and pluralize it
	tableName = fmt.Sprintf("%ss", strings.ToLower(tableName))
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

func getStructFieldsArray(dbStruct interface{}) ([]string, db.DatabaseError) {
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
