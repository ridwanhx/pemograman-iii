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

	// Menambahkan endpoint GET /api/mahasiswa untuk mengambil seluruh data mahasiswa
	api.Get("/mahasiswa", handler.GetAllMahasiswa);
	api.Get("/mahasiswa/:npm", handler.GetMahasiswaByNpm);
	
	api.Post("/mahasiswa", handler.InsertMahasiswa);

	api.Put("/mahasiswa/:npm", handler.UpdateMahasiswaByNpm)

	api.Patch("/mahasiswa/:npm", handler.UpdateMahasiswaByNpm)

	// api.Delete("/mahasiswa/:id", handler.DeleteMahasiswaById)
	api.Delete("/mahasiswa/:npm", handler.DeleteMahasiswaByNpm)
}