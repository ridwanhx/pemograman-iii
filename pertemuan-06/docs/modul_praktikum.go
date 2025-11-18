// --------------------------------------------------------------
// MEMBANGUN REST API (GET dan POST) dengan Go (Fiber) + Supabase
// --------------------------------------------------------------

// (1) Requirements
// Golang >= 1.21
// Git
// VS Code + ekstensi Go
// Akun Supabase + proyek aktif (PostgreSQL)
// Pengetahuan dasar HTTP/JSON

// Package Go yang akan digunakan:
// github.com/gofiber/fiber/v2	// framework Go / fiber
// gorm.io/gorm dan gorm.io/driver/postgres	// Object Relational Mapping (ORM)
// github.com/joho/godotenv	// memuat variabel dari file .env
// github.com/lib/pq	// untuk berinteraksi dengan database PostgreSQL

// Struktur Folder
/*
latihan/ 
├── config/ 
│   └── db.go 
├── handler/ 
│   └── mahasiswa_handler.go 
├── model/ 
│   └── struct.go 
├── repository/ 
│   └── mahasiswa_repository.go 
├── router/ 
│   └── routes.go 
├── main.go 
└── .env 
*/

// Goals
// -- Membuat REST API sederhana dengan dua endpoint:
// GET /api/<nama_data>	-> ambil semua data
// POST /api/<nama_data>	-> tambah data baru

// Alur data
// Client (Postman) -> Router -> Handler -> Repository (GORM) -> PostgreSQL (Supabase)

// Peran masing-masing komponen
// (1) Router: Memetakan URL -> handler mana yang dieksekusi.
// (2) Handler: Logika HTTP; validasi input, panggil repository, balas JSON.
// (3) Repository: Berkomunikasi dengan database (via GORM).
// (4) Model: Struktur data / tabel.
// (5) Config: Inisialisasi koneksi database (Supabase DSN).
// (6) main.go: titik masuk; init DB, auto-migrate, jalankan server
