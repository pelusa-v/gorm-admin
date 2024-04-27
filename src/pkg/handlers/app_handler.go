package handlers

import (
	"html/template"
	"io/fs"
)

type AppHandler interface {
	RegisterSimplePage(tmpl *template.Template, templateName string, route string, tmplDataFunc func() any)
	RegisterPkPage(tmpl *template.Template, templateName string, route string, tmplDataFunc func(pk string) any)
	RegisterStatic(fs fs.FS)
	// RegisterHomePage(tmpl *template.Template)
	// RegisterModelDetailPage(modelType reflect.Type, tmpl *template.Template)
}

type BaseHandler struct {
}
