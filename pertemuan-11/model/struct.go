package model


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

// TableName memastikan GORM mencari tabel bernama "users" di Supabase
func (User) TableName() string { return "users" }