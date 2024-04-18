package main

import "github.com/pelusa-v/gorm-admin/src/samples"

func main() {
	// app := fiber.New()

	// app.Get("/", func(c *fiber.Ctx) error {
	// 	return c.SendString("Hello, World!")
	// })

	// app.Listen(":3000")
	db := samples.NewDbInstance()
	samples.TestListHandler(db)
}
