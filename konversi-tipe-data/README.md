# 🔄 Konversi Tipe Data di Golang (Casting & strconv)

Repositori ini mendokumentasikan berbagai cara melakukan konversi antar tipe data di Go, mulai dari tipe numerik dasar hingga penggunaan paket `strconv` untuk konversi string.

## 🚀 Poin Penting Eksperimen

### 1. Numeric Casting (Konversi Angka)
Di Go, konversi antar tipe angka harus dilakukan secara **eksplisit**. Kamu tidak bisa langsung menjumlahkan `int` dengan `float64` tanpa melakukan casting terlebih dahulu.
* **Sintaks:** `TipeDataBaru(variabel)`
* **Pembulatan:** Saat melakukan casting dari `float` ke `int`, Go akan membuang angka di belakang koma (truncation), bukan membulatkannya ke atas/bawah secara matematis.

### 2. Bahaya Overflow & Integer Wrap-around
Ini adalah bagian krusial dalam koding sistem. Setiap tipe data memiliki batas kapasitas.
* **Kasus:** Mengonversi `int` (300) ke `uint8` (kapasitas maksimal 255).
* **Hasil:** Nilai akan "berputar" kembali ke angka terkecil (wrap-around). 
* **Rumus:** `300 - 256 = 44`. Maka hasilnya adalah **44**.



### 3. Konversi String dengan `strconv`
Karena `string` bukan tipe numerik, kita tidak bisa menggunakan casting biasa seperti `int(string)`. Go menyediakan paket bawaan `strconv`.
* **Itoa (Integer to ASCII):** Mengubah angka menjadi teks.
* **Atoi (ASCII to Integer):** Mengubah teks menjadi angka. Fungsi ini mengembalikan dua nilai: **hasil** dan **error**. Sangat penting untuk mengecek error jika teks yang dimasukkan bukan angka valid (misal: "abc").

---

## 🛠️ Ringkasan Metode Konversi

| Skenario | Cara | Catatan |
| :--- | :--- | :--- |
| **int ↔ float** | `float64(i)` atau `int(f)` | Hati-hati kehilangan presisi desimal. |
| **int ↔ string** | `strconv.Itoa()` / `strconv.Atoi()` | Membutuhkan penanganan error untuk Atoi. |
| **Beda Ukuran Bit** | `int32(i16)` | Aman jika dari kecil ke besar. |
| **Negative ↔ Uint** | `uint(negatif)` | **Sangat Berbahaya!** Menghasilkan angka sangat besar karena bit pertama (sign bit) terbaca sebagai nilai. |

---

## 💡 Catatan Mekanisme Memori
Setiap tipe data numerik di Go memiliki ukuran bit yang berbeda (8, 16, 32, 64 bit). Saat kamu melakukan konversi ke tipe yang lebih kecil sementara nilainya melebihi kapasitas tipe tersebut, maka terjadi **Data Loss** atau **Overflow**, di mana nilai akan kembali ke batas minimum tipe data tersebut.



---