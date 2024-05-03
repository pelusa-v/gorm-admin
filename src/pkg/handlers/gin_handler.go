package handlers

import (
	"html/template"
	"io/fs"
	"reflect"
)

type GinHandler struct {
	BaseHandler
	App *string
}

//	func (handler *GinHandler) Register() {
//		fmt.Println("Registering admin in Gin app")
//	}
func (handler *GinHandler) RegisterSimplePage(tmpl *template.Template, templateName string, route string, tmplDataFunc func() any) {

}

func (handler *GinHandler) RegisterPkPage(tmpl *template.Template, templateName string, route string, tmplDataFunc func(pk string) any) {
}

func (handler *GinHandler) RegisterStatic(fs fs.FS) {
}

func (handler *GinHandler) RegisterCreateEndpoint(route string, typeToCreate reflect.Type, actionCreateFunc func(data interface{}) error) {
}
