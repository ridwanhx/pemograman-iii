package handler

import (
	"latihan/model"
	"latihan/repository"

	"github.com/gofiber/fiber/v2"
)

func GetAllMahasiswa(c *fiber.Ctx) error {
	data, err := repository.GetAllMahasiswa()

	if (err != nil) {
		return c.Status(500).JSON(fiber.Map{
			"message": "Gagal Mengambil data",
			"error": err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"message": "Berhasil mengambil data mahasiswa",
		"data": data,
	})
}

func InsertMahasiswa(c *fiber.Ctx) error {
	var mahasiswa model.Mahasiswa

	if err := c.BodyParser(&mahasiswa); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"message": "Format data salah.",
			"error": err.Error(),
		})
	}

	if err := repository.InsertMahasiswa(mahasiswa); err != nil {
		return c.Status(500).JSON(fiber.Map{
			"message": "Gagal menambahkan data baru.",
			"error": err.Error(),
		})
	}

	return c.Status(201).JSON(fiber.Map{
		"message": "Data mahasiswa berhasil ditambahkan.",
		"data": mahasiswa,
	})
}

func GetMahasiswaByNpm(c *fiber.Ctx) error {
	npm := c.Params("npm")

	data, err := repository.GetMahasiswaByNpm(npm)

	if err != nil {
		return c.Status(404).JSON(fiber.Map{
			"message": "Mahasiswa dengan NPM tersebut tidak ditemukan.",
			"error": err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"message": "Berhasil mengambil data mahasiswa berdasarkan NPM",
		"data": data,
	})
}

func UpdateMahasiswaByNpm(c *fiber.Ctx) error {
	npm := c.Params("npm")
	var updateData map[string]interface{}

	if err := c.BodyParser(&updateData); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"message": "Format data salah",
			"error": err.Error(),
		})
	}

	if err := repository.UpdateMahasiswaByNpm(npm, updateData); err != nil {
		return c.Status(500).JSON(fiber.Map{
			"message": "Gagal update data mahasiswa",
			"error": err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"message": "Data mahasiswa berhasil diubah.",
		"data": updateData,
	})
}

func DeleteMahasiswaById(c *fiber.Ctx) error {
	id := c.Params("id")

	if err := repository.DeleteMahasiswaById(id); err != nil {
		return c.Status(500).JSON(fiber.Map{
			"message": "Gagal menghapus data mahasiswa",
			"error": err.Error(),
		})
	}
	return c.JSON(fiber.Map{
		"message": "Data mahasiswa berhasil dihapus",
	})
}