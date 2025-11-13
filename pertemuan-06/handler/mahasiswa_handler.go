package handler

import (
	"pertemuan-06/model"
	"pertemuan-06/repository"
	"github.com/gofiber/fiber/v2"
)

// implementasi penggunaan fiber
func GetAllMahasiswa(c *fiber.Ctx) error {
	// inisiasi variables (data, err), merujuk pada apa yang kita lakukan di repository
	// dimana kita membuat dua return yaitu data (model Mahasiswa) dan error
	// maka dari itu wajib membuat sejumlah yang kita definisikan di repository
	data, err := repository.GetAllMahasiswa()

	if err != nil {
		// status response code 500 / Internal Server Error, keluarannya adalah JSON dan dimasukkan kedalam Map
		return c.Status(500).JSON(fiber.Map{
			"error":   err.Error(),
			"message": "Gagal mengambil data mahasiswa",
		})
	}

	// jika berhasil, buat fiber Map yang mengembalikan data dan message nya.
	return c.JSON(fiber.Map{
		"data":    data,
		"message": "Berhasil mengambil data mahasiswa",
	})
}

func InsertMahasiswa(c *fiber.Ctx) error {
	var mhs model.Mahasiswa

	if err := c.BodyParser(&mhs); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": err.Error(),
			"message": "Format data salah",
		})
	}

	if err := repository.InsertMahasiswa(mhs); err != nil {
		return c.Status(500).JSON(fiber.Map{
			"error": err.Error(),
			"message": "Gagal memasukkan data mahasiswa",
		})
	}

	return c.Status(201).JSON(fiber.Map{
		"message": "Berhasil memasukkan data mahasiswa",
		"data": mhs,
	})
}