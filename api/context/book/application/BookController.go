package application

import (
	"Modgo/context/book/domain/commands"

	"github.com/gofiber/fiber/v2"
)

// BookController ...
type BookController struct {
	commands commands.BookCommandHandler
}

// NewBookController ...
func NewBookController(commands commands.BookCommandHandler) BookController {
	return BookController{
		commands: commands,
	}
}

// GetAll ...
func (c BookController) GetAll(ctx *fiber.Ctx) error {
	books, _ := c.commands.GetAllBooksHandle()

	if len(books) == 0 {
		ctx.Status(404).JSON(fiber.Map{
			"message": "No Books found",
		})
	}

	ctx.Status(200).JSON(books)
	return nil
}

// GetByID ...
func (c BookController) GetByID(ctx *fiber.Ctx) error {
	id := ctx.Params("id")
	book, _ := c.commands.GetByIDHandle(id)

	ctx.Status(200).JSON(book)
	return nil
}

// Create ...
func (c BookController) Create(ctx *fiber.Ctx) error {
	var book commands.CreateBookCommand

	if err := ctx.BodyParser(&book); err != nil {
		ctx.Status(400).JSON(fiber.Map{
			"message": "Invalid Request",
		})
		return nil
	}

	c.commands.CreateBookHandle(book)

	ctx.Status(fiber.StatusNoContent).JSON(nil)
	return nil
}

// Update ...
func (c BookController) Update(ctx *fiber.Ctx) error {
	id := ctx.Params("id")
	var book commands.UpdateBookCommand

	if err := ctx.BodyParser(&book); err != nil {
		ctx.Status(400).JSON(fiber.Map{
			"message": "Invalid Request",
		})
		return nil
	}

	c.commands.UpdateBookHandle(id, book)
	ctx.Status(fiber.StatusNoContent).JSON(nil)
	return nil
}

// Delete ...
func (c BookController) Delete(ctx *fiber.Ctx) error {
	id := ctx.Params("id")

	c.commands.DeleteBookHandle(id)
	ctx.Status(fiber.StatusNoContent).JSON(nil)
	return nil
}
