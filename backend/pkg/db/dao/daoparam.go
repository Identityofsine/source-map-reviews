package dao

import (
	"fmt"
	"reflect"
	"strings"

	"github.com/identityofsine/fofx-go-gin-api-template/internal/constants/exception"
	"github.com/identityofsine/fofx-go-gin-api-template/pkg/db"
)

// DaoFieldInfo contains information about a struct field's DAO annotations
type DaoFieldInfo struct {
	FieldName    string
	DbColumnName string
	IsPrimaryKey bool
	IsOmitted    bool
	Value        interface{}
}

// GetPrimaryKeyFields returns the primary key fields from a struct
func GetPrimaryKeyFields(dbStruct interface{}) ([]DaoFieldInfo, db.DatabaseError) {
	structType := reflect.TypeOf(dbStruct)
	structValue := reflect.ValueOf(dbStruct)

	if structType.Kind() != reflect.Struct {
		return nil, db.NewDatabaseError("GetPrimaryKeyFields", "Provided type is not a struct", "not-a-struct", exception.CODE_BAD_REQUEST)
	}

	var pkFields []DaoFieldInfo
	for i := 0; i < structType.NumField(); i++ {
		field := structType.Field(i)

		// Check if field has dao:"pk" tag
		if field.Tag.Get("dao") == "pk" {
			dbFieldName := field.Name
			if field.Tag.Get("db") != "" {
				dbFieldName = field.Tag.Get("db")
			}

			fieldInfo := DaoFieldInfo{
				FieldName:    field.Name,
				DbColumnName: dbFieldName,
				IsPrimaryKey: true,
				IsOmitted:    false,
				Value:        structValue.Field(i).Interface(),
			}
			pkFields = append(pkFields, fieldInfo)
		}
	}

	return pkFields, nil
}

// GetInsertFields returns fields suitable for INSERT operations (excludes primary keys with zero values and omitted fields)
func GetInsertFields(dbStruct interface{}) ([]DaoFieldInfo, db.DatabaseError) {
	structType := reflect.TypeOf(dbStruct)
	structValue := reflect.ValueOf(dbStruct)

	if structType.Kind() != reflect.Struct {
		return nil, db.NewDatabaseError("GetInsertFields", "Provided type is not a struct", "not-a-struct", exception.CODE_BAD_REQUEST)
	}

	var insertFields []DaoFieldInfo
	for i := 0; i < structType.NumField(); i++ {
		field := structType.Field(i)
		fieldValue := structValue.Field(i)
		daoTag := field.Tag.Get("dao")

		// Skip omitted fields
		if daoTag == "omit" {
			continue
		}

		// For primary keys, only include if they have non-zero values
		if daoTag == "pk" {
			zeroValue := reflect.Zero(field.Type)
			if reflect.DeepEqual(fieldValue.Interface(), zeroValue.Interface()) {
				continue // Skip zero-value primary keys (auto-increment)
			}
		}

		dbFieldName := field.Name
		if field.Tag.Get("db") != "" {
			dbFieldName = field.Tag.Get("db")
		}

		fieldInfo := DaoFieldInfo{
			FieldName:    field.Name,
			DbColumnName: dbFieldName,
			IsPrimaryKey: daoTag == "pk",
			IsOmitted:    false,
			Value:        fieldValue.Interface(),
		}
		insertFields = append(insertFields, fieldInfo)
	}

	return insertFields, nil
}

// GetNonPrimaryKeyFields returns all non-primary key fields that are not omitted
func GetNonPrimaryKeyFields(dbStruct interface{}, includeOmitted bool) ([]DaoFieldInfo, db.DatabaseError) {
	structType := reflect.TypeOf(dbStruct)
	structValue := reflect.ValueOf(dbStruct)

	if structType.Kind() != reflect.Struct {
		return nil, db.NewDatabaseError("GetNonPrimaryKeyFields", "Provided type is not a struct", "not-a-struct", exception.CODE_BAD_REQUEST)
	}

	var fields []DaoFieldInfo
	for i := 0; i < structType.NumField(); i++ {
		field := structType.Field(i)
		daoTag := field.Tag.Get("dao")

		// Skip primary keys
		if daoTag == "pk" {
			continue
		}

		// Handle omitted fields
		isOmitted := daoTag == "omit"
		if isOmitted && !includeOmitted {
			continue
		}

		dbFieldName := field.Name
		if field.Tag.Get("db") != "" {
			dbFieldName = field.Tag.Get("db")
		}

		fieldInfo := DaoFieldInfo{
			FieldName:    field.Name,
			DbColumnName: dbFieldName,
			IsPrimaryKey: false,
			IsOmitted:    isOmitted,
			Value:        structValue.Field(i).Interface(),
		}
		fields = append(fields, fieldInfo)
	}

	return fields, nil
}

// GetChangedFields returns fields that have been modified from their zero values
func GetChangedFields(dbStruct interface{}) ([]DaoFieldInfo, db.DatabaseError) {
	structType := reflect.TypeOf(dbStruct)
	structValue := reflect.ValueOf(dbStruct)

	if structType.Kind() != reflect.Struct {
		return nil, db.NewDatabaseError("GetChangedFields", "Provided type is not a struct", "not-a-struct", exception.CODE_BAD_REQUEST)
	}

	var changedFields []DaoFieldInfo
	for i := 0; i < structType.NumField(); i++ {
		field := structType.Field(i)
		fieldValue := structValue.Field(i)
		daoTag := field.Tag.Get("dao")

		// Skip primary keys and omitted fields for change detection
		if daoTag == "pk" || daoTag == "omit" {
			continue
		}

		// Check if field value is different from zero value
		zeroValue := reflect.Zero(field.Type)
		if !reflect.DeepEqual(fieldValue.Interface(), zeroValue.Interface()) {
			dbFieldName := field.Name
			if field.Tag.Get("db") != "" {
				dbFieldName = field.Tag.Get("db")
			}

			fieldInfo := DaoFieldInfo{
				FieldName:    field.Name,
				DbColumnName: dbFieldName,
				IsPrimaryKey: false,
				IsOmitted:    false,
				Value:        fieldValue.Interface(),
			}
			changedFields = append(changedFields, fieldInfo)
		}
	}

	return changedFields, nil
}

// ValidatePrimaryKey checks if the primary key fields have valid values
func ValidatePrimaryKey(dbStruct interface{}) db.DatabaseError {
	pkFields, err := GetPrimaryKeyFields(dbStruct)
	if err != nil {
		return err
	}

	if len(pkFields) == 0 {
		return db.NewDatabaseError("ValidatePrimaryKey", "No primary key fields found", "no-primary-key", exception.CODE_BAD_REQUEST)
	}

	for _, pkField := range pkFields {
		// Check if primary key value is zero value
		fieldType := reflect.TypeOf(pkField.Value)
		zeroValue := reflect.Zero(fieldType)

		if reflect.DeepEqual(pkField.Value, zeroValue.Interface()) {
			return db.NewDatabaseError("ValidatePrimaryKey", fmt.Sprintf("Primary key field %s has zero value", pkField.FieldName), "invalid-primary-key", exception.CODE_BAD_REQUEST)
		}
	}

	return nil
}

// BuildPrimaryKeyWhereClause builds a WHERE clause for primary key fields
func BuildPrimaryKeyWhereClause(dbStruct interface{}) (string, []interface{}, db.DatabaseError) {
	pkFields, err := GetPrimaryKeyFields(dbStruct)
	if err != nil {
		return "", nil, err
	}

	if len(pkFields) == 0 {
		return "", nil, db.NewDatabaseError("BuildPrimaryKeyWhereClause", "No primary key fields found", "no-primary-key", exception.CODE_BAD_REQUEST)
	}

	var conditions []string
	var values []interface{}

	for i, pkField := range pkFields {
		conditions = append(conditions, fmt.Sprintf("%s = $%d", pkField.DbColumnName, i+1))
		values = append(values, pkField.Value)
	}

	whereClause := strings.Join(conditions, " AND ")
	return whereClause, values, nil
}

// CheckIfPrimaryKeyExistsInDatabase checks if a record with the given primary key exists
func CheckIfPrimaryKeyExistsInDatabase(dbStruct interface{}) (bool, db.DatabaseError) {
	tableName, err := getDbModelNameFromStruct(dbStruct)
	if err != nil {
		return false, err
	}

	whereClause, args, err := BuildPrimaryKeyWhereClause(dbStruct)
	if err != nil {
		return false, err
	}

	// Use a simple COUNT query instead of scanning all columns
	rows, err := db.Query[int]("SELECT COUNT(*) FROM "+tableName+" WHERE "+whereClause, args...)
	if err != nil {
		return false, err
	}

	if len(*rows) == 0 {
		return false, nil
	}

	count := (*rows)[0]
	return count > 0, nil
}
