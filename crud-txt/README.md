# 🏦 Go File System: Fundamental CRUD & Persistence

Project ini mendemonstrasikan bagaimana Golang berinteraksi dengan sistem penyimpanan (Disk) melalui file teks (.txt). Fokus utama materi ini adalah memahami siklus hidup data, izin akses file, dan perbedaan antara memori sementara (RAM) vs penyimpanan permanen (Disk).

---

## 🏗️ Implementasi CRUD pada File TXT

Berbeda dengan database, operasi CRUD pada file teks di Go dilakukan dengan menggunakan package `os` dan `bufio`. Berikut adalah implementasinya:

| Operasi | Fungsi Golang | Deskripsi |
| :--- | :--- | :--- |
| **CREATE** | `os.Create()` | Membuat file baru. Jika file sudah ada, isinya akan di-reset (**Truncate**). |
| **READ** | `os.Open()` & `bufio.Scanner` | Membaca file baris demi baris (**Streaming**). Sangat efisien terhadap penggunaan RAM. |
| **UPDATE (Append)** | `os.OpenFile(O_APPEND)` | Menambahkan data baru di baris terakhir tanpa menghapus data lama. Cocok untuk *Logging*. |
| **UPDATE (Overwrite)**| `os.WriteFile()` | Menghapus seluruh isi lama dan menggantinya dengan data baru dalam bentuk `[]byte`. |
| **DELETE** | `os.Remove()` | Menghapus file secara permanen dari sistem operasi. |

---

## 🔍 Open vs. OpenFile: Mana yang Harus Digunakan?

Pemilihan fungsi pembukaan file bergantung pada **niat (intent)** terhadap file tersebut:

1. **`os.Open(name)`**
   * **Mode**: `Read-Only` (Hanya Baca).
   * **Kapan digunakan?** Saat ingin menampilkan isi laporan atau membaca konfigurasi. Paling aman karena mencegah perubahan data yang tidak sengaja.

2. **`os.OpenFile(name, flag, perm)`**
   * **Mode**: Fleksibel (Write, Append, Create).
   * **Flags Penting**:
     * `os.O_APPEND`: Menambah data di baris bawah.
     * `os.O_CREATE`: Membuat file otomatis jika belum ada.
     * `os.O_WRONLY`: Akses hanya untuk menulis.
   * **Kapan digunakan?** Saat perlu mencatat log transaksi yang terus bertambah.



---

## ⚠️ Engineering Insights: RAM vs Disk

Project ini memberikan pemahaman penting tentang **Persistence**:
* **Data di RAM (Volatile)**: Data yang diproses melalui Goroutine atau disimpan di variabel. Sangat cepat, tapi akan hilang jika program berhenti atau laptop mati.
* **Data di Disk (Non-Volatile)**: Data yang ditulis ke file TXT. Tetap ada selamanya (permanen) meskipun sistem dimatikan.



---

## 🛡️ Best Practices yang Diterapkan

* **Error Handling**: Setiap interaksi I/O selalu divalidasi dengan pengecekan `err` untuk mencegah *crash* akibat masalah izin akses atau disk penuh.
* **Resource Cleanup**: Menggunakan `defer file.Close()` untuk menjamin *file descriptor* ditutup, mencegah kebocoran memori (**Memory Leak**).
* **Byte Conversion**: Menggunakan konversi `[]byte` pada fungsi `WriteFile`, memahami bahwa penyimpanan fisik memproses data dalam bentuk byte mentah.

---

## 🚀 Cara Menjalankan
```bash
go run main.go