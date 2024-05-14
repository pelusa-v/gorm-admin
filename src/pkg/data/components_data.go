package data

import (
	"fmt"
	"reflect"
)

type SideBarData struct {
	AdminName string
	Models    []Model
}

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

type ModelObjectListItem struct {
	ModelObject ModelObject
	// Pk                    interface{}
	// Fields                []reflect.StructField
	// FieldsValues          []reflect.Value
	DetailURL             string
	DeleteURL             string
	DeleteObjectModalData DeleteObjectModalData
}

type DeleteObjectModalData struct {
	ModalId        string
	CloseModalId   string
	OpenModalId    string
	DeleteButtonId string
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
		// DetailURL:    fmt.Sprintf("/admin/%s/%v", objectValue.Type().Name(), objectValue.FieldByName(pkField.Name)),
		// DeleteURL:    fmt.Sprintf("/admin/%s/actions/delete/%v", objectValue.Type().Name(), objectValue.FieldByName(pkField.Name)),
		Pk: objectValue.FieldByName(pkField.Name),
	}
}
