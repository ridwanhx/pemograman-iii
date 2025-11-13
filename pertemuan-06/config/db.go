package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// deklarasi variabel global
// variabel ini nantinya bisa kita pakai / panggil di repository
var DB *gorm.DB

// inisiasi satu function yang nantinya bisa kita gunakan di beberapa tempat (reusable)
// fungsi ini akan terus dipanggil dengan memanggil fungsi GetDB dibawah
func InitDB() {
	// melakukan load / membaca file .env
	_ = godotenv.Load()

	// deklarasi variabel untuk menampung nilai dari SUPABASE_DSN
	dsn := os.Getenv("SUPABASE_DSN")

	// lakukan validasi
	// jika dsn tidak ditemukan
	if dsn == "" {
		log.Fatal("SUPABASE_DSN tidak ditemukan. Pastikan .env berisi DNS Supabase.")
	}

	// melakukan open koneksi ke dsn
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	
	// jika gagal koneksi (koneksi mati, dan lain sebagainya)
	if err != nil {
		log.Fatalf("Gagal konek ke database: %v", err)
	}

	// jika koneksi ke DB berhasil
	DB = db
	log.Println("✅ Koneksi ke PostgreSQL (Supabase) Berhasil.")
}

// akan / bisa kita gunakan di tempat lain / di function lain
func GetDB() *gorm.DB {
	if DB == nil {
		log.Fatal("DB belum diinisialisasi. Panggil config.InitDB terlebih dahulu.")
	}
	return DB
}