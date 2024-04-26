package handlers

import (
	"html/template"
)

type BuiltInHandler struct {
	BaseHandler
}

// func (handler *BuiltInHandler) Register() {
// 	fmt.Println("Registering admin in BuiltIn http app")
// }

func (handler *BuiltInHandler) RegisterSimplePage(tmpl *template.Template, route string, tmplDataFunc func() any) {

}

func (handler *BuiltInHandler) RegisterPkPage(tmpl *template.Template, route string, tmplDataFunc func(pk string) any) {
}

func (handler *BuiltInHandler) RegisterStatic(staticFolder string, staticPath string) {
}
