package repository

import (
	"pertemuan-06/config"
	"pertemuan-06/model"
)

// inisiasi fungsi untuk mengambil semua data mahasiswa
// ([]model.Mahasiswa, error) => merujuk pada return / nilai kembalian nantinya yang dihasilkan oleh fungsi ini.
// apakah itu akan mengembalikan model Mahasiswa (berhasil), ataupun error (gagal).
// []<nama_package>.<Nama_Struct>
func GetAllMahasiswa() ([]model.Mahasiswa, error) {
	// tampung data-data mahasiswa
	var data []model.Mahasiswa

	// panggil GetDB untuk permintaan koneksi ke Postgresql
	// ORM : Object Relational Mapping
	// kalau di golang, disebutnya gorm
	// seperti pada contoh dibawah, GetDB().Find() merujuk pada konsep ORM / gorm
	result := config.GetDB().Find(&data)
	return data, result.Error
}

func InsertMahasiswa(mhs model.Mahasiswa) error {
	result := config.GetDB().Create(&mhs)
	return result.Error
}