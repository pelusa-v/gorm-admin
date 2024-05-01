package data

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
}

type SelectInput struct {
	Options []SelectInputOption
}

type SelectInputOption struct {
	Label string
	Name  string
	Value string
}

func (form *FormData) SetFormInputs(model *DbModel) {
	fields := GetObjectFields(model.modelType)
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
