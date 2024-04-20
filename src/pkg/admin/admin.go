package admin

import (
	"embed"

	"github.com/gofiber/fiber/v2"
	"github.com/pelusa-v/gorm-admin/src/pkg/handlers"
	"gorm.io/gorm"
)

//go:embed templates/*
var AdminTemplates embed.FS

type Admin struct {
	Handler handlers.AppHandler
}

func NewFiberAdmin(app *fiber.App, db *gorm.DB) *Admin {
	handler := &handlers.FiberHandler{
		App: app,
		BaseHandler: handlers.BaseHandler{
			TemplatesFs: AdminTemplates,
			GormDB:      db,
		},
	}
	admin := &Admin{Handler: handler}
	return admin
}

func NewGinAdmin(app *string, db *gorm.DB) *Admin {
	handler := &handlers.GinHandler{
		App: app,
		BaseHandler: handlers.BaseHandler{
			TemplatesFs: AdminTemplates,
			GormDB:      db,
		},
	}
	admin := &Admin{Handler: handler}
	return admin
}

func NewAdmin(db *gorm.DB) *Admin {
	handler := &handlers.BuiltInHandler{
		BaseHandler: handlers.BaseHandler{
			TemplatesFs: AdminTemplates,
			GormDB:      db,
		},
	}
	admin := &Admin{Handler: handler}
	return admin
}

func (admin *Admin) Register() {
	admin.Handler.Register()
}

func (admin *Admin) RegisteModel(model string) {
	admin.Handler.RegisterModel(model)
}
