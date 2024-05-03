package handlers

import (
	"bytes"
	"fmt"
	"html/template"
	"io/fs"
	"net/http"
	"reflect"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/filesystem"
	"github.com/pelusa-v/gorm-admin/src/pkg/data"
)

type FiberHandler struct {
	BaseHandler
	App *fiber.App
}

// Handle /admin/Product/1 route, (:id)
func (handler *FiberHandler) RegisterSimplePage(tmpl *template.Template, templateName string, route string, tmplDataFunc func() any) {
	handler.App.Get(route, func(c *fiber.Ctx) error {
		var tmplOutput bytes.Buffer
		// err := tmpl.Execute(&tmplOutput, tmplDataFunc())
		err := tmpl.ExecuteTemplate(&tmplOutput, templateName, tmplDataFunc())
		if err != nil {
			panic(err)
		}

		c.Set("Content-Type", "text/html")
		return c.SendString(tmplOutput.String())
	})
}

func (handler *FiberHandler) RegisterPkPage(tmpl *template.Template, templateName string, route string, tmplDataFunc func(pk string) any) {
	handler.App.Get(route, func(c *fiber.Ctx) error {
		pk := c.Params("pk")

		var tmplOutput bytes.Buffer
		err := tmpl.ExecuteTemplate(&tmplOutput, templateName, tmplDataFunc(pk))
		if err != nil {
			panic(err)
		}

		c.Set("Content-Type", "text/html")
		return c.SendString(tmplOutput.String())
	})
}

func (handler *FiberHandler) RegisterCreateEndpoint(route string, typeToCreate reflect.Type, actionCreateFunc func(data interface{}) error) {

	RegisterFiberEndpoint(route, POST, handler, func(c *fiber.Ctx) error {
		// dataToCreate := reflect.New(typeToCreate).Elem()
		// fmt.Println(dataToCreate)
		// fmt.Println(string(c.Body()))
		// if err := c.BodyParser(dataToCreate.Addr().Interface()); err != nil {
		// 	panic(err)
		// }

		dataToCreate, err := data.GetObjectInstanceFromBytes(c.Body(), typeToCreate)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": "Cannot parse JSON",
			})
		}
		fmt.Println(dataToCreate)

		err = actionCreateFunc(dataToCreate)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": err.Error(),
			})
		}

		fmt.Println("Created...")
		return c.SendStatus(201)
	})
}

func (handler *FiberHandler) RegisterStatic(fs fs.FS) {
	handler.App.Use("/gorm-admin-statics", filesystem.New(filesystem.Config{
		Root: http.FS(fs),
	}))
}

func RegisterFiberEndpoint(route string, method RequestMethod, appHandler *FiberHandler, controller fiber.Handler) {
	switch method.Name {
	case GET.Name:
		appHandler.App.Get(route, controller)
	case POST.Name:
		appHandler.App.Post(route, controller)
	case PUT.Name:
		appHandler.App.Put(route, controller)
	case DELETE.Name:
		appHandler.App.Delete(route, controller)
	}
}
