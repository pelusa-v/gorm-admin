package admin

import (
	"embed"
	"html/template"
	"reflect"

	"github.com/gofiber/fiber/v2"
	"github.com/pelusa-v/gorm-admin/src/pkg/handlers"
	"gorm.io/gorm"
)

// static/bootstrap/dist/css/bootstrap.min.css
// //go:embed static/bootstrap/*

//go:embed templates/*
//go:embed static/styles/*
//go:embed static/js/*
var AdminTemplates embed.FS

type Admin struct {
	Handler     handlers.AppHandler
	TemplatesFs embed.FS
	Models      []reflect.Type
	GormDB      *gorm.DB
	Name        string
}

func NewFiberAdmin(app *fiber.App, db *gorm.DB) *Admin {
	handler := &handlers.FiberHandler{
		App:         app,
		BaseHandler: handlers.BaseHandler{},
	}
	admin := &Admin{Handler: handler, TemplatesFs: AdminTemplates, GormDB: db}
	return admin
}

func NewGinAdmin(app *string, db *gorm.DB) *Admin {
	handler := &handlers.GinHandler{
		App:         app,
		BaseHandler: handlers.BaseHandler{},
	}
	admin := &Admin{Handler: handler, TemplatesFs: AdminTemplates, GormDB: db}
	return admin
}

func NewAdmin(db *gorm.DB) *Admin {
	handler := &handlers.BuiltInHandler{
		BaseHandler: handlers.BaseHandler{},
	}
	admin := &Admin{Handler: handler, TemplatesFs: AdminTemplates, GormDB: db}
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
	admin.registerHomePage()
}

func (admin *Admin) RegisterModel(model any) {
	modelType := reflect.TypeOf(model) // Add models validation (validate that is a db model, validate against db)
	admin.Models = append(admin.Models, modelType)

	admin.registerModelDetailPage(modelType)
	admin.registerModelObjectDetailPage(modelType)
	admin.registerModelObjectCreatePage(modelType)
	admin.registerModelObjectCreateEndpoint(modelType)
	admin.registerModelObjectDeleteEndpoint(modelType)
}

func (admin *Admin) Configure(name string) {
	admin.Name = name
}
