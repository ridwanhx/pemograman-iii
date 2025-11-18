package repository

import (
	"latihan/config"
	"latihan/model"
)

// Fungsi untuk mengambil seluruh data movies dari database
// Mengembalikan slice berisi data movies dan error jika terjadi kesalahan
func GetAllMovie() ([]model.Movies, error) {
	// Inisialisasi slice untuk menampung data movies
	var data []model.Movies;
	result := config.GetDB().Find(&data);
	// - config.GetDB() → mengambil instance koneksi database yang sudah diinisialisasi sebelumnya.
	// - .Find(&data) → menjalankan query untuk mengambil semua record dari tabel yang sesuai dengan struct data
	return data, result.Error;
}

// InsertMovie menyimpan data movie baru ke database
// Mengembalikan error jika proses insert gagal
func InsertMovie(movie model.Movies) error {
	// Eksekusi perintah insert menggunakan GORM
	result := config.GetDB().Create(&movie);
	return result.Error;
}