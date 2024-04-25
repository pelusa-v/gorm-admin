package data

import (
	"fmt"
	"reflect"
)

type HomePageData struct {
	Models []Model
}

type ModelDetailPageData struct {
	Model              string
	ModelObjects       []ModelObject
	ModelObjectsFields []reflect.StructField
}

type ModelObjectDetailPageData struct {
	Model       string
	ModelObject ModelObject
}

type ModelObject struct {
	Fields       []reflect.StructField
	FieldsValues []reflect.Value
	DetailURL    string
}

type Model struct {
	Name string
}

func (item *Model) DetailURL() string {
	return fmt.Sprintf("/admin/%s", item.Name)
}

func GetHomePageData(modelTypes *[]reflect.Type) HomePageData {
	data := HomePageData{}
	modelsList := make([]Model, len(*modelTypes))
	for i, modelType := range *modelTypes {
		model := Model{Name: modelType.Name()}
		modelsList[i] = model
	}
	data.Models = modelsList
	return data
}

func GetModelDetailPageData(model DbModel) ModelDetailPageData {
	modelFields := GetObjectFields(model.modelType)
	data := ModelDetailPageData{Model: model.modelType.Name(), ModelObjectsFields: modelFields}

	var modelObjects []ModelObject
	objects := model.ListObjects()
	for _, o := range objects {
		modelObject := MapModelObject(o)
		modelObjects = append(modelObjects, modelObject)
	}
	data.ModelObjects = modelObjects

	return data
}

func GetModelObjectDetailPageData(model DbModel, pk string) ModelObjectDetailPageData {
	object := model.GetObject(pk)
	data := ModelObjectDetailPageData{Model: model.modelType.Name(), ModelObject: MapModelObject(object)}
	return data
}

func MapModelObject(o interface{}) ModelObject {
	objectValue := reflect.ValueOf(o)
	if objectValue.Kind() == reflect.Ptr {
		objectValue = objectValue.Elem()
	}

	objectFields := GetObjectFields(objectValue.Type())
	objectFieldsValues := GetObjectFieldsValues(objectValue)

	fmt.Printf("objectFields: %d\n", len(objectFields))
	fmt.Printf("objectFieldsValues: %d\n", len(objectFieldsValues))

	for i := 0; i < len(objectFields); i++ {
		fmt.Printf("%v - %v (%v)\n", objectFields[i].Name, objectFields[i].Type, objectFields[i].Tag)
	}
	fmt.Println("------------------------------------")

	for i := 0; i < len(objectFieldsValues); i++ {
		fmt.Printf("%v\n", objectFieldsValues[i])
	}

	return ModelObject{
		Fields:       objectFields,
		FieldsValues: objectFieldsValues,
	}
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

// func GetFieldsAndValues(objectValue reflect.Value) ([]reflect.StructField, []reflect.Value) {
// 	if objectValue.Kind() == reflect.Ptr {
// 		objectValue = objectValue.Elem()
// 	}

// 	var objectFields []reflect.StructField
// 	var objectFieldsValues []reflect.Value

// 	objectType := objectValue.Type()
// 	for i := 0; i < objectValue.NumField(); i++ {
// 		fieldType := objectType.Field(i)
// 		fieldValue := objectValue.Field(i)

// 		if FieldHasEmbeddedStructs(fieldType) {
// 			embeddedFields, embeddedFieldsValues := GetFieldsAndValues(fieldValue)
// 			objectFields = append(objectFields, embeddedFields...)
// 			objectFieldsValues = append(objectFieldsValues, embeddedFieldsValues...)
// 		} else {
// 			objectFields = append(objectFields, fieldType)
// 			objectFieldsValues = append(objectFieldsValues, fieldValue)
// 		}
// 	}

// 	return objectFields, objectFieldsValues
// }
