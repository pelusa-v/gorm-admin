package data

import (
	"fmt"
	"reflect"
)

type SideBarData struct {
	AdminName string
	Models    []Model
}

type HomePageData struct {
	SideBarData
}

type ModelDetailPageData struct {
	SideBarData
	Model              string
	ModelObjects       []ModelObject
	ModelObjectsFields []reflect.StructField
	PreviousURL        string
	AddURL             string
}

type ModelObjectDetailPageData struct {
	SideBarData
	Model       string
	ModelObject ModelObject
	PreviousURL string
}

type ModelObjectCreatePageData struct {
	SideBarData
	Model                  string
	PreviousURL            string
	CreateObjectForm       FormData
	CreateObjectURL        string
	RedirectAfterCreateURL string
}

type ModelObject struct {
	Pk                    interface{}
	Fields                []reflect.StructField
	FieldsValues          []reflect.Value
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

type Model struct {
	Name string
}

func (item *Model) DetailURL() string {
	return fmt.Sprintf("/admin/%s", item.Name)
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
	modelFields := GetObjectFields(modelType)
	data := ModelDetailPageData{Model: modelType.Name(), ModelObjectsFields: modelFields, PreviousURL: "/admin",
		AddURL: fmt.Sprintf("/admin/%s/actions/create", model.modelType.Name())}
	data.Models = manager.GetSidebarModels()
	data.AdminName = manager.GetSidebarName()

	var modelObjects []ModelObject
	objects := model.ListObjects()
	for _, o := range objects {
		modelObject := MapModelObject(o)
		deleteModalData := DeleteObjectModalData{
			ModalId:        fmt.Sprintf("delete-modal-%v", modelObject.Pk),
			CloseModalId:   fmt.Sprintf("close-delete-modal-%v", modelObject.Pk),
			OpenModalId:    fmt.Sprintf("open-delete-modal-%v", modelObject.Pk),
			DeleteButtonId: fmt.Sprintf("delete-action-%v", modelObject.Pk),
		}
		modelObject.DeleteObjectModalData = deleteModalData
		modelObjects = append(modelObjects, modelObject)
	}
	data.ModelObjects = modelObjects

	return data
}

func (manager *TemplateManager) GetModelObjectDetailPageData(model DbModel, pk string) ModelObjectDetailPageData {
	object := model.GetObject(pk)
	data := ModelObjectDetailPageData{Model: model.modelType.Name(), ModelObject: MapModelObject(object),
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
	templateForm.SetFormInputs(&model)

	data := ModelObjectCreatePageData{Model: model.modelType.Name(), PreviousURL: fmt.Sprintf("/admin/%s", model.modelType.Name()),
		CreateObjectURL: fmt.Sprintf("/admin/%s/actions/create", model.modelType.Name()), CreateObjectForm: *templateForm,
		RedirectAfterCreateURL: fmt.Sprintf("/admin/%s", model.modelType.Name())}
	data.Models = manager.GetSidebarModels()
	data.AdminName = manager.GetSidebarName()
	return data
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
		DetailURL:    fmt.Sprintf("/admin/%s/%v", objectValue.Type().Name(), objectValue.FieldByName(pkField.Name)),
		DeleteURL:    fmt.Sprintf("/admin/%s/actions/delete/%v", objectValue.Type().Name(), objectValue.FieldByName(pkField.Name)),
		Pk:           objectValue.FieldByName(pkField.Name),
	}
}
