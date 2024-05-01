package handlers

import (
	"bytes"
	"html/template"
	"io/fs"
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/filesystem"
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
		// err := tmpl.Execute(&tmplOutput, tmplDataFunc(pk))
		err := tmpl.ExecuteTemplate(&tmplOutput, templateName, tmplDataFunc(pk))
		if err != nil {
			panic(err)
		}

		c.Set("Content-Type", "text/html")
		return c.SendString(tmplOutput.String())
	})
}

func (handler *FiberHandler) RegisterCreateEndpoint(route string, redirect string, actionCreateFunc func(data interface{}) error) {

	RegisterFiberEndpoint(route, POST, handler, func(c *fiber.Ctx) error {
		var dataToCreate interface{}
		err := actionCreateFunc(c.BodyParser(dataToCreate))
		if err != nil {
			panic(err)
			return c.SendStatus(500)
		}
		// return c.Redirect(redirect)
		return c.SendStatus(200)
	})
}

// func (handler *FiberHandler) RegisterCreateEndpoint2(route string, redirect string, typeToCreate reflect.Type, fieldsToExtract []string,
// 	actionCreateFunc func(reflect.Type, map[string]interface{}) any) {

// 	handler.App.Get(route, func(c *fiber.Ctx) error {
// 		formfields := make(map[string]interface{}, len(fieldsToExtract))
// 		for _, f := range fieldsToExtract {
// 			formfields[f] = c.FormValue(f)
// 		}

// 		actionCreateFunc(typeToCreate, formfields)
// 		return c.Redirect(redirect)
// 	})
// }

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
