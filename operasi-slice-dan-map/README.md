# 🏛️ Operasi Slice dan Map: Pendalaman Struktur Data Dinamis

Repositori ini berisi laboratorium praktis mengenai dua struktur data koleksi utama di Golang: **Slice** dan **Map**. Proyek ini bertujuan untuk membedah bagaimana Go mengelola memori, kapasitas, dan sifat **Tipe Referensi** (*Reference Type*).

---

## 📦 1. Eksplorasi Slice (Array Dinamis)

Slice di Go bukan sekadar array, melainkan sebuah *header* yang merujuk ke array dasar.

### Poin Utama Eksperimen:
* **Mekanisme Referensi:** Saat melakukan slicing (`array[1:3]`), slice berbagi alamat memori dengan array asli. Mengubah isi slice akan berdampak langsung pada array tersebut.
* **Manajemen Kapasitas (`cap`):** Memahami bagaimana Go memesan ruang memori.
* **Efek `append` & Alokasi Ulang:** Jika `append` dilakukan melampaui kapasitas (`cap`), Go akan membuat array baru di memori. Pada titik ini, slice "lepas" dari referensi array aslinya.



---

## 🗺️ 2. Eksplorasi Map (Hash Table)

Map adalah kumpulan pasangan kunci-nilai (*key-value*) yang sangat efisien untuk pencarian data (O(1)).

### Poin Utama Eksperimen:
* **Inisialisasi Aman:** Praktik menggunakan `make()` dan literal `{}` untuk menghindari *Nil Map* yang dapat menyebabkan aplikasi *panic*.
* **Operasi CRUD:** Demonstrasi cara menambah, mengubah, menghapus (`delete`), dan melihat panjang data (`len`).
* **Sifat Referensi Murni:** Berbeda dengan tipe data dasar, Map selalu berbagi alamat memori yang sama saat disalin ke variabel baru.



---

## 🛠️ Ringkasan Perbandingan Teknis

| Fitur | Slice | Map |
| :--- | :--- | :--- |
| **Identitas** | Berdasarkan Index (0, 1, 2...) | Berdasarkan Kunci (*Key*) |
| **Tipe Memori** | Referensi (Pointer ke Array Dasar) | Referensi (Pointer ke hmap) |
| **Fungsi Utama** | `append()`, `len()`, `cap()` | `delete()`, `len()` |
| **Kapan Berubah?** | Selalu sinkron (kecuali jika `append` jebol kapasitas) | Selalu sinkron antar variabel alias |

---

## 💡 Kesimpulan: "Pass-by-Reference"

Melalui eksperimen ini, dapat disimpulkan bahwa Go sangat efisien dalam mengelola data besar. Daripada melakukan fotokopi (*copy*) data yang berat, Go menggunakan pointer internal untuk berbagi akses data asli. 

> **Insight:** Developer harus berhati-hati saat melakukan `append` pada slice yang berasal dari array, karena hubungan referensi bisa terputus secara otomatis saat terjadi alokasi memori baru.



---