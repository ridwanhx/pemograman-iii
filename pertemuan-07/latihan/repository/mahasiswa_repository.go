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