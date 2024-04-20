package admin

import (
	"embed"

	"github.com/gofiber/fiber/v2"
	"github.com/pelusa-v/gorm-admin/src/pkg/handlers"
)

//go:embed templates/*
var AdminTemplates embed.FS

type Admin struct {
	Handler handlers.AppHandler
}

func NewFiberAdmin(app *fiber.App) *Admin {
	handler := &handlers.FiberHandler{App: app, TemplatesFs: AdminTemplates}
	admin := &Admin{Handler: handler}
	return admin
}

func NewGinAdmin(app *string) *Admin {
	handler := &handlers.GinHandler{App: app, TemplatesFs: AdminTemplates}
	admin := &Admin{Handler: handler}
	return admin
}

func NewAdmin() *Admin {
	handler := &handlers.BuiltInHandler{TemplatesFs: AdminTemplates}
	admin := &Admin{Handler: handler}
	return admin
}

func (admin *Admin) Register() {
	admin.Handler.Register()
}
