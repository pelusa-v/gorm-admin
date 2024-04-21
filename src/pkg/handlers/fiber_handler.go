package handlers

import (
	"bytes"
	"html/template"

	"github.com/gofiber/fiber/v2"
)

type FiberHandler struct {
	BaseHandler
	App *fiber.App
}

func (handler *FiberHandler) RegisterPage(tmpl *template.Template, route string, tmplDataFunc func() any) {
	handler.App.Get(route, func(c *fiber.Ctx) error {
		var tmplOutput bytes.Buffer
		err := tmpl.Execute(&tmplOutput, tmplDataFunc())
		if err != nil {
			panic(err)
		}

		c.Set("Content-Type", "text/html")
		return c.SendString(tmplOutput.String())
	})
}

// func (handler *FiberHandler) RegisterHomePage(tmpl *template.Template) {
// 	handler.App.Get("/admin", func(c *fiber.Ctx) error {
// 		tmplOutput := data.GetHomePage(handler.Models, tmpl)
// 		c.Set("Content-Type", "text/html")
// 		return c.SendString(tmplOutput.String())
// 	})
// }

// func (handler *FiberHandler) RegisterModelDetailPage(modelType reflect.Type, tmpl *template.Template) {
// 	model := data.NewDbModel(modelType, handler.GormDB)

// 	handler.App.Get(fmt.Sprintf("/admin/%s", modelType.Name()), func(c *fiber.Ctx) error {
// 		tmplOutput := data.GetModelDetailPage(*model, tmpl)
// 		c.Set("Content-Type", "text/html")
// 		return c.SendString(tmplOutput.String())
// 	})
// }
