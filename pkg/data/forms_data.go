package data

import (
	"reflect"
)

type FormData struct {
	SimpleInputs []SimpleInput
	SelectInputs []SelectInput
}

type SimpleInput struct {
	Id       string
	Label    string
	Name     string
	Type     string
	Disabled bool
	Required bool
	Value    interface{}
}

type SelectInput struct {
	Options []SelectInputOption
}

type SelectInputOption struct {
	Label string
	Name  string
	Value string
}

func (form *FormData) SetFormInputs(model *DbModel, allTypes *[]reflect.Type) {
	fields := GetObjectFields(model.modelType, allTypes)
	for _, f := range fields {
		input := SimpleInput{}
		input.Id = f.Name
		input.Name = f.Name
		input.Label = f.Name
		// input.Disabled = IsPkField(f)
		input.Type = GetHtmlInputType(f)
		form.SimpleInputs = append(form.SimpleInputs, input)

		// fmt.Println(f.Tag)
		// fmt.Println(f.Name)
		// fmt.Println(f.Index)
		// fmt.Println(f.Type)
		// fmt.Println("---------------")
	}
}

func (form *FormData) SetFormInputsValues(model *DbModel, allTypes *[]reflect.Type, modelEntity interface{}) {
	fields := GetObjectFields(model.modelType, allTypes)

	entity := reflect.ValueOf(modelEntity)
	if entity.Kind() == reflect.Ptr {
		entity = entity.Elem()
	}

	for _, f := range fields {
		input := SimpleInput{}
		input.Id = f.Name
		input.Name = f.Name
		input.Label = f.Name
		input.Type = GetHtmlInputType(f)

		// fmt.Println("************ WATCH HERE ***********")
		// fmt.Println(entity.FieldByName(f.Name).String())
		// fmt.Println(entity.FieldByName(f.Name))
		// fmt.Println(entity.FieldByName(f.Name))
		// fmt.Println("***********************************")

		input.Value = entity.FieldByName(f.Name)

		// input.Value = entity.Elem().FieldByName(f.Name).Elem().String()
		form.SimpleInputs = append(form.SimpleInputs, input)
	}
}
