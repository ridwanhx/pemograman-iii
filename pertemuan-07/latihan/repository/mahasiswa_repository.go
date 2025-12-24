package repository

import (
	"latihan/config"
	"latihan/model"
)

func GetAllMahasiswa() ([]model.Mahasiswa, error) {
	var data []model.Mahasiswa
	result := config.GetDB().Find(&data);
	return data, result.Error
}

func InsertMahasiswa(mhs model.Mahasiswa) error {
	result := config.GetDB().Create(&mhs)
	return result.Error
}

func GetMahasiswaByNpm(npm string) ([]model.Mahasiswa, error) {
	var data []model.Mahasiswa
	result := config.GetDB().Where("npm = ?", npm).First(&data)
	return data, result.Error
}

func UpdateMahasiswaByNpm(npm string, updateData map[string]interface{}) error {
	result := config.GetDB().Model(&model.Mahasiswa{}).Where("npm = ?", npm).Updates(updateData)
	
	return result.Error
}

func DeleteMahasiswaById(id string) error {
	result := config.GetDB().Where("id = ?", id).Delete(&model.Mahasiswa{})

	return result.Error
}

// Tambahkan/Pastikan fungsi ini ada di mahasiswa_repository.go
func DeleteMahasiswaByNpm(npm string) error {
    // Gunakan Where("npm = ?", ...) agar mencari berdasarkan kolom NPM, bukan ID
    result := config.DB.Where("npm = ?", npm).Delete(&model.Mahasiswa{})
    return result.Error
}