package handlers

import (
	"html/template"
	"io/fs"
	"reflect"
)

type AppHandler interface {
	RegisterSimplePage(tmpl *template.Template, templateName string, route string, tmplDataFunc func() any)
	RegisterPkPage(tmpl *template.Template, templateName string, route string, tmplDataFunc func(pk string) any)
	RegisterStatic(fs fs.FS)
	RegisterCreateEndpoint(route string, typeToCreate reflect.Type, actionCreateFunc func(data interface{}) error)
	// RegisterHomePage(tmpl *template.Template)
	// RegisterModelDetailPage(modelType reflect.Type, tmpl *template.Template)
}

type BaseHandler struct {
}

type RequestMethod struct {
	Name string
}

var GET RequestMethod = RequestMethod{Name: "GET"}
var POST RequestMethod = RequestMethod{Name: "POST"}
var PUT RequestMethod = RequestMethod{Name: "PUT"}
var DELETE RequestMethod = RequestMethod{Name: "DELETE"}
