package data

import (
	"fmt"
	"reflect"
)

type ModelDetailPageData struct {
	Model   string
	Objects []DbObjectInstance
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
	data := ModelDetailPageData{Model: model.modelType.Name(), Objects: model.ListObjects()}
	return data
}
