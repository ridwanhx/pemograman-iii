// Melakukan Fetch API
// fetch("http://localhost:3000/api/mahasiswa")
//   .then((res) => res.json())
//   .then((data) => console.info(data));

const URL_API_MAHASISWA = "http://localhost:3000/api/mahasiswa";

document.addEventListener("DOMContentLoaded", function () {
    const tbody = document.getElementById("tbody-mahasiswa");
    const countMhs = document.getElementById("jumlah-mahasiswa");

  fetch(URL_API_MAHASISWA)
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
</tr>`;
      });

      tbody.innerHTML = rows;
    })
    .catch(() => {
      tbody.innerHTML = `<tr>
  <td colspan="5" class="text-center text-red-500 py-4">Gagal mengambil data mahasiswa</td>
</tr>`;
    });
});
