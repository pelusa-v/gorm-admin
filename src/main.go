package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/pelusa-v/gorm-admin/src/pkg/admin"
	"github.com/pelusa-v/gorm-admin/src/samples"
)

func main() {
	app := fiber.New()
	db := samples.NewDbInstance()

	admin := admin.NewFiberAdmin(app, db)
	admin.Register()
	admin.RegisterModel(samples.User{})
	admin.RegisterModel(samples.Product{})

	// db := samples.NewDbInstance()
	// samples.TestListHandler(db)

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	app.Listen(":3000")
}
