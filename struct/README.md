# 🏗️ Go Structs: Model Data & OOP Style

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