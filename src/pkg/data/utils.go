package data

import (
	"encoding/json"
	"reflect"
	"time"

	"gorm.io/gorm"
)

const HTML_INPUT_TEXT_TYPE = "text"
const HTML_INPUT_EMAIL_TYPE = "email"
const HTML_INPUT_DATE_TYPE = "date"
const HTML_INPUT_DATETIME_TYPE = "datetime-local"
const HTML_INPUT_BOOL_TYPE = "checkbox"
const HTML_INPUT_NUMBER_TYPE = "number"

const GORM_PK_DEFAULT_TAG_NAME string = "gorm"
const GORM_PK_DEFAULT_TAG_VALUE string = "primaryKey"
const GORM_PK_DEFAULT_NAME string = "ID"

func FieldHasEmbeddedStructs(f reflect.StructField) bool {
	// return f.Anonymous || f.Type == reflect.TypeOf(gorm.DeletedAt{})
	return f.Anonymous
}

func IsPkField(f reflect.StructField) bool {
	return f.Tag.Get(GORM_PK_DEFAULT_TAG_NAME) == GORM_PK_DEFAULT_TAG_VALUE || f.Name == GORM_PK_DEFAULT_NAME
}

func GetHtmlInputType(f reflect.StructField) string {
	switch f.Type.Kind() {
	case reflect.String:
		return HTML_INPUT_TEXT_TYPE
	case reflect.Uint, reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return HTML_INPUT_NUMBER_TYPE
	case reflect.Bool:
		return HTML_INPUT_BOOL_TYPE
	case reflect.Float32, reflect.Float64:
		return HTML_INPUT_NUMBER_TYPE
	case reflect.TypeOf(time.Time{}).Kind(), reflect.TypeOf(gorm.DeletedAt{}).Kind():
		return HTML_INPUT_DATETIME_TYPE
	default:
		return HTML_INPUT_TEXT_TYPE
	}
}

func FindPkField(fields []reflect.StructField) reflect.StructField {
	for _, f := range fields {
		if IsPkField(f) {
			return f
		}
	}

	panic("Gorm model doesn't have PK")
}

func GetObjectFieldsValues(objectValue reflect.Value) []reflect.Value {
	if objectValue.Kind() == reflect.Ptr {
		objectValue = objectValue.Elem()
	}

	var objectFieldsValues []reflect.Value

	objectType := objectValue.Type()
	for i := 0; i < objectValue.NumField(); i++ {
		fieldType := objectType.Field(i)
		fieldValue := objectValue.Field(i)

		if FieldHasEmbeddedStructs(fieldType) {
			embeddedFieldsValues := GetObjectFieldsValues(fieldValue)
			objectFieldsValues = append(objectFieldsValues, embeddedFieldsValues...)
		} else {
			objectFieldsValues = append(objectFieldsValues, fieldValue)
		}
	}

	return objectFieldsValues
}

func GetObjectFields(objectType reflect.Type) []reflect.StructField {
	var objectFields []reflect.StructField

	for i := 0; i < objectType.NumField(); i++ {
		fieldType := objectType.Field(i)

		if FieldHasEmbeddedStructs(fieldType) {
			embeddedFields := GetObjectFields(fieldType.Type)
			objectFields = append(objectFields, embeddedFields...)
		} else {
			objectFields = append(objectFields, fieldType)
		}
	}

	return objectFields
}

// func GetTypesNames(objectsTypes *[]reflect.Type) []string {
// 	var names []string
// 	for _, objectType := range *objectsTypes {
// 		names = append(names, objectType.Name())
// 	}
// 	return names
// }

func GetObjectInstanceFromBytes(data []byte, typ reflect.Type) (interface{}, error) {
	instancePtr := reflect.New(typ).Interface() // Create a new pointer to a type instance

	err := json.Unmarshal(data, instancePtr) // Unmarshal into the pointer
	if err != nil {
		return nil, err
	}

	return instancePtr, nil
}
