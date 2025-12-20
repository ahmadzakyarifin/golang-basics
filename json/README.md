# 🚀 Golang JSON Handling: Marshalling & Unmarshalling

Dokumentasi ini memberikan penjelasan mendalam per-sintaks mengenai cara mengelola data format **JSON** di Go menggunakan paket standar `encoding/json`. Fokus utama materi ini adalah memahami proses **Memetakan (Mapping)** data antara format JSON mentah dengan struktur data di Go.

---

## 🔍 Apa Itu "Memetakan" (Mapping)?

**Memetakan** adalah proses menjodohkan atau menyamakan "bahasa" antara file JSON dengan kode Go. Hal ini diperlukan karena JSON dan Go memiliki aturan penulisan yang berbeda:

* 📝 **Gaya Penulisan:** JSON biasanya menggunakan huruf kecil atau *snake_case* (`nama_lengkap`), sedangkan Go mewajibkan field diawali **Huruf Kapital** (`NamaLengkap`) agar bersifat *Exported*.
* 🏷️ **JSON Tags sebagai Label:** Kita memberikan label pada struct menggunakan tanda backtick ( **\`json:"nama"\`** ) untuk memberi tahu Go: *"Jika ada data bernama 'nama' di JSON, masukkan ke field 'Nama' di kode saya"*.


---

## 📂 Struktur Data & JSON Tags

Kita menggunakan **JSON Tags** untuk menyinkronkan kunci JSON dengan variabel Go agar data tidak kosong saat diproses:

```go
type User struct {
    Nama   string `json:"nama"`   // Memetakan key "nama" ke field Nama
    Alamat string `json:"alamat"` // Memetakan key "alamat" ke field Alamat
    Umur   int    `json:"umur"`   // Memetakan key "umur" ke field Umur
}