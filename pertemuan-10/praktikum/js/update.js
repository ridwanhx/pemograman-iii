document.addEventListener("DOMContentLoaded", async () => {
  const form = document.getElementById("form-update");
  const btnCancel = document.getElementById("btn-cancel");

  const params = new URLSearchParams(window.location.search);
  const npmParam = params.get("npm");

  if (!npmParam) {
    alert("NPM tidak ditemukan di URL.");
    window.location.href = "index.html";
    return;
  }

  btnCancel?.addEventListener("click", () => {
    window.location.href = "index.html";
  });

  // 1) Load detail
  try {
    const res = await fetch(`${API_MAHASISWA}/${npmParam}`);
    const json = await res.json();

    if (!res.ok) {
      alert(json?.message || "Gagal ambil detail mahasiswa");
      return;
    }

    const mhs = Array.isArray(json.data)
      ? json.data[0]
      : Array.isArray(json)
      ? json[0]
      : json.data;

    console.info(mhs);

    document.getElementById("npm").value = mhs.npm || npmParam;
    document.getElementById("nama").value = mhs.nama || "";
    document.getElementById("prodi").value = mhs.jurusan || "";
    document.getElementById("alamat").value = mhs.alamat || "";
    document.getElementById("hobi").value =
      mhs.hobi || Array.isArray(mhs.hobi) ? mhs.hobi.join(",") : mhs.hobi || "";
  } catch (err) {
    alert("Error koneksi ke server saat ambil detail.");
    console.error(err);
    return;
  }

  // 2) Submit update
  form.addEventListener("submit", async (e) => {
    e.preventDefault();

    const npm = document.getElementById("npm").value.trim();
    const nama = document.getElementById("nama").value.trim();
    const jurusan = document.getElementById("prodi").value.trim();
    const alamat = document.getElementById("alamat").value.trim();
    const hobiStr = document.getElementById("hobi").value.trim();

    function parseHobi(hobiStr) {
      if (!hobiStr) return [];
      // Mengubah "Membaca,Coding" menjadi ["Membaca", "Coding"]
      return hobiStr
        .split(",")
        .map((item) => item.trim())
        .filter((item) => item !== "");
    }

    const payload = { npm, nama, jurusan, alamat, hobi: parseHobi(hobiStr) };

    try {
      const res = await fetch(`${API_MAHASISWA}/${npm}`, {
        method: "PATCH",
        headers: { "Content-Type": "application/json" },
        body: JSON.stringify(payload),
      });

      const json = await res.json();

      if (!res.ok) {
        alert(json?.message || "Gagal update data");
        return;
      }

      alert(json?.message || "Berhasil update data");
      window.location.href = "index.html";
    } catch (err) {
      alert("Error koneksi ke server saat update");
      console.error(err);
    }
  });
});
