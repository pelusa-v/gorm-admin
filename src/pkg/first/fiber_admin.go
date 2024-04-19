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

type Admin struct {
	Handler AppHandler
}

type AppHandler interface {
	Register()
}

type FiberHandler struct {
	App *fiber.App
}

type BuiltInHandler struct {
}

type GinHandler struct {
	App *string
}

func (handler *FiberHandler) Register() {
	fmt.Println("Registering admin in Fiber app")

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
	fmt.Println("Registering admin in BuiltIn http app")
}

func (handler *GinHandler) Register() {
	fmt.Println("Registering admin in Gin app")
}

func NewFiberAdmin(app *fiber.App) *Admin {
	handler := &FiberHandler{App: app}
	admin := &Admin{Handler: handler}
	return admin
}

func NewGinAdmin(app *string) *Admin {
	handler := &GinHandler{App: app}
	admin := &Admin{Handler: handler}
	return admin
}

func NewAdmin() *Admin {
	handler := &BuiltInHandler{}
	admin := &Admin{Handler: handler}
	return admin
}

func (admin *Admin) Register() {
	admin.Handler.Register()
}
