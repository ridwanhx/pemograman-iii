document.addEventListener("DOMContentLoaded", async () => {
  const btnLogout = document.getElementById("btn-logout");

  // tombol muncul hanya kalau ada token
  if (btnLogout) {
    btnLogout.style.display = getToken() ? "inline-block" : "none";

    btnLogout.addEventListener("click", () => {
      const yakin = confirm("Yakin ingin logout?");
      if (!yakin) return; // kalau Cancel, hentikan

      clearToken();
      alert("Logout berhasil");
      window.location.href = "login.html";
    });
  }

  // wajib login
  if (!requireAuth()) return;

  const tbody = document.getElementById("tbody-mahasiswa");
  const jumlahMhs = document.getElementById("jumlah-mahasiswa");

  if (!tbody) {
    console.error(`Element tbody dengan id="tbody-mahasiswa" tidak ditemukan`);
  }

  // Load data
  try {
    const response = await authFetch(API_MAHASISWA);
    const hasil = await response.json();

    const data = hasil?.data ?? [];

    if (jumlahMhs) jumlahMhs.textContent = String(data.length);

    if (!data.length) {
      tbody.innerHTML = `
            <tr> 
                <td colspan="6" class="text-center text-gray-500 py-4"> 
                    Belum ada data mahasiswa 
                </td> 
            </tr>
            `;
      return;
    }
    tbody.innerHTML = data
      .map((mhs) => {
        const hobi = Array.isArray(mhs.hobi)
          ? mhs.hobi.join(", ")
          : mhs.hobi || "-";

        const npm = mhs.npm || "-";

        return ` 
          <tr class="h-18 border-b border-coolGray-100"> 
            <td class="px-4 bg-white text-sm font-medium">${npm}</td> 
            <td class="px-4 bg-white text-sm font-medium">${
              mhs.nama || "-"
            }</td> 
            <td class="px-4 bg-white text-sm">${mhs.jurusan || "-"}</td> 
            <td class="px-4 bg-white text-sm">${mhs.alamat || "-"}</td> 
            <td class="px-4 bg-white text-sm">${hobi}</td> 
            <td class="px-4 bg-white text-sm"> 
              <button class="btn-edit bg-yellow-400 text-white px-2 py-1 
rounded text-xs" data-npm="${npm}"> 
                Edit 
              </button> 
              <button class="btn-delete bg-red-500 text-white px-2 py-1 rounded 
text-xs ml-1" data-npm="${npm}"> 
                Hapus 
              </button> 
            </td> 
          </tr> 
        `;
      })
      .join("");

    // ===== Event delegation untuk Edit & Delete =====
    tbody.addEventListener("click", async (e) => {
      const target = e.target;

      // EDIT
      if (target.classList.contains("btn-edit")) {
        const npm = target.dataset.npm;
        window.location.href = `update.html?npm=${encodeURIComponent(npm)}`;
        return;
      }

      // DELETE
      if (target.classList.contains("btn-delete")) {
        const npm = target.dataset.npm;

        if (!confirm(`Yakin hapus mahasiswa NPM ${npm}?`)) return;

        try {
          const res = await authFetch(`${API_MAHASISWA}/${npm}`, {
            method: "DELETE",
          });
          const json = await res.json();

          if (!res.ok) {
            alert(json?.message || "Gagal hapus data");
            return;
          }

          alert(json?.message || "Berhasil hapus data");
          window.location.reload();
        } catch (err) {
          console.error(err);
          alert("Terjadi error saat menghapus data");
        }
      }
    });
  } catch (err) {
    console.error(err);
    tbody.innerHTML = ` 
      <tr> 
        <td colspan="6" class="text-center text-red-500 py-4"> 
          Gagal mengambil data mahasiswa 
        </td> 
      </tr> 
    `;
  }
});
