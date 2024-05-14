package data

import "reflect"

type HomePageData struct {
	SideBarData
}

type ModelDetailPageData struct {
	SideBarData
	Model                string
	ModelObjectListItems []ModelObjectListItem
	ModelObjectsFields   []reflect.StructField
	PreviousURL          string
	AddURL               string
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
