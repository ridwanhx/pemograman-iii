package router

import (
	"pertemuan-11/config"
	"pertemuan-11/handler"

	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {
	// Grouping untuk versi API
	api := app.Group("/api")

	// PUBLIC
	api.Post("/login", handler.Login)
	api.Post("/register", handler.CreateUser)

	// PROTECTED
	protected := api.Group("", config.JWTMiddleware())

	// Route untuk User (butuh token)
	protected.Get("/users", handler.GetUsers)
	protected.Post("/users", handler.CreateUser)
	protected.Get("/users/:username", handler.GetUserByUsername)
}