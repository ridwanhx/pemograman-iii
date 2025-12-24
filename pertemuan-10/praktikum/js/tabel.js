// Melakukan Fetch API
// fetch("http://localhost:3000/api/mahasiswa")
//   .then((res) => res.json())
//   .then((data) => console.info(data));

// const URL_API_MAHASISWA = "http://localhost:3000/api/mahasiswa";

document.addEventListener("DOMContentLoaded", function () {
  const tbody = document.getElementById("tbody-mahasiswa");
  const countMhs = document.getElementById("jumlah-mahasiswa");

  fetch(API_MAHASISWA)
    .then((response) => response.json())
    .then((result) => {
      const data = result.data;
      let rows = "";

      if (countMhs) {
        countMhs.textContent = data.length ? data.length : "0";
      }

      data.forEach((mhs) => {
        let hobi = Array.isArray(mhs.hobi)
          ? mhs.hobi.join(", ")
          : mhs.hobi || "-";

        rows += `<tr class="h-18 border-b border-coolGray-100">
  <td class="px-4 bg-white text-sm font-medium">${mhs.npm || "-"}</td>
  <td class="px-4 bg-white text-sm font-medium">${mhs.nama || "-"}</td>
  <td class="px-4 bg-white text-sm">${mhs.jurusan || "-"}</td>
  <td class="px-4 bg-white text-sm">${mhs.alamat || "-"}</td>
  <td class="px-4 bg-white text-sm">${mhs.hobi || "-"}</td>
  <td class="px-4 bg-white text-sm">
  <button class="btn-edit bg-yellow-400 text-white px-2 py-1 rounded text-xs" data-npm="${
    mhs.npm
  }">Edit</button>
  <button class="btn-delete bg-red-500 text-white px-2 py-1 rounded text-xs" data-npm="${
    mhs.npm
  }">Hapus</button>
  </td>
</tr>`;
      });

      tbody.innerHTML = rows;

      // tombol edit
      document.querySelectorAll(".btn-edit").forEach((btn) => {
        btn.addEventListener("click", () => {
          const npm = btn.dataset.npm;
          window.location.href = `update.html?npm=${encodeURIComponent(npm)}`;
        });
      });

      // tombol delete
      document.querySelectorAll(".btn-delete").forEach((btn) => {
        btn.addEventListener("click", async () => {
          const npm = btn.dataset.npm;

          if (
            !confirm(`Yakin ingin menghapus data mahasiswa dengan NPM ${npm}`)
          )
            return;

          try {
            const res = await fetch(`${API_MAHASISWA}/${npm}`, {
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
            alert("Error koneksi ke server saat delete");
            console.error(err);
          }
        });
      });
    })
    .catch(() => {
      tbody.innerHTML = `<tr>
  <td colspan="5" class="text-center text-red-500 py-4">Gagal mengambil data mahasiswa</td>
</tr>`;
    });
});
