package server

import (
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

func Init() {
	app := fiber.New()
	app.Use(recover.New())
	// app.Use(logger.New())

	// app.Get("/list/:id", func(c *fiber.Ctx) error {
	// 	return c.SendString(c.Params("id"))
	// })
	controller.setUserControllers(app)

	log.Fatal(app.Listen(":" + os.Getenv("PORT")))
}
