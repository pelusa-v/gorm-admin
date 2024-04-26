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

// Handle /admin/Product/1 route, (:id)
func (handler *FiberHandler) RegisterSimplePage(tmpl *template.Template, route string, tmplDataFunc func() any) {
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

func (handler *FiberHandler) RegisterPkPage(tmpl *template.Template, route string, tmplDataFunc func(pk string) any) {
	handler.App.Get(route, func(c *fiber.Ctx) error {
		pk := c.Params("pk")

		var tmplOutput bytes.Buffer
		err := tmpl.Execute(&tmplOutput, tmplDataFunc(pk))
		if err != nil {
			panic(err)
		}

		c.Set("Content-Type", "text/html")
		return c.SendString(tmplOutput.String())
	})
}

func (handler *FiberHandler) RegisterStatic(staticFolder string, staticPath string) {
	handler.App.Static(staticPath, staticFolder, fiber.Static{
		Browse: true,
	})
}
