package book

import (
	"Modgo/context/book/application"
	"Modgo/context/book/domain/commands"
	infrasctructure "Modgo/context/book/infrastructure"
	"database/sql"

	"github.com/gofiber/fiber/v2"
)

// SetupBookApplication ...
func SetupBookApplication(db *sql.DB, router fiber.Router) {
	g := router.Group("/books")

	repository := infrasctructure.NewBookRepository(db)
	commands := commands.NewBookCommandHandler(repository)
	controller := application.NewBookController(commands)

	g.Get("/", controller.GetAll)
	g.Get("/:id", controller.GetByID)
	g.Post("/", controller.Create)
	g.Put("/:id", controller.Update)
	g.Delete("/:id", controller.Delete)
}
