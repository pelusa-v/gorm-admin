package data

import (
	"fmt"
	"reflect"
)

type Model struct {
	Name string
}

func (item *Model) DetailURL() string {
	return fmt.Sprintf("/admin/%s", item.Name)
}

type ModelObject struct {
	Pk           interface{}
	Fields       []reflect.StructField
	FieldsValues []reflect.Value
	TypeName     string
}

func MapModelObject(o interface{}) ModelObject {
	objectValue := reflect.ValueOf(o)
	if objectValue.Kind() == reflect.Ptr {
		objectValue = objectValue.Elem()
	}

	objectFields := GetObjectFields(objectValue.Type())
	objectFieldsValues := GetObjectFieldsValues(objectValue)
	pkField := FindPkField(objectFields)

	return ModelObject{
		Fields:       objectFields,
		FieldsValues: objectFieldsValues,
		TypeName:     objectValue.Type().Name(),
		Pk:           objectValue.FieldByName(pkField.Name),
	}
}
