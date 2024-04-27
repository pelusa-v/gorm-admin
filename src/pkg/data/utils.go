package data

import (
	"reflect"

	"gorm.io/gorm"
)

const GORM_PK_DEFAULT_TAG_NAME string = "gorm"
const GORM_PK_DEFAULT_TAG_VALUE string = "primaryKey"
const GORM_PK_DEFAULT_NAME string = "ID"

func FieldHasEmbeddedStructs(f reflect.StructField) bool {
	return f.Anonymous || f.Type == reflect.TypeOf(gorm.DeletedAt{})
}

func FindPkField(fields []reflect.StructField) reflect.StructField {
	for _, f := range fields {
		if f.Tag.Get(GORM_PK_DEFAULT_TAG_NAME) == GORM_PK_DEFAULT_TAG_VALUE || f.Name == GORM_PK_DEFAULT_NAME {
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
