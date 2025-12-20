# 🏦 Simulasi Concurrency: Analogi Kasir Bank & Manajemen Tugas

Project ini mendemonstrasikan bagaimana Golang mengelola banyak proses secara bersamaan (**Goroutine**) dengan tetap menjaga integritas data menggunakan **Mutex** dan sinkronisasi waktu dengan **WaitGroup**.

---

## 🎭 Analogi Besar: Operasi Bank

Bayangkan sebuah Bank yang memiliki satu **Buku Tabungan Raksasa** (variabel `saldo`) yang diletakkan di meja tengah.

### 1. `wg.Add(1)` — Manajer Mencatat Kehadiran
Sebelum Kasir mulai bekerja, **Manajer Bank** (fungsi `main`) memegang alat hitung (*counter*). Setiap kali ia memanggil kasir (`go kasir`), ia memencet tombol: **+1**.
* **Logika**: Manajer mencatat **sebelum** kasir berangkat. Jika catatan total adalah 5, maka bank tidak boleh tutup sampai angka itu kembali ke 0.

### 2. `go kasir` — Balapan di Belakang Layar
Saat perintah `go` dipanggil, kasir-kasir tidak berjalan berurutan. Mereka **berlari bersamaan** menuju Buku Tabungan.
* **Non-Deterministic**: Di belakang layar (CPU), sistem membagi tenaga secara dinamis. Kasir 3 mungkin berlari lebih cepat dari Kasir 1. Itulah alasan mengapa tulisan "Selesai" di terminal muncul tidak berurutan.



### 3. `mutex.Lock()` & `Unlock()` — Pena Ajaib
Kelima kasir sampai di depan Buku Tabungan secara bersamaan. Tanpa keamanan, mereka akan berebutan menulis dan merusak data (**Race Condition**).
* **`Lock()`**: Kasir pertama yang sampai akan merebut **Satu-satunya Pena Ajaib**. Begitu pena dipegang, kasir lain **wajib mematung** (berhenti sejenak) meskipun mereka sudah siap menulis.
* **`Unlock()`**: Setelah selesai mencatat, kasir meletakkan pena kembali di meja sehingga kasir lain bisa kembali berebut giliran.

### 4. `defer wg.Done()` — Laporan Tugas
Setelah seorang kasir menyelesaikan transaksinya, ia mendatangi Manajer dan melapor. Manajer memencet tombol pengurang pada alat hitungnya: **-1**.

### 5. `wg.Wait()` — Gembok Pintu Bank
Fungsi `main` (Manajer) sebenarnya sangat cepat selesai. Namun, perintah `wg.Wait()` memaksanya berdiri di pintu keluar.
* **Logika**: Selama angka di alat hitung belum **0**, Manajer tetap berjaga. Jika tidak ada `Wait()`, Bank akan tutup (program mati) saat kasir-kasir di dalam sebenarnya masih sibuk bekerja.

---

## 💡 Engineering Insights (Behind the Scenes)

### ⚠️ Race Condition vs. Independent Tasks
Penting untuk memahami kapan kita membutuhkan bantuan **Mutex**:
* **Gunakan Mutex**: Jika banyak Goroutine mengakses **satu aset yang sama** (seperti variabel `saldo`). Mutex memastikan proses "Baca-Ubah-Tulis" dilakukan secara **Atomic** (satu orang sampai tuntas) agar tidak ada hitungan yang hilang.
* **Cukup WaitGroup**: Jika Goroutine menjalankan **tugas yang berbeda-beda** (misal: satu Masak, satu Nyapu, satu Cuci Piring). Karena mereka tidak berebutan aset yang sama, kita hanya perlu WaitGroup untuk tahu kapan semua tugas berbeda itu selesai.



### ⚙️ Go Scheduler (M:P:N Model)
Go mampu menjalankan ribuan Goroutine meski CPU hanya memiliki sedikit Core (inti).
* **Work Stealing**: Jika salah satu Core CPU menganggur, Go Scheduler akan otomatis "mencuri" tugas dari Core lain agar beban kerja merata.
* **Concurrent, Not Sequential**: Urutan fungsi di kode tidak menjamin urutan eksekusi. CPU yang menentukan siapa yang jalan duluan berdasarkan ketersediaan *resource*.



### 🛡️ Race Detector (Fitur Detektif)
Go menyediakan alat untuk mendeteksi apakah kode kita aman dari tabrakan data. Selalu gunakan flag `-race` saat tahap testing untuk memastikan tidak ada akses memori yang bocor atau bertabrakan:
```bash
go run -race main.go