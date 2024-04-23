package handlers

import (
	"html/template"
)

type AppHandler interface {
	RegisterSimplePage(tmpl *template.Template, route string, tmplDataFunc func() any)
	RegisterPkPage(tmpl *template.Template, route string, tmplDataFunc func(pk string) any)
	// RegisterHomePage(tmpl *template.Template)
	// RegisterModelDetailPage(modelType reflect.Type, tmpl *template.Template)
}

type BaseHandler struct {
}
