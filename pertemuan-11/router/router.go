package router

import (
    "pertemuan-11/config"
    "pertemuan-11/handler"

    "github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {
    // Grouping untuk versi API
    api := app.Group("/api")

    // =========================
    // PUBLIC ROUTES
    // =========================
    api.Post("/login", handler.Login)
    api.Post("/register", handler.CreateUser)

    // =========================
    // PROTECTED ROUTES (butuh JWT)
    // =========================
    protected := api.Group("", config.JWTMiddleware())

    // ---- USER ROUTES ----
    protected.Get("/users", handler.GetUsers)
    protected.Post("/users", handler.CreateUser)
    protected.Get("/users/:username", handler.GetUserByUsername)

    // ---- MAHASISWA ROUTES ----
    protected.Get("/mahasiswa", handler.GetAllMahasiswa)           // ambil semua data
    protected.Post("/mahasiswa", handler.CreateMahasiswa)          // tambah data baru
    protected.Get("/mahasiswa/:npm", handler.GetMahasiswaByNpm)    // ambil data berdasarkan NPM
    protected.Put("/mahasiswa/:npm", handler.UpdateMahasiswaByNpm) // update data berdasarkan NPM
    protected.Delete("/mahasiswa/:npm", handler.DeleteMahasiswaByNpm) // hapus data berdasarkan NPM
}
