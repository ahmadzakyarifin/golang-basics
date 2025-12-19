# 🏛️ Go Memory Mastery: Analogi Google Doc & Pointer Koleksi

Repositori ini mendokumentasikan eksperimen mendalam mengenai cara kerja memori di Golang. Fokus utama proyek ini adalah membedah perilaku **Pass-by-Value** (fotokopi data) dan **Pass-by-Reference** (berbagi alamat) pada tipe data kompleks.

---

## 🧠 Analogi "Google Doc": Memahami Slice & Pointer

Banyak developer bingung: *"Katanya Slice itu Reference, tapi kenapa kalau di-append di dalam fungsi, datanya nggak nambah di fungsi main?"*. Berikut adalah visualisasi logikanya:

### 1. Ubah Isi: "Kolaborasi Dokumen"
**Ibarat:** Slice adalah sebuah **Link Google Doc**. Saat Anda mengirim Slice ke fungsi, Anda memberikan Link tersebut.
* **Aksi:** Jika fungsi mengubah isi tulisan di dalam dokumen tersebut (`slc[0].Umur = 23`), Anda yang memegang link aslinya akan melihat perubahan tersebut.
* **Kesimpulan:** Mengubah isi index tidak butuh pointer ke slice (`*[]`).



### 2. Append Data: "Ganti Dokumen Baru"
**Ibarat:** Dokumen Google Doc tersebut memiliki batas halaman (Kapasitas). Saat dilakukan `append` dan kapasitas penuh, Go otomatis **membuat dokumen baru** yang lebih besar.
* **Masalah:** Fungsi sekarang memegang **Link Baru**, sedangkan Anda di `main` masih memegang **Link Lama**. Inilah kenapa data tidak bertambah di `main`.
* **Solusi di Proyek Ini:** Kita menggunakan **Pointer ke Slice** (`*[]*Mahasiswa`) agar fungsi bisa langsung mengganti "Link" yang Anda pegang di `main` dengan link dokumen baru.



---

## 🧪 Analisis Kasus Spesifik (Deep Dive)

### A. Pointer ke Struct (`*Mahasiswa`)
Digunakan pada fungsi `ubahUmur`. Mengirim alamat memori memastikan kita tidak melakukan duplikasi data (fotokopi) struct yang berat, serta memastikan perubahan bersifat permanen.

### B. Array of Struct (`*[N]T`)
Array di Go adalah *Value Type*. Jika dikirim tanpa pointer, Go menyalin seluruh paket data array. Dengan `*[2]Mahasiswa`, kita hanya mengirim alamat (8 byte) namun mendapatkan akses penuh ke elemen asli.

### C. Map of Pointers (`map[string]*Mahasiswa`)
Map adalah tipe data referensi murni (Grup WhatsApp). Kita menggunakan pointer ke struct di dalam map agar:
1. Kita diizinkan mengubah field struct secara langsung (`m["x"].Nama = "Anton"`).
2. Efisiensi memori karena map hanya menyimpan alamat objek.



---

## 🛠️ Ringkasan Ketentuan Mutasi

| Tipe Data | Perilaku Default | Teknik Update | Status di Main |
| :--- | :--- | :--- | :--- |
| **Struct** | Value (Copy) | **Pointer** (`*T`) | ✅ Berubah |
| **Array** | Value (Copy) | **Pointer** (`*[N]T`) | ✅ Berubah |
| **Slice (Isi)** | Reference-like | **Langsung** | ✅ Berubah |
| **Slice (Add)** | Pass-by-Value | **Pointer / Return** | ✅ Berubah |
| **Map** | Reference | **Langsung** | ✅ Berubah |

---

## 🏆 Best Practice Produksi: Return vs Pointer

Meskipun dalam proyek ini kita mendemonstrasikan **Pointer ke Slice** (`*[]T`), di dunia kerja pola **Return** lebih disukai untuk urusan `append`.

**Kenapa Return lebih baik?**
1. **Idiomatic Go:** Mengikuti cara kerja fungsi internal `append(s, v)`.
2. **Readability:** Kode lebih eksplisit karena ada proses penugasan ulang (`s = tambah(s)`).
3. **Safety:** Menghindari kerumitan manipulasi alamat memori yang rawan *bug* bagi pemula.

**Gunakan Pointer ke Slice (`*[]T`) hanya jika:**
- Fungsi memiliki terlalu banyak nilai kembalian (*return values*).
- Diwajibkan oleh arsitektur *library* atau *interface* tertentu.

---