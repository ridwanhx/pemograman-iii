package model

import (
	"github.com/lib/pq"
)

// inisialisasi structs
type Movies struct {
	Id string `json:"id" gorm:"column:id;primaryKey;type:uuid;default:gen_random_uuid()"`
	Title string `json:"title" gorm:"column:title;type:varchar(100);not null"`
	Director string `json:"director" gorm:"column:director;type:varchar(100);not null"`
	Categories pq.StringArray `json:"categories" gorm:"column:categories;type:text[];not null"`
	Casts map[string]string `json:"casts" gorm:"column:casts;type:jsonb;not null"`
	Synopsis string `json:"synopsis" gorm:"column:synopsis;type:text"`
	Rating float64 `json:"rating" gorm:"column:rating;type:float;not null"`
	ReleaseYear int `json:"release_year" gorm:"column:release_year;type:integer"`
}

type Mahasiswa struct {
	Id string `json:"id" gorm:"column:id;primaryKey;type:uuid;default:gen_random_uuid()"`
	Npm string `json:"npm" gorm:"column:npm;type:char(9);not null"`
	Nama string `json:"nama" gorm:"column:nama;type:varchar(100);not null"`
	Email string `json:"email" gorm:"column:email;type:varchar(100);not null"`
	Jurusan string `json:"jurusan" gorm:"column:jurusan;type:varchar(100)"`
	Ipk float64 `json:"ipk" gorm:"column:ipk;type:float;not null"`
	Alamat string `json:"alamat" gorm:"column:alamat;type:text;not null"`
	Hobi pq.StringArray `json:"hobi" gorm:"column:hobi;type:text[]"`
}

// Menentukan nama tabel di database untuk struct Movies
func (Movies) TableName() string {
	return "movies";
}

func (Mahasiswa) TableName() string {
	return "mahasiswa";
}