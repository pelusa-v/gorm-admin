package handlers

import (
	"html/template"
	"io/fs"
)

type BuiltInHandler struct {
	BaseHandler
}

// func (handler *BuiltInHandler) Register() {
// 	fmt.Println("Registering admin in BuiltIn http app")
// }

func (handler *BuiltInHandler) RegisterSimplePage(tmpl *template.Template, templateName string, route string, tmplDataFunc func() any) {

}

func (handler *BuiltInHandler) RegisterPkPage(tmpl *template.Template, templateName string, route string, tmplDataFunc func(pk string) any) {
}

func (handler *BuiltInHandler) RegisterStatic(fs fs.FS) {
}

func (handler *BuiltInHandler) RegisterCreateEndpoint(route string, actionCreateFunc func(data interface{}) error) {
}

func (handler *BuiltInHandler) RegisterDeleteEndpoint(route string, actionFunc func(pk interface{}) error) {
}
