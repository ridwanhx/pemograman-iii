package router

import (
	"pertemuan-06/handler"

	"github.com/gofiber/fiber/v2"
)

// inisiasi setup route
func SetupRoutes(app *fiber.App) {
	// ketika nanti ingin mengakses /api
	api := app.Group("/api")

	// inisialisasi endpoint berdasarkan request method nya
	api.Get("/mahasiswa", handler.GetAllMahasiswa)
	api.Post("/mahasiswa", handler.InsertMahasiswa)
}