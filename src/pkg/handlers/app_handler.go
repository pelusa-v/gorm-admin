package handlers

import (
	"html/template"
)

type AppHandler interface {
	RegisterPage(tmpl *template.Template, route string, tmplDataFunc func() any)
	// RegisterHomePage(tmpl *template.Template)
	// RegisterModelDetailPage(modelType reflect.Type, tmpl *template.Template)
}

type BaseHandler struct {
}
