package main

import (
	"pertemuan-06/config"
	"pertemuan-06/model"
	"pertemuan-06/router"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func main() {
	// instansiasi fiber / inisialisasi penggunaan fiber
	app := fiber.New()

	// inisiasi koneksi ke DB
	// <nama_package>.<nama_function>
	config.InitDB()
	
	// Membuat auto migrate model
	// yang akan kita migrate adalah struct Mahasiswa, dan letak daripada struct tersebut ada di package model
	config.GetDB().AutoMigrate(&model.Mahasiswa{})
	
	// logging request
	// misal nantinya ada yang akses API (GET, POST, dll)
	// ketika ada permintaan dari client, nanti akan ada log nya
	app.Use(logger.New())

	// panggil setup route yang sudah kita buat sebelumnya di package router
	// jangan sampai terlewat, karena jika kita tidak menambahkan router ke main, nanti pada saat dilakukan pengetesan akan muncul error dengan status 404 Not found
	router.SetupRoutes(app)

	// inisialisasi port
	app.Listen(":3000")
}