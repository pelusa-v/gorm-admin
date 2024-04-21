package data

import (
	"bytes"
	"fmt"
	"html/template"
	"reflect"
)

type ModelDetailPageData struct {
	Model   string
	Objects []interface{}
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

func GetHomePage(modelTypes []reflect.Type, tmpl *template.Template) bytes.Buffer {
	data := HomePageData{}
	modelsList := make([]ModelsListItemData, len(modelTypes))
	for i, modelType := range modelTypes {
		model := ModelsListItemData{Name: modelType.Name()}
		modelsList[i] = model
	}
	data.Models = modelsList

	var tmplOutput bytes.Buffer
	err := tmpl.Execute(&tmplOutput, data)
	if err != nil {
		panic(err)
	}

	return tmplOutput
}

func GetModelDetailPage(model DbModel, tmpl *template.Template) bytes.Buffer {
	var tmplOutput bytes.Buffer
	data := ModelDetailPageData{Model: model.modelType.Name(), Objects: model.ListOjects()}
	err := tmpl.Execute(&tmplOutput, data)
	if err != nil {
		panic(err)
	}

	return tmplOutput
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
	data := ModelDetailPageData{Model: model.modelType.Name(), Objects: model.ListOjects()}
	return data
}
