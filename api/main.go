package main

import (
	"Modgo/context/book"
	"Modgo/domain/user/entities"
	"Modgo/utils"
	"database/sql"
	"errors"

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
				c.Status(500).JSON(fiber.Map{
					"message": erro.Error(),
				})
				return nil
			}

			c.Status(400).JSON(fiber.Map{
				"message": r.Error(),
			})

			return nil
		},
	})

	app.Get("/:name", func(c *fiber.Ctx) error {
		n := c.Params("name")

		if n == "teste" {
			return utils.NewResponseError("Usuário Inválido")
		}

		if len(n) <= 6 {
			return errors.New("Tamanho do nome inválido")
		}

		u := entities.User{
			Name: n,
		}

		return c.Status(fiber.StatusOK).JSON(u)
	})

	api := app.Group("/api")
	v1 := api.Group("/v1")
	book.SetupBookApplication(db, v1)

	app.Listen(":3000")
}
