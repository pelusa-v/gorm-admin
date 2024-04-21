package handlers

import (
	"html/template"
	"reflect"

	"gorm.io/gorm"
)

type AppHandler interface {
	// Register()
	RegisterModel(model reflect.Type)
	RegisterHomePage(tmpl *template.Template)
	RegisterModelDetailPage(modelType reflect.Type, tmpl *template.Template)
}

type BaseHandler struct {
	GormDB *gorm.DB
	Models []reflect.Type
}

func (handler *BaseHandler) RegisterModel(model reflect.Type) {
	handler.Models = append(handler.Models, model)
}
