package handlers

import (
	"embed"
	"reflect"

	"gorm.io/gorm"
)

type AppHandler interface {
	Register()
	RegisterModel(model reflect.Type)
}

type BaseHandler struct {
	GormDB      *gorm.DB
	TemplatesFs embed.FS
}
