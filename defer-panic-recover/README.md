# ⚠️ Error Management: Defer, Panic, & Recover

Dokumentasi ini menjelaskan mekanisme penanganan kondisi darurat dan pengelolaan sumber daya (*resource cleanup*) di Golang.

---

## ⏳ 1. Defer (Cleanup)
`defer` digunakan untuk memastikan sebuah perintah dijalankan tepat di akhir fungsi, sebelum fungsi tersebut keluar (*return*).

* 📝 **Aturan LIFO (Last-In-First-Out):** Go menjalankan `defer` dengan urutan terbalik dari penulisan. Fungsi yang paling akhir ditulis akan dijalankan paling pertama.
* ⚡ **Urutan Eksekusi:** Jika menulis `defer A` lalu `defer B`, maka urutannya adalah **B dulu, baru A**.
* 💡 **Analogi:** Seperti melepas pakaian; kita melepas Jaket (terakhir dipakai) sebelum melepas Kemeja.

---

## 🚨 2. Panic (Fatal Error)
`panic` menghentikan alur eksekusi normal program saat terjadi kondisi yang tidak mungkin dipulihkan.

* 📝 **Efek Berhenti:** Baris kode di bawah `panic` **tidak akan dieksekusi**.
* ⚡ **Defer Tetap Jalan:** Meskipun alur berhenti, Go menjamin semua fungsi `defer` yang sudah terdaftar tetap akan dijalankan.
* 💡 **Kapan Digunakan:** Hanya untuk error kritis yang membuat aplikasi tidak bisa berjalan sama sekali.

---

## 🛡️ 3. Recover (Rescue)
`recover` adalah fungsi untuk menangkap data dari `panic` agar aplikasi tidak berhenti secara paksa (*crash*).

* 📝 **Wajib di dalam Defer:** `recover` hanya bisa bekerja jika diletakkan di dalam fungsi yang dipanggil lewat `defer`.
* ⚡ **Pemulihan Alur:** Jika `panic` tertangkap, program tidak mati, namun akan langsung lompat keluar fungsi tersebut dan melanjutkan alur di fungsi pemanggilnya.



---

## 💡 Engineering Insights
* 🛠️ **Kenapa Tidak Lanjut Eksekusi?** Karena `panic` menandakan status fungsi sudah tidak valid atau berbahaya untuk dilanjutkan.
* 🛑 **Error vs Panic:** Gunakan tipe data `error` untuk masalah biasa. Gunakan `panic` hanya jika sistem benar-benar tidak bisa bekerja.
* 🗑️ **Resource Safety:** Selalu tulis `defer close()` segera setelah berhasil membuka sebuah *resource* agar tidak terjadi kebocoran memori.

---

## 🚀 Cara Menjalankan
Gunakan perintah berikut di terminal:
```bash
go run main.go