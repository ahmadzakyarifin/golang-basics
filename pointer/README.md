# 📍 Go Pointers: Understanding Memory Addresses

Materi ini membahas konsep dasar Pointer di Go: bagaimana mengelola data menggunakan alamat memori alih-alih menduplikasi nilainya.

---

## 1. Konsep Dasar
Di Go, setiap variabel disimpan di sebuah "rumah" di dalam RAM yang memiliki **Alamat Memori**.

* **Value:** Isi dari rumah tersebut (contoh: "Zaky").
* **Pointer:** Kertas berisi alamat rumah tersebut (contoh: `0xc000...`).

### Operator Penting:
1. **Operator `&` (Address-of):** Digunakan untuk mengambil alamat memori dari sebuah variabel.
2. **Operator `*` (Asterisk/Dereferencing):** - Jika di tipe data (`*string`), artinya variabel tersebut adalah pointer.
   - Jika di depan variabel pointer (`*namePtr`), artinya kita sedang mengakses/mengubah **isi** dari alamat tersebut.

---

## 2. Pass by Value vs Pass by Reference

### Pass by Value (Default Go)
Saat kita mengirim variabel ke fungsi, Go membuat **fotokopi** datanya.
* **Kelebihan:** Data asli aman dari perubahan yang tidak disengaja.
* **Kekurangan:** Memakan memori lebih banyak jika data yang disalin sangat besar.

### Pass by Reference (Menggunakan Pointer)
Saat kita mengirim **alamat memori** (`&variabel`) ke fungsi.
* **Kelebihan:** Sangat hemat memori (hanya mengirim alamat kecil) dan bisa mengubah data asli.
* **Kekurangan:** Harus hati-hati karena perubahan di dalam fungsi akan merubah data asli secara permanen.



---

## 💡 Analogi Sederhana
* **Pass by Value:** Anda memberikan **fotokopi** dokumen ke teman. Teman Anda mencoret-coret fotokopi itu, dokumen asli Anda tetap bersih.
* **Pass by Reference:** Anda memberikan **kunci rumah** ke teman. Teman Anda masuk dan mengecat tembok rumah, saat Anda pulang, tembok Anda sudah berubah warna.

---

## 🚀 Kesimpulan
Pointer bukan tentang "membuat kode jadi keren", tapi tentang **Efisiensi Memori**. Gunakan pointer saat Anda perlu mengubah data asli atau saat menangani data yang ukurannya sangat besar.