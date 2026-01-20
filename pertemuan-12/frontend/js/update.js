document.addEventListener("DOMContentLoaded", async () => {
  // Wajib login
  if (!requireAuth()) return;

  const form = document.getElementById("form-update");
  const btnCancel = document.getElementById("btn-cancel");

  // Ambil parameter NPM dari URL
  const params = new URLSearchParams(window.location.search);
  const npmParam = params.get("npm");

  if (!npmParam) {
    alert("NPM tidak ditemukan di URL. Contoh: update.html?npm=1224001");
    window.location.href = "index.html";
    return;
  }

  // Tombol cancel
  btnCancel?.addEventListener("click", () => {
    window.location.href = "index.html";
  });

  // 1) Load detail mahasiswa
  try {
    const res = await authFetch(`${API_MAHASISWA}/${npmParam}`);
    let json;
    try {
      json = await res.json();
    } catch {
      alert("Response server bukan JSON");
      return;
    }

    if (!res.ok) {
      alert(json?.message || "Gagal ambil detail mahasiswa");
      return;
    }

    const mhs = json.data || {};
    document.getElementById("npm").value = mhs.npm || npmParam;
    document.getElementById("nama").value = mhs.nama || "";
    document.getElementById("jurusan").value = mhs.jurusan || "";
    document.getElementById("alamat").value = mhs.alamat || "";
    document.getElementById("hobi").value = Array.isArray(mhs.hobi)
      ? mhs.hobi.join(",")
      : (mhs.hobi || "");
  } catch (err) {
    alert("Error koneksi ke server saat ambil detail");
    console.error(err);
    return;
  }

  // 2) Submit update mahasiswa
  form.addEventListener("submit", async (e) => {
    e.preventDefault();

    const npm = document.getElementById("npm").value.trim();
    const nama = document.getElementById("nama").value.trim();
    const jurusan = document.getElementById("jurusan").value.trim();
    const alamat = document.getElementById("alamat").value.trim();
    const hobiStr = document.getElementById("hobi").value.trim();

    const payload = {
      npm,
      nama,
      jurusan,
      alamat,
      hobi: parseHobi(hobiStr),
    };

    try {
      const res = await authFetch(`${API_MAHASISWA}/${npm}`, {
        method: "PUT",
        headers: { "Content-Type": "application/json" },
        body: JSON.stringify(payload),
      });

      let json;
      try {
        json = await res.json();
      } catch {
        alert("Response server bukan JSON");
        return;
      }

      if (!res.ok) {
        alert(json?.message || "Gagal update data");
        return;
      }

      alert(json?.message || "Berhasil update data");
      window.location.href = "index.html";
    } catch (err) {
      console.error(err);
      alert("Terjadi error saat update data");
    }
  });
});
