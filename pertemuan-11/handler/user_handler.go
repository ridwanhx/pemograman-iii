package handler

import (
	"net/http"
	"pertemuan-11/model"
	"pertemuan-11/repository"

	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
)

func GetUsers(c *fiber.Ctx) error {
	users, err := repository.GetAllUser()

	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"message": "Gagal mengambil data user",
			"error":   err.Error(),
		})
	}

	return c.Status(http.StatusOK).JSON(fiber.Map{
		"message": "Berhasil mengambil data user",
		"data":    users,
	})
}

func CreateUser(c *fiber.Ctx) error {
	var request model.CreateUserRequest
	
	if err := c.BodyParser(&request); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"message": "Format data salah",
		})
	}

	if request.Username == "" || request.Password == "" {
		return c.Status(400).JSON(fiber.Map{
			"message": "Username & password wajib diisi.",
		})
	}

	// Cek username sudah ada atau belum
	if _, err := repository.FindByUsername(request.Username); err == nil {
		return c.Status(409).JSON(fiber.Map{
			"message": "Username sudah digunakan",
		})
	}

	// Hash password
	hash, err := bcrypt.GenerateFromPassword([]byte(request.Password), bcrypt.DefaultCost)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"message": "Gagal hash password",
			"error":   err.Error(),
		})
	}

	// Map ke Model User (PostgreSQL/GORM menggunakan ID uint/otomatis)
	newUser := model.User{
		Username: request.Username,
		Password: string(hash),
		Role:     "user",
	}

	user, err := repository.CreateUser(newUser)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"message": "Gagal membuat user baru",
			"error":   err.Error(),
		})
	}

	return c.Status(http.StatusCreated).JSON(fiber.Map{
		"message": "User berhasil dibuat",
		"data":    user,
	})
}

func GetUserByUsername(c *fiber.Ctx) error {
	username := c.Params("username")

	user, err := repository.FindByUsername(username)
	if err != nil {
		return c.Status(http.StatusNotFound).JSON(fiber.Map{
			"message": "User tidak ditemukan",
			"error":   err.Error(),
		})
	}

	return c.Status(http.StatusOK).JSON(fiber.Map{
		"message": "User ditemukan",
		"data":    user,
	})
}