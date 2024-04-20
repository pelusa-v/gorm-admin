package handlers

import (
	"bytes"
	"embed"
	"fmt"
	"text/template"

	"github.com/gofiber/fiber/v2"
)

type FiberHandler struct {
	App         *fiber.App
	TemplatesFs embed.FS
}

func (handler *FiberHandler) Register() {
	fmt.Println("Registering admin in Fiber app")

	tmpl, err := template.ParseFS(handler.TemplatesFs, "templates/home.html")
	if err != nil {
		panic(err)
	}

	handler.App.Get("/admin", func(c *fiber.Ctx) error {

		// Send the HTML content as the response
		var tmplOutput bytes.Buffer
		err = tmpl.Execute(&tmplOutput, nil)
		if err != nil {
			return err
		}

		// Send the rendered HTML as the response
		c.Set("Content-Type", "text/html")
		return c.SendString(tmplOutput.String())
	})
}
