# 🚀 Go Memory Mutation Lab: Update & Pointers

Repositori ini mendokumentasikan eksperimen mendalam mengenai berbagai cara memperbarui (**Update**) nilai data di dalam bahasa Go. Fokus utama adalah membedah perbedaan perilaku memori antara **Value Type** dan **Reference Type**.

## 📌 Deskripsi Eksperimen
Eksperimen ini membagi cara mutasi data menjadi 4 model utama:
1. **Langsung**: Mengubah nilai variabel secara direct.
2. **Variabel Lain**: Menguji apakah perubahan pada alias mempengaruhi data asli.
3. **Variabel Pointer**: Menggunakan operator `&` dan `*` untuk memanipulasi alamat memori.
4. **Fungsi & Receiver**: Menggunakan parameter pointer dan *Pointer Receiver* pada Struct.

---

## 🧠 Konsep Penting

### 1. Value Type vs Reference Type
Di Go, setiap tipe data memiliki cara berbeda saat berpindah antar variabel atau fungsi:

* **Value Type (`int`, `Array`, `Struct`)**: Bersifat seperti fotokopi. Saat dikirim ke fungsi atau variabel lain, Go menyalin nilainya. Mengubah salinan **TIDAK** mempengaruhi data asli.
* **Reference Type (`Slice`, `Map`)**: Bersifat seperti alamat rumah. Variabel hanya menyimpan referensi ke lokasi data sebenarnya. Perubahan pada satu alias **AKAN** mempengaruhi data asli.



### 2. Header pada Slice
Meskipun Slice terlihat seperti referensi, ia sebenarnya adalah sebuah *struct* kecil (Header) yang berisi: **Pointer ke Array Dasar, Length, dan Capacity**. Inilah alasan mengapa mengubah isi index slice di dalam fungsi akan merubah data asli tanpa perlu menggunakan pointer `*[]`.



---

## 🛠️ Ringkasan Temuan (Update Tabel)

| Tipe Data | Model Mutasi | Efek ke Data Asli | Mengapa? |
| :--- | :--- | :--- | :--- |
| **int** | Variabel Lain | ❌ Tetap | Pass-by-value (Fotokopi) |
| **int** | Pointer/Fungsi `*` | ✅ Berubah | Mengakses alamat memori asli |
| **Array** | Pointer/Fungsi `*` | ✅ Berubah | Mencegah penyalinan data besar |
| **Slice** | Variabel Lain/Alias| ✅ Berubah | Berbagi pointer array dasar yang sama |
| **Map** | Fungsi Biasa | ✅ Berubah | Map secara internal adalah pointer |
| **Struct**| Pointer Receiver | ✅ Berubah | Memanipulasi field di alamat asli |

---

## 🧪 Analisis Kode

### Pointer Receiver pada Struct
Pada eksperimen Struct, digunakan metode **Pointer Receiver**:
```go
func (m *Mahasiswa) updateReceiver(baru string) { m.Nama = baru }