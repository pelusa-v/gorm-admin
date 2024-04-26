package admin

import (
	"fmt"
	"io/fs"
	"reflect"

	"github.com/pelusa-v/gorm-admin/src/pkg/data"
)

func (admin *Admin) registerHomePage() {
	homePageTemplate := admin.template("templates/home.html")

	admin.Handler.RegisterSimplePage(homePageTemplate, "/admin", func() any {
		return data.GetHomePageData(&admin.Models)
	})

	subFS, _ := fs.Sub(admin.TemplatesFs, "static/styles")
	admin.Handler.RegisterStatic(subFS)
}

func (admin *Admin) registerModelDetailPage(modelType reflect.Type) {
	dbModel := data.NewDbModel(modelType, admin.GormDB)
	modelDetailPageTemplate := admin.template("templates/ModelDetail.html")
	modelDetailPageRoute := fmt.Sprintf("/admin/%s", modelType.Name())

	admin.Handler.RegisterSimplePage(modelDetailPageTemplate, modelDetailPageRoute, func() any {
		return data.GetModelDetailPageData(*dbModel)
	})
}

func (admin *Admin) registerModelObjectDetailPage(modelType reflect.Type) {
	dbModel := data.NewDbModel(modelType, admin.GormDB)
	modelObjectDetailPageTemplate := admin.template("templates/ModelObjectDetail.html")
	modelObjectDetailPageRoute := fmt.Sprintf("/admin/%s/:pk", modelType.Name())

	admin.Handler.RegisterPkPage(modelObjectDetailPageTemplate, modelObjectDetailPageRoute, func(pk string) any {
		return data.GetModelObjectDetailPageData(*dbModel, pk)
	})
}
