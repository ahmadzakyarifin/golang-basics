# 🏗️ Go Structs: Model Data & OOP Style

# Apa itu Struct di Go?

> **Singkatnya:** Struct adalah cara kita membuat "Tipe Data Kustom" sendiri yang membungkus beberapa variabel berbeda tipe menjadi satu kesatuan yang rapi.

---

## 🧐 Analogi Sederhana: "Formulir KTP"

Bayangkan kamu ingin menyimpan data penduduk. Sebuah KTP memiliki banyak jenis data yang berbeda:
1.  **Nama** (Teks/String)
2.  **Umur** (Angka/Integer)
3.  **Sudah Menikah?** (Benar/Salah / Boolean)

Tanpa struct, data ini akan berserakan. Dengan struct, kita "menjahit" data-data ini menjadi satu objek bernama `Penduduk`.

---

## ❌ Masalah: Tanpa Struct (Berantakan)

Jika kita ingin menyimpan data user tanpa struct, kita harus membuat variabel terpisah untuk setiap atribut. Ini akan **sangat ribet** jika usernya ada banyak.

```go
// Data User 1
var namaUser1 string = "Budi"
var umurUser1 int = 25
var emailUser1 string = "budi@gmail.com"

// Data User 2
var namaUser2 string = "Siti"
var umurUser2 int = 22
// ... Bayangkan kalau ada 1000 user, kodingan akan sangat kacau!

Materi ini menjelaskan bagaimana Go mengelola kumpulan data melalui **Struct** dan memberikan "perilaku" pada data tersebut melalui **Method**.

---

## 1. Aturan Penamaan (Access Modifier)
Di Go, hak akses tidak ditentukan oleh kata kunci `public/private`, melainkan oleh huruf pertama:
* **Huruf Besar (PascalCase):** `Person`, `Nama`. Bersifat **Exported** (Bisa diakses oleh package lain/Public).
* **Huruf Kecil (camelCase):** `person`, `umur`. Bersifat **Unexported** (Hanya bisa diakses di package yang sama/Private).

---

## 2. Inisialisasi Struct
Terdapat beberapa cara mendefinisikan struct:
1. **Named Fields:** Memberikan nilai langsung ke nama field-nya (Paling aman).
2. **Positional Fields:** Mengisi nilai sesuai urutan field (Kurang aman jika struktur berubah).
3. **Keyword `new`:** Membuat struct kosong dan mengembalikan **alamat memorinya** (Pointer).
4. **Anonymous Struct:** Membuat struct tanpa nama tipe untuk keperluan sementara.

---

## 3. Manipulasi Data: Value vs Pointer Receiver

Dalam "OOP" ala Go, kita menggunakan Receiver Function untuk membuat Method:

### A. Value Receiver `(p Person)`
* Go membuat **fotokopi** dari struct.
* Perubahan di dalam method **tidak akan merubah** data asli.
* Cocok untuk fungsi yang hanya menampilkan data (*Getter*).

### B. Pointer Receiver `(p *Person)`
* Go menggunakan **alamat memori asli**.
* Perubahan di dalam method **akan merubah** data asli secara permanen.
* Sangat efisien untuk struct berukuran besar karena tidak ada duplikasi data.

---

## 💡 Engineering Note
"Gunakan **Pointer Receiver** secara konsisten jika mayoritas method dalam struct bertujuan untuk mengubah nilai. Ini menjaga performa aplikasi tetap ringan karena menghindari alokasi memori berlebih untuk penyalinan data."

---

🌏 Implementasi Nyata: Database & API
Di dunia kerja (Backend Development), Struct sangat vital karena fungsinya sebagai Representasi Data.

1. Struct sebagai Cermin Database 🗄️
Bayangkan Database (SQL) seperti Excel. Struct adalah definisi kolomnya.

Tabel Database: Punya kolom id, username, created_at.

Struct Go: Harus dibuat sama persis field-nya agar data bisa masuk.

2. Struct sebagai Cetakan JSON (API) 🌐
Saat membuat Web API (misalnya dengan http.NewServeMux), data dikirim dalam format JSON. Kita menggunakan "Struct Tags" (tulisan di dalam tanda kutip miring/backtick) untuk mencocokkan data.