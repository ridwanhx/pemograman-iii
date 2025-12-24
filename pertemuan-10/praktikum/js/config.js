const API_BASE = "http://127.0.0.1:3000/api";
const API_MAHASISWA = `${API_BASE}/mahasiswa`;

// helper kecil untuk parsing hobi dari input "a,b,c" => ["a", "b", "c"]
function parseHobi(input) {
  if (!input) return [];

  return input
    .split(", ")
    .map((x) => x.trim())
    .filter((x) => x.length > 0);
}
