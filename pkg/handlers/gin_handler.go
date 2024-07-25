package handlers

import (
	"bytes"
	"html/template"
	"io/fs"
	"net/http"

	"github.com/gin-gonic/gin"
)

type GinHandler struct {
	BaseHandler
	App *gin.Engine
}

//	func (handler *GinHandler) Register() {
//		fmt.Println("Registering admin in Gin app")
//	}
func (handler *GinHandler) RegisterSimplePage(tmpl *template.Template, templateName string, route string, tmplDataFunc func() any) {
	handler.App.GET(route, func(c *gin.Context) {
		var tmplOutput bytes.Buffer
		err := tmpl.ExecuteTemplate(&tmplOutput, templateName, tmplDataFunc())
		if err != nil {
			panic(err)
		}
		c.Set("Content-Type", "text/html")
		c.Data(http.StatusOK, "text/html; charset=utf-8", []byte(tmplOutput.String()))
	})
}

func (handler *GinHandler) RegisterPkPage(tmpl *template.Template, templateName string, route string, tmplDataFunc func(pk string) any) {
}

func (handler *GinHandler) RegisterStatic(fs fs.FS) {
	handler.App.StaticFS("/gorm-admin-statics", http.FS(fs))
}

// func (handler *FiberHandler) RegisterStatic2(fs fs.FS) {
// 	handler.App.Use("/gorm-admin-statics", filesystem.New(filesystem.Config{
// 		Root: http.FS(fs),
// 	}))
// }

func (handler *GinHandler) RegisterCreateEndpoint(route string, actionCreateFunc func(data interface{}) error) {
}

func (handler *GinHandler) RegisterDeleteEndpoint(route string, actionFunc func(pk interface{}) error) {
}
