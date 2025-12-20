# 🔄 Control Flow: Looping & Percabangan di Go

Dokumentasi ini adalah rangkuman mengenai logika pengambilan keputusan dan teknik perulangan yang saya terapkan di Golang. Fokusnya adalah pada penulisan kode yang efisien dan pemanfaatan fitur unik Go seperti *short statement*.

---

## 🔀 1. Percabangan (Conditional)

Percabangan di Go dibuat sangat minimalis namun tetap fungsional untuk menangani logika yang kompleks.

### **A. If, Else If, Else**
* 📝 **Tanpa Kurung:** Go menghilangkan kebutuhan tanda kurung `()` pada kondisi, sehingga kode terlihat lebih bersih.
* ⚡ **Short Statement:** Saya sering menggunakan deklarasi variabel singkat langsung di baris `if`. Variabel ini hanya hidup di dalam blok tersebut, sehingga membantu menjaga kebersihan memori (*Scope Management*).
* 💡 **Kegunaan:** Sangat membantu untuk validasi data atau pengecekan error secara cepat sebelum lanjut ke logika berikutnya.



[Image of if-else flowchart in programming]


### **B. Switch Case**
* 📝 **Implicit Break:** Salah satu fitur favorit saya di Go adalah tidak perlunya menulis `break` secara manual; Go akan otomatis berhenti setelah menemukan `case` yang cocok.
* ⚡ **Multi-Value:** Kita bisa menggabungkan beberapa nilai dalam satu `case` menggunakan koma untuk efisiensi baris.
* 💡 **Kapan Digunakan:** Saya memilih `switch` jika ada banyak kondisi tetap (seperti kategori menu atau status transaksi) agar kode tidak "berantakan" dengan banyak `else-if`.

---

## 🔄 2. Perulangan (Looping)

Keunikan Go adalah hanya memiliki satu kata kunci perulangan, yaitu **`for`**, namun bisa digunakan dalam berbagai gaya.

### **A. For Standar & While Style**
* 📝 **Standard For:** Digunakan untuk perulangan dengan hitungan pasti (misal: 1 sampai 10).
* 💡 **While-Like:** Cukup dengan `for [kondisi] { ... }` untuk perulangan yang bergantung pada variabel dinamis di luar loop.

### **B. For Range (Efisiensi Koleksi)**
* 📝 **Sifat:** Cara paling aman untuk membedah data di dalam **Slice, Array, atau Map** tanpa risiko error indeks.
* ⚡ **Blank Identifier (`_`):** Jika saya hanya butuh datanya tanpa butuh indeksnya, saya menggunakan tanda `_` agar kode tetap bersih dan tidak error.
* 💡 **Kegunaan:** Standar utama saya saat harus mengolah daftar data dari hasil query database atau response API.



---

## 📊 Tabel Referensi Cepat

| Logika | Sintaks | Alasan Memilih |
| :--- | :--- | :--- |
| **Dinamis** | `if-else` | Cocok untuk pengecekan perbandingan kompleks (>, <, ==). |
| **Fixed Value** | `switch` | Lebih rapi untuk mencocokkan status atau pilihan menu yang sudah pasti. |
| **Data List** | `for range` | Cara terbaik untuk memproses isi Array/Slice/Map secara menyeluruh. |

---

## 💡 Engineering Tips & Best Practices
* 🛠️ **Minimal Scope:** Selalu prioritaskan penggunaan *short statement* di `if` jika variabel tersebut tidak digunakan di tempat lain.
* 🛑 **Break & Continue:** Manfaatkan `break` untuk berhenti paksa dan `continue` untuk melompati iterasi data yang tidak valid agar performa loop lebih cepat.
* 🗑️ **Avoid Infinite Loop:** Selalu pastikan ada kondisi yang akan bernilai `false` agar aplikasi tidak memakan CPU secara berlebihan.

---

## 🚀 Cara Menjalankan
Pastikan kamu berada di direktori yang benar, lalu jalankan:
```bash
go run main.go