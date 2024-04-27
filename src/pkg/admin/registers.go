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
