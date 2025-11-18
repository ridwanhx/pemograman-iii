package config

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// inisiasi variabel global DB (nantinya akan kita panggil di package lain)
var DB *gorm.DB;

// Inisialisasi Database
func InitDB() {
	// (1) Load environment dari file .env
	_ = godotenv.Load();

	// (2) Inisialisasi dsn
	dsn := os.Getenv("SUPABASE_DSN");

	// (3) Kondisi jika dsn tidak ditemukan
	if (dsn == "") {
		log.Fatalf("SUPABASE_DSN tidak ditemukan. Pastikan SUPABASE_DSN ada di file .env");
	}

	// (4) Membuka koneksi ke DB PostgreSQL menggunakan library gorm
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{});

	// (5) kondisi jika koneksi ke database gagal
	if (err != nil) {
		log.Fatalf("Gagal koneksi ke database: %v", err);
	}

	// (6) Inisiasi db sebagai nilai baru untuk DB (var global)
	DB = db;

	// (7) Pesan berhasil konek ke database
	fmt.Println("✅ Koneksi ke PostgreSQL (Supabase) berhasil.");
}

// Fungsi untuk mengambil instance koneksi database (gorm.DB) yang sudah diinisialisasi
func GetDB() *gorm.DB {
	// Hentikan program jika DB belum diinisialisasi
	if (DB == nil) {
		log.Fatal("DB belum diinisialisasi. Panggil config.InitDB() terlebih dahulu");
	}
	return DB;
}