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
