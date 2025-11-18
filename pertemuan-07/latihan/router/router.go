package router

import (
	"latihan/handler"

	"github.com/gofiber/fiber/v2"
)

// *fiber.App adalah pointer ke instance utama aplikasi Fiber.
// Parameter app di sini adalah instance yang diinisialisasi di main.go
// Fungsi SetupRoutes digunakan untuk mendefinisikan semua route aplikasi
func SetupRoutes(app *fiber.App) {
	// Membuat grup route dengan prefix (awalan) /api
	api := app.Group("/api");

	// Menambahkan endpoint GET /api/movies untuk mengambil seluruh data movies
	api.Get("/movies", handler.GetAllMovie);
	api.Post("/movies", handler.InsertMovie);
}