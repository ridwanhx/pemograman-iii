package handler

import (
	"fmt" // Tambahkan fmt untuk konversi ID ke string
	"os"
	"strconv"

	"pertemuan-11/config"
	"pertemuan-11/model"
	"pertemuan-11/repository"

	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
)

func Login(c *fiber.Ctx) error {
	var req model.LoginRequest

	if err := c.BodyParser(&req); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"message": "Format data salah",
		})
	}

	// 1. Cari user di PostgreSQL (Supabase)
	user, err := repository.FindByUsername(req.Username)
	if err != nil {
		return c.Status(401).JSON(fiber.Map{
			"message": "Username atau password salah",
		})
	}

	// 2. Komparasi password
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password)); err != nil {
		return c.Status(401).JSON(fiber.Map{
			"message": "Username atau password salah",
		})
	}

	// 3. Konfigurasi expiry token
	expStr := os.Getenv("JWT_EXPIRES_MINUTES")
	expMin, _ := strconv.Atoi(expStr)
	if expMin == 0 {
		expMin = 60 
	}

	// 4. Generate Token 
	// PENYESUAIAN: Karena user.ID adalah uint (Postgres), gunakan fmt.Sprint untuk menjadikannya string
	userIDStr := fmt.Sprint(user.ID) 
	token, err := config.GenerateToken(userIDStr, user.Username, user.Role, expMin)

	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"message": "Gagal membuat token",
			"error":   err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"message": "Login berhasil",
		"token":   token,
		"user": fiber.Map{
			"username": user.Username,
			"role":     user.Role,
		},
	})
}