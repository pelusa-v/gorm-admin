package main

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/pelusa-v/gorm-admin/src/pkg/admin"
	"github.com/pelusa-v/gorm-admin/src/samples"
)

func main() {
	app := fiber.New()
	db := samples.NewDbInstance()

	admin := admin.NewFiberAdmin(app, db)
	fmt.Println("Name")
	fmt.Printf("Here is:%s//////", admin.Name)
	admin.Register()
	admin.Configure("My awesome project")
	fmt.Println("Name")
	fmt.Printf("Here is:%s//////", admin.Name)
	admin.RegisterModel(samples.User{})
	admin.RegisterModel(samples.Product{})
	admin.RegisterModel(samples.Car{})

	// db := samples.NewDbInstance()
	// samples.TestListHandler(db)

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	app.Listen(":3000")
}
