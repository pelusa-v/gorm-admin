package admin

import (
	"fmt"
	"io/fs"
	"reflect"

	"github.com/pelusa-v/gorm-admin/src/pkg/data"
)

func (admin *Admin) registerHomePage() {
	templates := admin.template("templates/*.html")
	templateManager := data.NewTemplateManager(&admin.Name, &admin.Models)

	admin.Handler.RegisterSimplePage(templates, "home.html", "/admin", func() any {
		return templateManager.GetHomePageData()
	})

	subFS, _ := fs.Sub(admin.TemplatesFs, "static/styles")
	admin.Handler.RegisterStatic(subFS)
}

func (admin *Admin) registerModelDetailPage(modelType reflect.Type) {
	dbModel := data.NewDbModel(modelType, admin.GormDB)
	templates := admin.template("templates/*.html")
	templateManager := data.NewTemplateManager(&admin.Name, &admin.Models)
	modelDetailPageRoute := fmt.Sprintf("/admin/%s", modelType.Name())

	admin.Handler.RegisterSimplePage(templates, "ModelDetail.html", modelDetailPageRoute, func() any {
		return templateManager.GetModelDetailPageData(*dbModel)
	})
}

func (admin *Admin) registerModelObjectDetailPage(modelType reflect.Type) {
	dbModel := data.NewDbModel(modelType, admin.GormDB)
	templates := admin.template("templates/*.html")
	templateManager := data.NewTemplateManager(&admin.Name, &admin.Models)
	modelObjectDetailPageRoute := fmt.Sprintf("/admin/%s/:pk", modelType.Name())

	admin.Handler.RegisterPkPage(templates, "ModelObjectDetail.html", modelObjectDetailPageRoute, func(pk string) any {
		return templateManager.GetModelObjectDetailPageData(*dbModel, pk)
	})
}

func (admin *Admin) registerModelObjectCreatePage(modelType reflect.Type) {
	dbModel := data.NewDbModel(modelType, admin.GormDB)
	templates := admin.template("templates/*.html")
	templateManager := data.NewTemplateManager(&admin.Name, &admin.Models)
	modelObjectCreatePageRoute := fmt.Sprintf("/admin/%s/actions/create", modelType.Name())

	admin.Handler.RegisterSimplePage(templates, "ModelObjectCreate.html", modelObjectCreatePageRoute, func() any {
		return templateManager.GetModelObjectCreatePageData(*dbModel)
	})
}

func (admin *Admin) registerModelObjectCreateEndpoint(modelType reflect.Type) {
	dbModel := data.NewDbModel(modelType, admin.GormDB)
	modelObjectCreateRoute := fmt.Sprintf("/admin/%s/actions/create", modelType.Name())
	modelDetailPageRoute := fmt.Sprintf("/admin/%s", modelType.Name())
	admin.Handler.RegisterCreateEndpoint(modelObjectCreateRoute, modelDetailPageRoute, func(data interface{}) error {
		return dbModel.CreateObject(data)
	})
}
