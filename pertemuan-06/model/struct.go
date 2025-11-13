package model

import "github.com/lib/pq"

type Mahasiswa struct {
	// gorm itu sama kaya ORM
	// reminder: buat struct itu pake backtick
	NPM string `json:"npm" gorm:"column:npm;primaryKey;type:varchar(20);not null"`
	Nama string `json:"nama" gorm:"column:nama;type:varchar(100);not null"`
	Prodi string `json:"prodi" gorm:"column:prodi;type:varchar(100);not null"`
	Alamat string `json:"alamat" gorm:"column:alamat;type:varchar(200);not null"`
	Hobi   pq.StringArray `json:"hobi" gorm:"column:hobi;type:text[]"`
}

// deklarasi nama tabel
// func (<nama_struct>) TableName() string
func (Mahasiswa) TableName() string {
	// inisiasi nama tabel = mahasiswa
	return "mahasiswa"
}