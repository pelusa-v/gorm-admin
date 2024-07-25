package data

type SideBarData struct {
	AdminName string
	Models    []Model
}

type ModelObjectListItem struct {
	ModelObject           ModelObject
	DetailURL             string
	DeleteURL             string
	UpdateURL             string
	DeleteObjectModalData DeleteObjectModalData
}

type DeleteObjectModalData struct {
	ModalId        string
	CloseModalId   string
	OpenModalId    string
	DeleteButtonId string
}
