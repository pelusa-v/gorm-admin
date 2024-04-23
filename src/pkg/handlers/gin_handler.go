package handlers

import (
	"html/template"
)

type GinHandler struct {
	BaseHandler
	App *string
}

//	func (handler *GinHandler) Register() {
//		fmt.Println("Registering admin in Gin app")
//	}
func (handler *GinHandler) RegisterSimplePage(tmpl *template.Template, route string, tmplDataFunc func() any) {

}

func (handler *GinHandler) RegisterPkPage(tmpl *template.Template, route string, tmplDataFunc func(pk string) any) {
}
