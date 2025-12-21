# 🏦 Go Concurrency Master: Mutex vs. Channels (The Bank Story)

Project ini mendemonstrasikan bagaimana Golang mengelola banyak tugas secara independen (**Concurrency**) melalui analogi operasional bank. Kita akan membedah perbedaan mendalam antara penggunaan **Mutex** (Shared Memory) dan **Channels** (Message Passing).

---

## 🏎️ Apa itu Concurrency?
**Concurrency** adalah kemampuan program untuk menangani banyak pekerjaan secara mandiri. Di Go, kita menggunakan **Goroutine**. Ibarat seorang kasir yang bisa melayani nasabah sambil menunggu mesin cetak selesai bekerja, CPU akan terus berganti tugas dengan sangat cepat sehingga semua pekerjaan tampak berjalan bersamaan.

---

## 🎭 Analogi Utama: Manajemen Saldo Bank

Bayangkan sebuah Bank dengan satu **Buku Tabungan Raksasa** (variabel `saldo`). Ada 5 Kasir (Goroutine) yang masing-masing harus mencatat 1.000 koin (total 5.000 transaksi).

### 1. Metode Mutex (Rebutan Pena Ajaib)
Semua Kasir berlari ke meja yang sama untuk menulis di satu buku yang sama.
* **Mekanisme**: Kasir yang sampai duluan mengambil **Pena Ajaib** dan mengunci buku (**Lock**). Kasir lain wajib berhenti total dan mengantre. Setelah selesai, pena diletakkan kembali (**Unlock**) agar kasir lain bisa berebut lagi.
* **Risiko**: Jika kasir lupa mengembalikan pena, bank akan macet selamanya (**Deadlock**).



### 2. Metode Channel (Pipa Teleportasi)
Kasir tidak pernah menyentuh Buku Besar secara langsung. Buku itu dijaga oleh satu **Petugas Khusus (Consumer)** di ruangan terpisah.
* **Mekanisme**: Kasir cukup melempar koin melalui **Pipa (Channel)**. Koin tersebut akan mengalir langsung ke tangan Petugas Khusus untuk dicatat.
* **Kelebihan**: Data saldo terisolasi. Tidak ada rebutan pena, yang ada hanyalah aliran data yang rapi dari pengirim ke penerima.

---

## 📦 Membedah Jenis Pipa (Channel) & Mekanisme Antrean

Bagaimana koin berpindah dari Kasir ke Petugas? Ada dua kondisi pipa:

### A. Pipa Tanpa Buffer (Unbuffered) - "Serah Terima Satu-Satu"
* **Mekanisme**: Pipa tidak memiliki ruang simpan. Pintu pipa hanya terbuka jika ada orang di **kedua ujungnya** di saat yang sama.
* **Proses**: Kasir harus memegang koin di mulut pipa dan menunggu petugas menangkapnya. Begitu satu koin diterima, barulah kasir berikutnya bisa menaruh koinnya.
* **Efek**: Proses ini kaku dan lambat karena terjadi "serah terima tangan ke tangan" untuk setiap koin. Kasir harus sering berhenti menunggu kesiapan petugas.



### B. Pipa Dengan Buffer (Buffered) - "Sistem Laci Borongan"
* **Mekanisme**: Pipa memiliki **Laci Penampung** dengan kuota tertentu (misal: 100).
* **Proses**: Kasir bisa menaruh koin secara "borongan" tanpa menunggu petugas siap. Misalnya, Kasir 1 langsung menaruh 50 koin, disusul Kasir 2 menaruh 30 koin. 
* **Kapan Menunggu?**: Kasir baru akan berhenti/menunggu jika **Laci Penampung sudah Penuh**. Begitu petugas mengambil koin dari ujung laci, barulah kasir bisa mengisi sisa koinnya lagi.
* **Efek**: Proses jauh lebih cepat karena kasir bisa mencicil banyak koin sekaligus ke dalam laci tanpa terhambat oleh kecepatan petugas, selama kapasitas laci masih tersedia.

## 🛡️ Aturan Emas Sinkronisasi (Urutan Eksekusi)

Agar sistem berjalan tanpa error (Panic/Deadlock), urutan pengakhiran harus dilakukan sebagai berikut:

1. **Wait for Producers (`wg.Wait`)**: Pastikan semua pengirim data (Kasir) sudah menyelesaikan tugasnya.
2. **Close Communication (`close(ch)`)**: Tutup pipa pengiriman. Ini memberi tahu penerima (Petugas) bahwa tidak ada lagi data yang akan masuk.
3. **Wait for Consumer (`<-finish`)**: Pastikan penerima telah selesai memproses semua data sisa yang ada di dalam antrean (Buffer).

## 🕵️ Mengapa Consumer Harus Berjalan di Latar Belakang (Goroutine)?

Consumer atau "Petugas Pencatat" wajib dijalankan sebagai Goroutine (`go`) karena alasan berikut:

1. **Non-Blocking Flow**: Consumer bertugas mengosongkan pipa. Jika Consumer tidak stand-by sejak awal, pipa (buffer) akan cepat penuh dan mengakibatkan para Kasir berhenti bekerja (Terhenti/Stuck).
2. **Dual Action**: Consumer bekerja secara konkuren; ia terus memantau pipa dan langsung melakukan pencatatan saldo sesaat setelah data diterima.
3. **Penyelamat Data**: Dengan bantuan sinyal `finish`, Consumer menjamin bahwa meskipun semua Kasir sudah pulang, ia akan tetap tinggal di kantor sampai koin terakhir di dalam laci selesai dicatat dengan benar.

---

## ⚙️ Engineering Insights: Di Balik Layar

### 🔄 Siapa yang Membawa Koin?
Di Go, Channel adalah **Pipa Teleportasi otomatis**. Hambatannya bukan di "jarak", tapi di **"kapasitas"**. Tanpa buffer, pengirim dan penerima harus sinkron satu-satu. Dengan buffer, pengirim diberikan kebebasan untuk "menimbun" hasil kerjanya selama tempat penampungan masih ada.

### ⚙️ Manajemen Otomatis (Go Scheduler)
Satu hal luar biasa dari Go adalah kita tidak perlu pusing mengatur urutan kerja di CPU. 
* **Efisiensi**: Go akan mengatur sendiri mana Goroutine yang harus jalan duluan dan mana yang harus "istirahat" sebentar saat pipa penuh.
* **Fokus pada Alur**: Sebagai pengembang, kita cukup fokus pada bagaimana data mengalir di dalam **Channel** dan bagaimana **Buffer** membantu mempercepat pekerjaan kasir. Sisanya, biarkan Go yang bekerja untuk kita.


### 🏁 Memahami Output Terminal: Mengapa Hasilnya "Acak"?

Saat menjalankan program, Anda mungkin melihat urutan laporan penyelesaian yang tidak berurutan, contohnya:
- Kasir 1: Selesai...
- Kasir 4: Selesai...
- Kasir 2: Selesai...

**Apa yang sebenarnya terjadi di balik layar?**

1. **Laporan Garis Finish**: Pesan "Selesai" di terminal hanya muncul ketika seorang Kasir telah menuntaskan seluruh 1.000 koinnya. Urutan yang tidak berurutan (misal: 1 lalu 4, bukan 1 lalu 2) adalah bukti otentik bahwa mereka **balapan di jalur yang sama secara bersamaan**.
2. **Siapa Cepat Dia Habis**: Setiap Kasir (Goroutine) berlomba memperebutkan akses ke CPU dan ruang kosong di pipa (Buffer). Kasir 4 bisa muncul lebih dulu karena ia mendapatkan "jatah waktu" yang lebih banyak dari sistem untuk mengirim koin, sehingga ia mencapai garis finish (menyelesaikan loop) lebih cepat daripada Kasir 2.
3. **Pipa yang Berdesakan**: Di dalam pipa, koin-koin dari Kasir 1 hingga 5 masuk secara **acak dan selang-seling**. Petugas (Consumer) pun mengambil koin tersebut secara acak siapa pun yang paling cepat menaruhnya di lubang pipa.
4. **Kecepatan Bukan Urutan**: Output ini membuktikan bahwa Concurrency di Go mengutamakan **Efisiensi** (menyelesaikan semua tugas secepat mungkin) daripada **Urutan** (menunggu Kasir 1 selesai baru Kasir 2 mulai).

> **Kesimpulan**: Terminal hanya melaporkan hasil akhir dari sebuah kompetisi lari. Meskipun laporannya muncul satu-satu, kenyataannya semua pelari (Kasir) sedang berada di lintasan secara bersamaan dan saling salip-menyalip di setiap meternya.
---


## 🚀 Kesimpulan Filosofis
**"Don't communicate by sharing memory, share memory by communicating."**
* **Mutex**: Proteksi manual pada aset yang diperebutkan (Shared Memory).
* **Channel**: Aliran data otomatis melalui pipa pesan (Message Passing).
* **Buffer**: Memberikan kapasitas "borongan" agar pekerja tidak sering berhenti menunggu, sehingga meningkatkan kecepatan secara drastis.

---

## 🛠️ Cara Menjalankan
```bash
# Jalankan program utama
go run main.go

# Cek tabrakan data (Race Condition)
go run -race main.go