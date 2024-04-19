package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/pelusa-v/gorm-admin/src/pkg/first"
)

func main() {
	app := fiber.New()

	handler := &first.FiberHandler{App: app}
	admin := first.SimpleAdmin{Handler: handler}

	admin.Handler.Register()

	// db := samples.NewDbInstance()
	// samples.TestListHandler(db)

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	app.Listen(":3000")
}
