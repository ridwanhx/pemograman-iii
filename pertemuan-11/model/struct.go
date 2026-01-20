package model

import (
	"github.com/lib/pq"
)

type User struct {
    ID string `json:"id" gorm:"column:id;primaryKey;type:uuid;default:gen_random_uuid()"`
    Username string `gorm:"unique;not null" json:"username"`
    // Password ditandai json:"-" agar tidak muncul saat data dikirim ke user
    Password string `gorm:"not null" json:"-"` 
    Role     string `gorm:"default:user" json:"role"`
}

type CreateUserRequest struct {
    Username string `json:"username"`
    Password string `json:"password"`
}

type LoginRequest struct {
    Username string `json:"username"`
    Password string `json:"password"`
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

// TableName memastikan GORM mencari tabel bernama "users" di Supabase
func (User) TableName() string { return "users" }

func (Mahasiswa) TableName() string {
	return "mahasiswa";
}