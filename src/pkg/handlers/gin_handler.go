package handlers

import (
	"html/template"
	"reflect"
)

type GinHandler struct {
	BaseHandler
	App *string
}

// func (handler *GinHandler) Register() {
// 	fmt.Println("Registering admin in Gin app")
// }

func (handler *GinHandler) RegisterHomePage(tmpl *template.Template) {

}

func (handler *GinHandler) RegisterModelDetailPage(modelType reflect.Type, tmpl *template.Template) {
}
