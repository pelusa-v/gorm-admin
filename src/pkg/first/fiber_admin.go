package first

import (
	"bytes"
	"embed"
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

func (admin *FiberAdmin) RegisterApp(app *fiber.App) {
	tmpl, err := template.ParseFS(homeHTML, "templates/home.html")
	if err != nil {
		panic(err)
	}

	app.Get("/admin", func(c *fiber.Ctx) error {

		// Send the HTML content as the response
		var tmplOutput bytes.Buffer
		err = tmpl.Execute(&tmplOutput, nil)
		if err != nil {
			return err
		}

		// Send the rendered HTML as the response
		c.Set("Content-Type", "text/html")
		return c.SendString(tmplOutput.String())

		// tmpl := template.Must(template.ParseGlob("pkg/**/*.html"))

		// var tmplOutput bytes.Buffer
		// err := tmpl.ExecuteTemplate(&tmplOutput, "home.html", nil)
		// if err != nil {
		// 	return err
		// }

		// return c.SendString(tmplOutput.String())
	})
}
