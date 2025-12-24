document.addEventListener("DOMContentLoaded", async () => {
  const form = document.getElementById("form-insert");
  const btnCancel = document.getElementById("btn-cancel");

  btnCancel?.addEventListener("click", () => {
    window.location.href = "index.html";
  });

  form.addEventListener("submit", async (e) => {
    e.preventDefault();

    const npm = document.getElementById("npm").value.trim();
    const nama = document.getElementById("nama").value.trim();
    const jurusan = document.getElementById("prodi").value.trim();
    const alamat = document.getElementById("alamat").value.trim();
    const hobiStr = document.getElementById("hobi").value.trim();
    if (!npm || !nama) {
      alert("Field NPM dan Nama wajib diisi!");
      return;
    }

    const payload = {
      npm,
      nama,
      jurusan,
      alamat,
      hobi: parseHobi(hobiStr),
    };

    try {
      const res = await fetch(API_MAHASISWA, {
        method: "POST",
        headers: { "Content-Type": "application/json" },
        body: JSON.stringify(payload),
      });

      const json = await res.json();

      if (!res.ok) {
        alert(json?.message || "Gagal insert data");
        return;
      }

      alert(json?.message || "Berhasil melakukan insert dataa");
      window.location.href = "index.html";
    } catch (err) {
      alert("Error koneksi ke server");
      console.error(err);
    }
  });
});
