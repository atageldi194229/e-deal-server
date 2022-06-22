package routes

import (
	"github.com/atageldi194229/e-deal-server/controllers"
	"github.com/gofiber/fiber/v2"
)

func welcome(c *fiber.Ctx) error {
	return c.SendString("Welcome to my awesome app")
}

func Setup(app *fiber.App) {
	api := app.Group("api")

	// Welcome
	api.Get("/", welcome)

	// Category
	categories := api.Group("categories")
	categories.Post("/", controllers.CreateCategory)
	categories.Get("/", controllers.GetCategories)
	categories.Get("/:id", controllers.GetCategory)
	categories.Put("/:id", controllers.UpdateCategory)
	categories.Delete("/:id", controllers.DeleteCategory)
}
