const API_BASE = "http://127.0.0.1:3000/api";
const API_MAHASISWA = `${API_BASE}/mahasiswa`;
const API_LOGIN = `${API_BASE}/login`;

// helper kecil untuk parsing hobi dari input a,b,c -> ["a", "b", "c"]
function parseHobi(input) {
  // jika input yang dikirimkan kosong (null, undefined, atau string kosong)
  // langsung kembalikan array kosong agar tidak terjadi error
  if (!input) return [];

  // jika input memiliki nilai, lakukan parsing
  return (
    input
      // pisahkan string berdasarkan tanda koma (,)
      .split(",")
      // hilangkan spasi di awal dan akhir setiap elemen array
      .map((x) => x.trim())
      // buang elemen yang panjangnya 0 (string kosong) agar hasil lebih bersih
      .filter((x) => x.length > 0)
  );
}

// JWT Helpers
const TOKEN_KEY = "token";

// Fungsi setToken menerima parameter 'token' yang akan disimpan ke localStorage
function setToken(token) {
  // Simpan nilai token ke dalam localStorage dengan key bernama TOKEN_KEY
  // localStorage.setItem(key, value) digunakan untuk menyimpan data secara permanen di browser
  localStorage.setItem(TOKEN_KEY, token);
}

// inisialisasi fungsi u/ mengambil token yang sebelumnya disimpan di localStorage
function getToken() {
  // ambil data dari localStorage dengan key bernama TOKEN_KEY
  return localStorage.getItem(TOKEN_KEY);
}

// fungsi u/ menghapus token
function clearToken() {
  // hapus token dari browser dengan key bernama TOKEN_KEY
  localStorage.removeItem(TOKEN_KEY);
}

// inisialisasi aturan auth (jika tidak punya token, lempar kembali ke halaman login)
function requireAuth() {
  // inisialisasi var token yang nilainya diambil dari fungsi getToken diatas
  const token = getToken();
  if (!token) {
    window.location.href = "login.html";
    return false;
  }
  return true;
}

// fetch yang otomatis bawa Authorization: Bearer <token>
async function authFetch(url, options = {}) {
  const token = getToken();

  // Buat object headers:
  // - Ambil headers bawaan dari parameter options (jika ada)
  // - Jika token tersedia, tambahkan header Authorization dengan format Bearer <token>
  const headers = {
    ...(options.headers || {}),
    ...(token ? { Authorization: `Bearer ${token}` } : {}),
  };

  // Lakukan request menggunakan fetch dengan menggabungkan options dan header baru
  const res = await fetch(url, { ...options, headers });

  // kalau token invalid / expired -> paksa login ulang
  if (res.status === 401) {
    // hapus token dari localStorage
    clearToken();
    // tampilkan pesan melalui alert
    alert("Sesi habis / token tidak valid. Silahkan login ulang.");
    // redirect kembali ke halaman login
    window.location.href = "login.html";
    // Reject promise agar bisa ditangani di caller
    return Promise.reject(new Error("Unauthorized"));
  }

  // kembalikan response hasil fetch
  return res;
}
