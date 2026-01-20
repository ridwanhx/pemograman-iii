document.addEventListener("DOMContentLoaded", () => {
  if (!requireAuth()) return;
  const form = document.getElementById("form-insert");
  const btnCancel = document.getElementById("btn-cancel");
  btnCancel?.addEventListener("click", () => {
    window.location.href = "index.html";
  });
  form.addEventListener("submit", async (e) => {
    e.preventDefault();

    const npm = document.getElementById("npm").value.trim();
    const nama = document.getElementById("nama").value.trim();
    const jurusan = document.getElementById("jurusan").value.trim();
    const alamat = document.getElementById("alamat").value.trim();
    const hobiStr = document.getElementById("hobi").value.trim();

    if (!npm || !nama) {
      alert("NPM dan Nama wajib diisi!");
      return;
    }

    const payload = { npm, nama, jurusan, alamat, hobi: parseHobi(hobiStr) };

    try {
      const res = await authFetch(API_MAHASISWA, {
        method: "POST",
        headers: { "Content-Type": "application/json" },
        body: JSON.stringify(payload),
      });

      const json = await res.json();

      if (!res.ok) {
        alert(json?.message || "Gagal insert data");
        return;
      }

      alert(json?.message || "Berhasil insert data");
      window.location.href = "index.html";
    } catch (err) {
      console.error(err);
    }
  });
});
