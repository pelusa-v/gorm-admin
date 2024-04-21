package admin

import (
	"embed"
	"html/template"
	"reflect"

	"github.com/gofiber/fiber/v2"
	"github.com/pelusa-v/gorm-admin/src/pkg/handlers"
	"gorm.io/gorm"
)

//go:embed templates/*
var AdminTemplates embed.FS

type Admin struct {
	Handler     handlers.AppHandler
	TemplatesFs embed.FS
}

func NewFiberAdmin(app *fiber.App, db *gorm.DB) *Admin {
	handler := &handlers.FiberHandler{
		App: app,
		BaseHandler: handlers.BaseHandler{
			GormDB: db,
		},
	}
	admin := &Admin{Handler: handler, TemplatesFs: AdminTemplates}
	return admin
}

func NewGinAdmin(app *string, db *gorm.DB) *Admin {
	handler := &handlers.GinHandler{
		App: app,
		BaseHandler: handlers.BaseHandler{
			GormDB: db,
		},
	}
	admin := &Admin{Handler: handler, TemplatesFs: AdminTemplates}
	return admin
}

func NewAdmin(db *gorm.DB) *Admin {
	handler := &handlers.BuiltInHandler{
		BaseHandler: handlers.BaseHandler{
			GormDB: db,
		},
	}
	admin := &Admin{Handler: handler, TemplatesFs: AdminTemplates}
	return admin
}

func (admin *Admin) template(templateFsPath string) *template.Template {
	tmpl, err := template.ParseFS(admin.TemplatesFs, templateFsPath)
	if err != nil {
		panic(err)
	}

	return tmpl
}

func (admin *Admin) Register() {
	admin.Handler.RegisterHomePage(admin.template("templates/home.html"))
}

func (admin *Admin) RegisterModel(model any) {

	modelType := reflect.TypeOf(model) // Add models validation (validate that is a db model, validate against db)
	admin.Handler.RegisterModel(modelType)
	admin.Handler.RegisterModelDetailPage(modelType, admin.template("templates/ModelDetail.html"))
}
