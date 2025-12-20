# 🔍 Mastering RegEx di Go

Dokumentasi ini merangkum teknik fundamental penggunaan Regular Expression di Golang menggunakan paket `regexp`. Materi ini mencakup validasi, ekstraksi, dan manipulasi teks secara efisien.

---

## 🏗️ 5 Pilar Wajib RegEx

Berdasarkan implementasi praktis, berikut adalah poin-poin kunci yang harus dikuasai:

### **1. Karakter Dasar & Sets**
* 📝 **`\d`**: Digunakan untuk menangkap angka (0-9).
* 📝 **`\w`**: Karakter alfanumerik (huruf, angka, `_`).
* 📝 **`[a-z]`**: Range spesifik untuk karakter tertentu.

### **2. Quantifiers (Jumlah)**
* 🔢 **`+`**: Minimal muncul 1 kali atau lebih.
* 🔢 **`{n,m}`**: Rentang jumlah karakter (misal: `{3,15}` berarti minimal 3 maksimal 15).
* 🔢 **`?`**: Karakter bersifat opsional.

### **3. Anchors (Ketajaman Posisi)**
* 📍 **`^`**: Mengunci pola agar wajib dimulai dari awal baris.
* 📍 **`$`**: Mengunci pola agar wajib berakhir di ujung baris.
* 💡 **Penting:** Tanpa anchor, validasi input bisa "bocor" karena RegEx akan mencari kecocokan di tengah teks sembarang.

### **4. Go Methods (regexp Package)**
* 🛠️ **`MustCompile`**: Mengunci pola menjadi objek RegEx. Sangat disarankan diletakkan di luar perulangan untuk menjaga performa CPU.
* 🛠️ **`MatchString`**: Standar untuk validasi input (mengembalikan boolean).
* 🛠️ **`FindAllString`**: Digunakan untuk mengambil data dari teks panjang (ekstraksi).
* 🛠️ **`ReplaceAllString`**: Senjata utama untuk fitur sensor atau pembersihan data string.



### **5. Case Insensitive Flag**
* ⚡ **`(?i)`**: Prefiks khusus agar pola tidak membedakan huruf besar dan kecil (misal: `(?i)go` akan cocok dengan "GO", "Go", atau "go").

---

## 💡 Engineering Insights
* ⚙️ **Raw Strings:** Selalu tulis pola menggunakan backtick (\`) agar tidak perlu melakukan *escaping* pada karakter backslash.
* ⚙️ **Performance:** Jangan memanggil `MustCompile` di dalam fungsi yang dijalankan berulang kali. Lakukan inisialisasi di tingkat global atau sekali saja di awal.



---

## 🚀 Cara Menjalankan
Pastikan file `main.go` sudah siap, lalu jalankan:
```bash
go run main.go