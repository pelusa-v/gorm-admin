package admin

import (
	"fmt"
	"reflect"

	"github.com/pelusa-v/gorm-admin/src/pkg/data"
)

func (admin *Admin) registerHomePage() {
	homePageTemplate := admin.template("templates/home.html")

	admin.Handler.RegisterPage(homePageTemplate, "/admin", func() any {
		return data.GetHomePageData(&admin.Models)
	})
}

func (admin *Admin) registerModelDetailPage(modelType reflect.Type) {
	dbModel := data.NewDbModel(modelType, admin.GormDB)
	modelDetailPageTemplate := admin.template("templates/ModelDetail.html")
	modelDetailPageRoute := fmt.Sprintf("/admin/%s", modelType.Name())

	admin.Handler.RegisterPage(modelDetailPageTemplate, modelDetailPageRoute, func() any {
		return data.GetModelDetailPageData(*dbModel)
	})
}
