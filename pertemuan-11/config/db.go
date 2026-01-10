package config

import (
	"log"
	"os"
	"pertemuan-11/model"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB() {
    dsn := os.Getenv("POSTGRE_URI")
    if dsn == "" {
        log.Fatal("POSTGRE_URI tidak ditemukan di .env")
    }

    db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
    if err != nil {
        log.Fatal("Gagal terhubung ke database:", err)
    }

    // TAMBAHKAN BARIS INI:
    // Ini akan otomatis membuat tabel "users" jika belum ada
    err = db.AutoMigrate(&model.User{})
    if err != nil {
        log.Fatal("Gagal migrasi database:", err)
    }

    log.Println("Berhasil terhubung ke Supabase & Migrasi Tabel Selesai")
    DB = db
}