package admin

import (
	"fmt"
	"io/fs"
	"reflect"

	"github.com/pelusa-v/gorm-admin/pkg/data"
)

func (admin *Admin) registerHomePage() {
	templates := admin.template("templates/*.gohtml")
	templateManager := data.NewTemplateManager(&admin.Name, &admin.Models)

	admin.Handler.RegisterSimplePage(templates, "home.gohtml", "/admin", func() any {
		return templateManager.GetHomePageData()
	})

	subFS, _ := fs.Sub(admin.TemplatesFs, "static/styles")
	admin.Handler.RegisterStatic(subFS)
}

func (admin *Admin) registerModelDetailPage(modelType reflect.Type) {
	dbModel := data.NewDbModel(modelType, admin.GormDB)
	templates := admin.template("templates/*.gohtml")
	templateManager := data.NewTemplateManager(&admin.Name, &admin.Models)
	modelDetailPageRoute := fmt.Sprintf("/admin/%s", modelType.Name())

	admin.Handler.RegisterSimplePage(templates, "ModelDetail.gohtml", modelDetailPageRoute, func() any {
		return templateManager.GetModelDetailPageData(*dbModel)
	})
}

func (admin *Admin) registerModelObjectDetailPage(modelType reflect.Type) {
	dbModel := data.NewDbModel(modelType, admin.GormDB)
	templates := admin.template("templates/*.gohtml")
	templateManager := data.NewTemplateManager(&admin.Name, &admin.Models)
	modelObjectDetailPageRoute := fmt.Sprintf("/admin/%s/:pk", modelType.Name())

	admin.Handler.RegisterPkPage(templates, "ModelObjectDetail.gohtml", modelObjectDetailPageRoute, func(pk string) any {
		return templateManager.GetModelObjectDetailPageData(*dbModel, pk)
	})
}

func (admin *Admin) registerModelObjectCreatePage(modelType reflect.Type) {
	dbModel := data.NewDbModel(modelType, admin.GormDB)
	templates := admin.template("templates/*.gohtml")
	templateManager := data.NewTemplateManager(&admin.Name, &admin.Models)
	modelObjectCreatePageRoute := fmt.Sprintf("/admin/%s/actions/create", modelType.Name())

	admin.Handler.RegisterSimplePage(templates, "ModelObjectCreate.gohtml", modelObjectCreatePageRoute, func() any {
		return templateManager.GetModelObjectCreatePageData(*dbModel)
	})
}

func (admin *Admin) registerModelObjectCreateEndpoint(modelType reflect.Type) {
	dbModel := data.NewDbModel(modelType, admin.GormDB)
	modelObjectCreateRoute := fmt.Sprintf("/admin/%s/actions/create", modelType.Name())
	admin.Handler.RegisterCreateEndpoint(modelObjectCreateRoute, func(data interface{}) error {
		return dbModel.CreateObject(data, modelType)
	})
}

func (admin *Admin) registerModelObjectUpdatePage(modelType reflect.Type) {
	dbModel := data.NewDbModel(modelType, admin.GormDB)
	templates := admin.template("templates/*.gohtml")
	templateManager := data.NewTemplateManager(&admin.Name, &admin.Models)
	modelObjectUpdatePageRoute := fmt.Sprintf("/admin/%s/actions/update/:pk", modelType.Name())

	admin.Handler.RegisterPkPage(templates, "ModelObjectCreate.gohtml", modelObjectUpdatePageRoute, func(pk string) any {
		return templateManager.GetModelObjectUpdatePageData(*dbModel, pk)
	})
}

func (admin *Admin) registerModelObjectUpdateEndpoint(modelType reflect.Type) {
	dbModel := data.NewDbModel(modelType, admin.GormDB)
	modelObjectUpdateRoute := fmt.Sprintf("/admin/%s/actions/update", modelType.Name())
	admin.Handler.RegisterCreateEndpoint(modelObjectUpdateRoute, func(data interface{}) error {
		return dbModel.UpdateObject(data, modelType)
	})
}

func (admin *Admin) registerModelObjectDeleteEndpoint(modelType reflect.Type) {
	dbModel := data.NewDbModel(modelType, admin.GormDB)
	modelObjectDeleteRoute := fmt.Sprintf("/admin/%s/actions/delete/:pk", modelType.Name())
	admin.Handler.RegisterDeleteEndpoint(modelObjectDeleteRoute, func(pk interface{}) error {
		return dbModel.DeleteObject(pk)
	})
}
