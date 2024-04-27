package data

import "reflect"

type TemplateManager struct {
	configurableData *ConfigurableData
}

func NewTemplateManager(name *string, models *[]reflect.Type) *TemplateManager {
	return &TemplateManager{
		configurableData: &ConfigurableData{
			Name:   name,
			Models: models,
		},
	}
}

type ConfigurableData struct {
	Name   *string
	Models *[]reflect.Type
}
