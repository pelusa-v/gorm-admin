package handlers

import (
	"embed"

	"gorm.io/gorm"
)

type AppHandler interface {
	Register()
	RegisterModel(model string)
}

type BaseHandler struct {
	GormDB      *gorm.DB
	TemplatesFs embed.FS
}
