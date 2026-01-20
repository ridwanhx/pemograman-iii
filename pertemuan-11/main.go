package main

import (
    "log"
    "pertemuan-11/config"
    "pertemuan-11/router"

    "github.com/gofiber/fiber/v2"
    "github.com/joho/godotenv"
)

func main() {
    // 1. Load .env
    if err := godotenv.Load(); err != nil {
        log.Fatal("Error Loading .env file")
    }

    // 2. Koneksi ke DB
    config.ConnectDB()

    // 3. Inisialisasi Fiber
    app := fiber.New()

    // 4. Aktifkan middleware CORS
    app.Use(config.CorsConfig())

    // 5. Setup Router
    router.SetupRoutes(app)

    // 6. Jalankan server
    log.Fatal(app.Listen(":3000"))
}
