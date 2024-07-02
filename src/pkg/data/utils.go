package data

import (
	"encoding/json"
	"fmt"
	"reflect"
	"regexp"
	"strings"
	"time"

	"gorm.io/gorm"
)

const HTML_INPUT_TEXT_TYPE = "text"
const HTML_INPUT_EMAIL_TYPE = "email"
const HTML_INPUT_DATE_TYPE = "date"
const HTML_INPUT_DATETIME_TYPE = "datetime-local"
const HTML_INPUT_BOOL_TYPE = "checkbox"
const HTML_INPUT_NUMBER_TYPE = "number"

const GORM_DEFAULT_TAG_NAME string = "gorm"
const GORM_PK_DEFAULT_TAG_VALUE string = "primaryKey"
const GORM_PK_DEFAULT_NAME string = "ID"
const GORM_EMBEDDED_DEFAULT_TAG_VALUE string = "embedded"
const GORM_EMBEDDED_PREFIX_DEFAULT_TAG_VALUE string = "embeddedPrefix"

// FieldHasEmbeddedStructs checks if a struct field has embedded structs.
func FieldHasEmbeddedStructs(f reflect.StructField) bool {
	return f.Anonymous ||
		strings.Contains(f.Tag.Get(GORM_DEFAULT_TAG_NAME), GORM_EMBEDDED_DEFAULT_TAG_VALUE)
}

// FieldHasEmbeddedPrefix checks if a struct field has an embedded prefix.
func FieldHasEmbeddedPrefix(f reflect.StructField) bool {
	return strings.Contains(f.Tag.Get(GORM_DEFAULT_TAG_NAME), GORM_EMBEDDED_PREFIX_DEFAULT_TAG_VALUE)
}

// IsPkField checks if a struct field is a primary key field.
func IsPkField(f reflect.StructField) bool {
	return f.Tag.Get(GORM_DEFAULT_TAG_NAME) == GORM_PK_DEFAULT_TAG_VALUE || f.Name == GORM_PK_DEFAULT_NAME
}

// IsVirtualField checks if a struct field is a virtual field (has one, belongs to, has many relations).
func IsVirtualField(f reflect.StructField, allTypes *[]reflect.Type) bool {
	for _, t := range *allTypes {
		if f.Type.Name() == t.Name() {
			return true
		}

		if f.Type.Kind() == reflect.Slice {
			if f.Type.Elem().Name() == t.Name() {
				return true
			}
		}
	}
	return false
}

// AddEmbeddedPrefixToField adds an embedded prefix to a struct field.
func AddEmbeddedPrefixToField(f *reflect.StructField) {
	re := regexp.MustCompile(fmt.Sprintf(`%s:([^";]+)`, GORM_EMBEDDED_PREFIX_DEFAULT_TAG_VALUE))
	matches := re.FindStringSubmatch(f.Tag.Get(GORM_DEFAULT_TAG_NAME))
	if len(matches) > 1 {
		f.Name = matches[1]
	}
}

// GetHtmlInputType returns the HTML input type for a struct field.
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

// FindPkField finds the primary key field in a list of struct fields.
func FindPkField(fields []reflect.StructField) reflect.StructField {
	for _, f := range fields {
		if IsPkField(f) {
			return f
		}
	}

	panic("Gorm model doesn't have PK")
}

// GetObjectFieldsValues returns the values of the fields in an object.
func GetObjectFieldsValues(objectValue reflect.Value, allTypes *[]reflect.Type) []reflect.Value {
	if objectValue.Kind() == reflect.Ptr {
		objectValue = objectValue.Elem()
	}

	var objectFieldsValues []reflect.Value

	objectType := objectValue.Type()
	for i := 0; i < objectValue.NumField(); i++ {
		fieldType := objectType.Field(i)
		fieldValue := objectValue.Field(i)

		if IsVirtualField(fieldType, allTypes) {
			continue
		}

		if FieldHasEmbeddedStructs(fieldType) {
			embeddedFieldsValues := GetObjectFieldsValues(fieldValue, allTypes)
			objectFieldsValues = append(objectFieldsValues, embeddedFieldsValues...)
		} else {
			objectFieldsValues = append(objectFieldsValues, fieldValue)
		}
	}

	return objectFieldsValues
}

// GetObjectFields returns the fields of an object.
func GetObjectFields(objectType reflect.Type, allTypes *[]reflect.Type) []reflect.StructField {
	var objectFields []reflect.StructField

	for i := 0; i < objectType.NumField(); i++ {
		fieldType := objectType.Field(i)

		// if FieldHasEmbeddedPrefix(fieldType) {
		// 	AddEmbeddedPrefixToField(&fieldType)
		// }

		if IsVirtualField(fieldType, allTypes) {
			// fmt.Println("---------- WATCH HERE!!! ----------")
			// fmt.Println(fieldType.Type.Name())
			// fmt.Println(fieldType.Type.Kind())
			// if fieldType.Type.Kind() == reflect.Slice {
			// 	fmt.Println(fieldType.Type.Elem())
			// 	fmt.Println(fieldType.Type.Elem().Name())
			// }
			// fmt.Println(fieldType.Type)
			// fmt.Println(fieldType.Name)
			// fmt.Println("*****************************")
			continue
		}

		if FieldHasEmbeddedStructs(fieldType) {
			embeddedFields := GetObjectFields(fieldType.Type, allTypes)
			objectFields = append(objectFields, embeddedFields...)
		} else {
			objectFields = append(objectFields, fieldType)
		}
	}

	return objectFields
}

// GetObjectInstanceFromBytes returns an object instance from JSON bytes.
func GetObjectInstanceFromBytes(data []byte) (interface{}, error) {
	// instancePtr := reflect.New(typ).Interface() // Create a new pointer to a type instance

	var instanceData map[string]interface{}
	err := json.Unmarshal(data, &instanceData) // Unmarshal into the pointer
	if err != nil {
		return nil, err
	}

	return instanceData, nil
}
