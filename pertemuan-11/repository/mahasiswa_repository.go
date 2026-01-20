package repository

import (
    "pertemuan-11/config"
    "pertemuan-11/model"
)

// Ambil semua data mahasiswa
func GetAllMahasiswa() ([]model.Mahasiswa, error) {
    var data []model.Mahasiswa
    result := config.GetDB().Find(&data)
    return data, result.Error
}

// Insert mahasiswa baru
func InsertMahasiswa(mhs model.Mahasiswa) error {
    result := config.GetDB().Create(&mhs)
    return result.Error
}

// Cari mahasiswa berdasarkan NPM
func GetMahasiswaByNpm(npm string) (model.Mahasiswa, error) {
    var data model.Mahasiswa
    result := config.GetDB().Where("npm = ?", npm).First(&data)
    return data, result.Error
}

// Update mahasiswa berdasarkan NPM
func UpdateMahasiswaByNpm(npm string, updateData map[string]interface{}) error {
    result := config.GetDB().Model(&model.Mahasiswa{}).Where("npm = ?", npm).Updates(updateData)
    return result.Error
}

// Delete mahasiswa berdasarkan ID
func DeleteMahasiswaById(id string) error {
    result := config.GetDB().Where("id = ?", id).Delete(&model.Mahasiswa{})
    return result.Error
}

// Delete mahasiswa berdasarkan NPM
func DeleteMahasiswaByNpm(npm string) error {
    result := config.GetDB().Where("npm = ?", npm).Delete(&model.Mahasiswa{})
    return result.Error
}
