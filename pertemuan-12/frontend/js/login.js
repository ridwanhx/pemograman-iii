document.addEventListener("DOMContentLoaded", () => {
  // ambil elemen form login
  const form = document.getElementById("form-login");

  form.addEventListener("submit", async (e) => {
    // cegah agar nilai tidak di submit oleh form
    e.preventDefault();

    // ambil elemen input (username & password)
    const username = document.getElementById("username").value.trim();
    const password = document.getElementById("password").value.trim();

    // jika username atau password tidak mengirimkan nilai (undefined, null, string kosong)
    if (!username || !password) {
      alert("Username dan password wajib diisi");
      return;
    }

    // do try catch
    try {
      // Inisialisasi request login menggunakan fetch API
      const res = await fetch(API_LOGIN, {
        // Method HTTP yang digunakan adalah POST karena kita mengirimkan data (username & password)
        method: "POST",
        // Header untuk memberitahu server bahwa data yang dikirim berupa JSON
        headers: { "Content-Type": "application/json" },
        // Body request berisi data login (username & password) dalam format JSON string
        body: JSON.stringify({ username, password }),
      });

      // Parsing response dari server menjadi format JSON agar mudah diolah
      const json = await res.json();

      // Jika response tidak OK (status bukan 200-299), berarti login gagal
      if (!res.ok) {
        // Tampilkan pesan error dari server jika ada, jika tidak ada gunakan pesan default "Login gagal"
        alert(json?.message || "Login gagal");
        // Hentikan eksekusi fungsi agar tidak lanjut ke proses berikutnya
        return;
        }
        
        // simpan token
        setToken(json.token);
        alert("Login berhasil!");
        window.location.href = "index.html";
    } catch (err) {
        alert("Gagal terhubung ke server saat login");
        console.error(err);
    }
  });
});
