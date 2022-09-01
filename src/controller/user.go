package controller

import (
	"github.com/gofiber/fiber"
	"go-fiber.com/src/db"
	"go-fiber.com/src/model"
)

type UserController struct{}

func setUserControllers(app *fiber.App) {
	var uc UserController
	u := app.Group("/user")
	u.Get("/", uc.List)

}

func (uc *UserController) List(c *fiber.Ctx) error {
	db := db.DBinstance
	var data []model.User
	db.Find(&data)
	return c.JSON(fiber.Map{"status": "success", "message": "done", "result": data})
}
