package route

import (
	"github.com/gofiber/fiber/v2"
	"github.com/jesusGZ/simple_server_of_fiber/pkg/handler"
	"gorm.io/gorm"
)

// Expenses route
func Expenses(app *fiber.App, db *gorm.DB) {
	h := &handler.ExpenseHandler{
		DB: db,
	}
	r := app.Group("/expenses")
	r.Get("/", h.Index)
	r.Get("/:id", h.Show)
	r.Post("/", h.Store)
	r.Put("/:id", h.Update)
	r.Delete("/:id", h.Destroy)
	r.Post("/:id/upload", h.Upload)
}
