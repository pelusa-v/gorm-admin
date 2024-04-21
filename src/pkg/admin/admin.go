package admin

import (
	"embed"
	"fmt"
	"html/template"
	"reflect"

	"github.com/gofiber/fiber/v2"
	"github.com/pelusa-v/gorm-admin/src/pkg/data"
	"github.com/pelusa-v/gorm-admin/src/pkg/handlers"
	"gorm.io/gorm"
)

//go:embed templates/*
var AdminTemplates embed.FS

type Admin struct {
	Handler     handlers.AppHandler
	TemplatesFs embed.FS
	Models      []reflect.Type
	GormDB      *gorm.DB
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
	// admin.Handler.RegisterHomePage(admin.template("templates/home.html"))

	homePageTemplate := admin.template("templates/home.html")
	admin.Handler.RegisterPage(homePageTemplate, "/admin", func() any {
		return data.GetHomePageData(&admin.Models)
	})
}

func (admin *Admin) RegisterModel(model any) {
	modelType := reflect.TypeOf(model) // Add models validation (validate that is a db model, validate against db)
	admin.Models = append(admin.Models, modelType)

	// admin.Handler.RegisterModelDetailPage(modelType, admin.template("templates/ModelDetail.html"))

	dbModel := data.NewDbModel(modelType, admin.GormDB)
	modelDetailPageTemplate := admin.template("templates/ModelDetail.html")
	modelDetailPageRoute := fmt.Sprintf("/admin/%s", modelType.Name())
	admin.Handler.RegisterPage(modelDetailPageTemplate, modelDetailPageRoute, func() any {
		return data.GetModelDetailPageData(*dbModel)
	})
}
