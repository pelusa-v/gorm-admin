package data

import (
	"fmt"
	"reflect"
)

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

func (manager *TemplateManager) GetSidebarModels() []Model {
	modelsList := make([]Model, len(*manager.configurableData.Models))
	for i, modelType := range *manager.configurableData.Models {
		model := Model{Name: modelType.Name()}
		modelsList[i] = model
	}
	return modelsList
}

func (manager *TemplateManager) GetSidebarName() string {
	if *manager.configurableData.Name == "" {
		return "GORM admin"
	}
	return *manager.configurableData.Name
}

func (manager *TemplateManager) GetHomePageData() HomePageData {
	data := HomePageData{}
	data.Models = manager.GetSidebarModels()
	data.AdminName = manager.GetSidebarName()
	return data
}

func (manager *TemplateManager) GetModelDetailPageData(model DbModel) ModelDetailPageData {
	modelType := model.modelType
	modelFields := GetObjectFields(modelType, manager.configurableData.Models)
	data := ModelDetailPageData{Model: modelType.Name(), ModelObjectsFields: modelFields, PreviousURL: "/admin",
		AddURL: fmt.Sprintf("/admin/%s/actions/create", model.modelType.Name())}
	data.Models = manager.GetSidebarModels()
	data.AdminName = manager.GetSidebarName()

	var modelObjectsListItems []ModelObjectListItem
	objects := model.ListObjects()
	for _, o := range objects {
		modelObjectListItem := ModelObjectListItem{}
		modelObject := MapModelObject(o, manager.configurableData.Models)
		modelObjectListItem.ModelObject = modelObject
		modelObjectListItem.DetailURL = fmt.Sprintf("/admin/%s/%v", modelObject.TypeName, modelObject.Pk)
		modelObjectListItem.DeleteURL = fmt.Sprintf("/admin/%s/actions/delete/%v", modelObject.TypeName, modelObject.Pk)
		modelObjectListItem.UpdateURL = fmt.Sprintf("/admin/%s/actions/update/%v", modelObject.TypeName, modelObject.Pk)
		modelObjectListItem.DeleteObjectModalData = DeleteObjectModalData{
			ModalId:        fmt.Sprintf("delete-modal-%v", modelObject.Pk),
			CloseModalId:   fmt.Sprintf("close-delete-modal-%v", modelObject.Pk),
			OpenModalId:    fmt.Sprintf("open-delete-modal-%v", modelObject.Pk),
			DeleteButtonId: fmt.Sprintf("delete-action-%v", modelObject.Pk),
		}
		modelObjectsListItems = append(modelObjectsListItems, modelObjectListItem)
	}
	data.ModelObjectListItems = modelObjectsListItems

	return data
}

func (manager *TemplateManager) GetModelObjectDetailPageData(model DbModel, pk string) ModelObjectDetailPageData {
	object := model.GetObject(pk)
	data := ModelObjectDetailPageData{Model: model.modelType.Name(), ModelObject: MapModelObject(object, manager.configurableData.Models),
		PreviousURL: fmt.Sprintf("/admin/%s", model.modelType.Name())}
	data.Models = manager.GetSidebarModels()
	data.AdminName = manager.GetSidebarName()
	return data
}

func (manager *TemplateManager) GetModelObjectCreatePageData(model DbModel) ModelObjectCreatePageData {
	templateForm := &FormData{
		SimpleInputs: make([]SimpleInput, 0),
		SelectInputs: make([]SelectInput, 0),
	}
	templateForm.SetFormInputs(&model, manager.configurableData.Models)

	data := ModelObjectCreatePageData{Model: model.modelType.Name(), PreviousURL: fmt.Sprintf("/admin/%s", model.modelType.Name()),
		SubmitObjectURL: fmt.Sprintf("/admin/%s/actions/create", model.modelType.Name()), SubmitObjectForm: *templateForm,
		RedirectAfterCreateURL: fmt.Sprintf("/admin/%s", model.modelType.Name())}
	data.Models = manager.GetSidebarModels()
	data.AdminName = manager.GetSidebarName()
	return data
}

func (manager *TemplateManager) GetModelObjectUpdatePageData(model DbModel, pk string) ModelObjectCreatePageData {
	templateForm := &FormData{
		SimpleInputs: make([]SimpleInput, 0),
		SelectInputs: make([]SelectInput, 0),
	}
	// templateForm.SetFormInputs(&model, manager.configurableData.Models)
	object := model.GetObject(pk)
	templateForm.SetFormInputsValues(&model, manager.configurableData.Models, object) // TODO

	data := ModelObjectCreatePageData{Model: model.modelType.Name(), PreviousURL: fmt.Sprintf("/admin/%s", model.modelType.Name()),
		SubmitObjectURL: fmt.Sprintf("/admin/%s/actions/update", model.modelType.Name()), SubmitObjectForm: *templateForm,
		RedirectAfterCreateURL: fmt.Sprintf("/admin/%s", model.modelType.Name())}
	data.Models = manager.GetSidebarModels()
	data.AdminName = manager.GetSidebarName()
	return data
}
