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

// Handle /admin/Product?id=1 route, (DON'T REGISTER ALL ID's ROUTES)
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
