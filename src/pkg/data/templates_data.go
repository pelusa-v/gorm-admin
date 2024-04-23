package data

import (
	"fmt"
	"reflect"
)

type ModelDetailPageData struct {
	Model              string
	ModelObjects       []ModelObject
	ModelObjectsFields []reflect.StructField
}

type ModelObject struct {
	Fields       []reflect.StructField
	FieldsValues []reflect.Value
	DetailURL    string
}

type HomePageData struct {
	Models []ModelsListItemData
}

type ModelsListItemData struct {
	Name string
}

func (item *ModelsListItemData) DetailURL() string {
	return fmt.Sprintf("/admin/%s", item.Name)
}

func GetHomePageData(modelTypes *[]reflect.Type) HomePageData {
	data := HomePageData{}
	modelsList := make([]ModelsListItemData, len(*modelTypes))
	for i, modelType := range *modelTypes {
		model := ModelsListItemData{Name: modelType.Name()}
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
		objectValue := reflect.ValueOf(o)
		var objectFields []reflect.StructField
		var objectFieldsValues []reflect.Value

		for i := 0; i < objectValue.NumField(); i++ {
			objectFields = append(objectFields, objectValue.Type().Field(i))
			objectFieldsValues = append(objectFieldsValues, objectValue.Field(i))
		}

		modelObjects = append(modelObjects, ModelObject{
			Fields:       objectFields,
			FieldsValues: objectFieldsValues,
		})
	}
	data.ModelObjects = modelObjects

	return data
}
