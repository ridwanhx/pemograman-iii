// --------------------------------------------
// PERTEMUAN 04 - Pengenalan Web Service & JSON
// --------------------------------------------

// # Apa itu JSON?
// JSON (JavaScript Object Notation) adalah format data ringan untuk pertukaran informasi antar sistem
// Sederhananya, JSON mirip seperti tabel data, tapi berbentuk teks
// Contoh:
/*
	{
		"name": "Muhammad Ridwan",
		"npm": 714251017,
		"email": "ridwan@mail.com",
		"age": 30,
		"hobbies": ["Coding", "Streaming", "Learning"],
		"is_active": true
	}
*/

// Struktur data JSON
// -- Objek JSON: {...}
// Gunakan tanda kurung kurawal {}

// -- Array JSON: [...]
// Gunakan tanda kurung siku []

// Struktur dasar JSON
// Nilai JSON dapat berupa:
// String, Angka, Boolean, Array, Object


// Kelebihan & (Kekurangan) JSON dibanding format lain
// JSON -> Ringan, mudah dibaca, mudah digunakan di berbagai bahasa pemograman (Tidak cocok untuk data yang skala nya sangat besar)
// XML -> Kuat dan formal (Berat, banyak tag, sulit dibaca manusia)
// CSV -> Sederhana (Tidak bisa menyimpan struktur kompleks)

// Karena itu, hampir di semua API modern (GO, PHP< Python, Java, Node.js) menggunakan JSON sebagai format standar

// Membuat JSON
// Ketikkan script JSON seperti pada contoh diatas
// Simpan file dalam format .json

// Mengecek Validitas JSON
// Gunakan situs https://jsonlint.com
// Copy-paste isi JSON -> klik Validate JSON
// Jika muncul pesan "Valid JSON", berarti formatnya sudah benar.