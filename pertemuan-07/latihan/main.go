package main

import (
	"latihan/config"
	"latihan/model"
	"latihan/router"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func main() {
	// (1) Inisialisasi objek baru dari Fiber sebagai instance utama aplikasi
	app := fiber.New();

	// (2) inisialisasi koneksi ke DB
	config.InitDB();

	// (3) get data movies
	config.DB.AutoMigrate(&model.Movies{});

	// (4) Logging request
	app.Use(logger.New());

	// (5) Panggil SetupRoutes yang sebelumnya kita inisialisasi di package router
	router.SetupRoutes(app);

	// (6) jalankan server
	app.Listen(":3000");
}