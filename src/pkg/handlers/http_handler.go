package handlers

import (
	"html/template"
	"reflect"
)

type BuiltInHandler struct {
	BaseHandler
}

// func (handler *BuiltInHandler) Register() {
// 	fmt.Println("Registering admin in BuiltIn http app")
// }

func (handler *BuiltInHandler) RegisterPage(tmpl *template.Template, route string, tmplDataFunc func() any) {

}

func (handler *BuiltInHandler) RegisterHomePage(tmpl *template.Template) {

}

func (handler *BuiltInHandler) RegisterModelDetailPage(modelType reflect.Type, tmpl *template.Template) {
}
