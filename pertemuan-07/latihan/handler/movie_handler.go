package handler

import (
	"latihan/model"
	"latihan/repository"

	"github.com/gofiber/fiber/v2"
)

func GetAllMovie(c *fiber.Ctx) error {
	// inisiasi data, error hasil dari operasi di package repository
	data, err := repository.GetAllMovie();

	// jika error
	if (err != nil) {
		// kembalikan status code 500, beserta data Map (key: value) berikut
		return c.Status(500).JSON(fiber.Map{
			"message": "Gagal mengambil data",
			"error": err.Error(),
		});
	}

	// kembalikan JSON, beserta data Map (key: value) berikut
	return c.JSON(fiber.Map{
		"message": "Berhasil mengambil data movies",
		"data": data,
	});
}

// Insert Movie
func InsertMovie(c *fiber.Ctx) error {
	// inisiasi object yang merupakan instance ke model Movies
	var movie model.Movies;

	// buat kondisi untuk mengecek kesesuaian format dari data yang dikirim dengan data yang ada pada model Movies
	// dilakukan parsing (proses membaca dan memecah data mentah menjadi struktur yang bisa dipahami oleh program) terhadap data yang dikirimkan
	// lalu di cek apakah menghasilkan error
	if err := c.BodyParser(&movie); err != nil {
		// jika ya, kembalikan status code 400 beserta Map yang berisi pesan kesalahannya
		c.Status(400).JSON(fiber.Map{
			"message": "Format data salah.",
		})
	}

	// buat kondisi apakah proses insert data ke database menghasilkan error
	if err := repository.InsertMovie(movie); err != nil {
		// jika ya, maka kembalikan status code 500 (Internal Server Error)
		c.Status(500).JSON(fiber.Map{
			// sertakan pesan umum bahwa data baru gagal ditambahkan
			"message": "Gagal menambahkan data baru.",
			// sertakan detail error yang lebih informatif
			"error": err.Error(),
		})
	}

	// kembalikan status code 201 untuk proses insert data yang berhasil
	return c.Status(201).JSON(fiber.Map{
		// sertakan pesan umum bahwa data baru berhasil ditambahkan
		"message": "Data movie baru berhasil ditambahkan.",
		// sertakan detail data yang berhasil ditambahkan ke database
		"data": movie,
	})
}
