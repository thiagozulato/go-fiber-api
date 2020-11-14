package main

import (
	"Modgo/context/book"
	"Modgo/domain/user/entities"
	"Modgo/utils"
	"database/sql"

	"github.com/gofiber/fiber/v2"
)

var (
	db *sql.DB
)

func main() {
	app := fiber.New(fiber.Config{
		ErrorHandler: func(c *fiber.Ctx, erro error) error {
			r, ok := erro.(utils.ResponseError)

			if !ok {
				c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
					"message": erro.Error(),
				})
				return nil
			}

			c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"message": r.Error(),
			})

			return nil
		},
	})

	app.Get("/", func(c *fiber.Ctx) error {
		return c.Status(fiber.StatusOK).JSON(entities.User{
			Name: "Go API Fiber Test",
		})
	})

	app.Get("/go/api/v1/demos", func(c *fiber.Ctx) error {
		return c.Status(fiber.StatusOK).SendString("Hello World")
	})

	api := app.Group("/api")
	v1 := api.Group("/v1")
	book.SetupBookApplication(db, v1)

	app.Listen(":3000")
}
