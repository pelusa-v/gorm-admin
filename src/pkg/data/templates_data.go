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

// func (item *ModelObject) DetailObjectDetailURL() string {
// 	return fmt.Sprintf("/admin/%s/%s", item.Name)
// }

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
	data := ModelDetailPageData{Model: model.modelType.Name(), ModelObjectsFields: model.GetModelFields()}

	var modelObjects []ModelObject
	objects := model.ListObjects()
	for _, o := range objects {
		modelObject := GetObjectFields(o)
		modelObjects = append(modelObjects, modelObject)
	}
	data.ModelObjects = modelObjects

	return data
}

func GetModelObjectDetailPageData(model DbModel, pk string) ModelObjectDetailPageData {
	object := model.GetObject(pk)
	data := ModelObjectDetailPageData{Model: model.modelType.Name(), ModelObject: GetObjectFields(object)}
	return data
}

func GetObjectFields(o interface{}) ModelObject {
	objectValue := reflect.ValueOf(o).Elem()
	var objectFields []reflect.StructField
	var objectFieldsValues []reflect.Value

	for i := 0; i < objectValue.NumField(); i++ {
		objectFields = append(objectFields, objectValue.Type().Field(i))
		objectFieldsValues = append(objectFieldsValues, objectValue.Field(i))
	}

	for i := 0; i < objectValue.NumField(); i++ {
		fmt.Println("---------------------------")
		fmt.Printf("NAME: %v :\n", objectValue.Type().Field(i).Name)
		fmt.Printf("TYPE: %v :\n", objectValue.Type().Field(i).Type)
		fmt.Printf("TAG: %v :\n", objectValue.Type().Field(i).Tag)
		fmt.Printf("%v :\n", objectValue.Type().Field(i))
		fmt.Println(objectValue.Field(i))
	}

	return ModelObject{
		Fields:       objectFields,
		FieldsValues: objectFieldsValues,
	}
}

// func GetPkField(o interface{}) ModelObject {

// }
