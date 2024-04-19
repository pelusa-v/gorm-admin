package first

import (
	"bytes"
	"embed"
	"fmt"
	"html/template"

	"github.com/gofiber/fiber/v2"
)

//go:embed templates/*
//go:embed home2.html
var homeHTML embed.FS

type FiberAdmin struct {
	Framework string
}

func GenerateAdmin() *FiberAdmin {
	return new(FiberAdmin)
}

type SimpleAdmin struct {
	Handler AppHandler
}

type AppHandler interface {
	Register()
}

type FiberHandler struct {
	App *fiber.App
}

type BuiltInHandler struct {
	App *int
}

type GinHandler struct {
	App *string
}

func (handler *FiberHandler) Register() {
	tmpl, err := template.ParseFS(homeHTML, "templates/home.html")
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

func (handler *BuiltInHandler) Register() {
	fmt.Println("Registering BuiltIn http app to admin")
}

func (handler *GinHandler) Register() {
	fmt.Println("Registering Gin app to admin")
}
