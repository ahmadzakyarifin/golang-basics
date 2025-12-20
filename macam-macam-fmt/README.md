# 🏛️ Analisis Mendalam: Standar I/O dan Arsitektur Error di Go

Repositori ini mendokumentasikan penggunaan paket `fmt` dan `errors` dalam skenario Backend. Fokus utama adalah pada manajemen output stream (Stdout vs Stderr), pemformatan data presisi, dan hierarki pelacakan error.

---

## 🔍 1. Bedah Keluarga fmt (Print, S-Variant, & F-Variant)

Go menyediakan berbagai cara untuk menangani output berdasarkan tujuan datanya di sistem operasi.

### **A. Dasar Output: Print & Println**
* **`fmt.Print`**: Mencetak teks mentah. Tidak menambah spasi otomatis antar argumen dan tidak menambah baris baru (`\n`) di akhir.
* **`fmt.Println`**: Standar untuk *debugging* cepat. Otomatis menambah spasi antar argumen dan baris baru di akhir.

### **B. Kendali Format Penuh: fmt.Printf**
Menggunakan **Verb** (simbol `%`) untuk memetakan variabel ke dalam format teks tertentu. Berikut adalah daftar verb yang paling umum digunakan:

| Verb | Deskripsi Teknikal | Contoh Penggunaan |
| :--- | :--- | :--- |
| **`%v`** | **Value**: Format default. Go menebak format terbaik. | Data umum / Struct |
| **`%+v`** | **Struct+**: Menampilkan isi struct beserta **nama field-nya**. | Logging Object |
| **`%T`** | **Type**: Menampilkan **tipe data** dari variabel tersebut. | Debugging Tipe |
| **`%d`** | **Digit**: Format untuk angka bulat (integer). | ID, Jumlah, Umur |
| **`%s`** | **String**: Format untuk teks. | Nama, Pesan |
| **`%f`** | **Float**: Format untuk angka desimal. | Harga, Koordinat |
| **`%.2f`**| **Precision**: Membatasi 2 angka di belakang koma. | Mata Uang / Rupiah |
| **`%-15s`**| **Padding**: Menambah spasi (15 karakter) rata kiri. | Perapihan Tabel CLI |
| **`%p`** | **Pointer**: Menampilkan alamat memori variabel. | Analisis Memory |

---

## 🚀 2. Varian Output: Sprintf vs Fprintf

### **fmt.Sprintf (S-Variant)**
* **Sintaks:** `pesan := fmt.Sprintf(...)`
* **Analogi:** Seperti "menulis surat tapi belum dikirim". Pesannya disimpan dulu di sebuah variabel (**String**).
* **Kasus:** Digunakan saat kita ingin menyiapkan pesan log atau notifikasi untuk disimpan ke database atau dikirim melalui API.

### **fmt.Fprintf (F-Variant)**
* **Sintaks:** `fmt.Fprintf(os.Stderr, ...)`
* **Analogi:** Seperti "menentukan kotak pos tujuan".
    * **`os.Stdout`**: Jalur normal (Layar terminal).
    * **`os.Stderr`**: Jalur khusus error. Sangat penting di Backend agar log kesalahan bisa dipisahkan dari output data normal oleh sistem monitor (seperti Docker log).



---

## 🛡️ 3. Manajemen Error & Investigasi dengan `errors.Is`

Di Go, error adalah sebuah nilai (object) yang bisa diinvestigasi secara hierarkis, bukan sekadar pesan teks.

### **A. errors.New vs fmt.Errorf**
* **`errors.New`**: Membuat error statis (**Sentinel Error**). Cocok untuk error konstan seperti `ErrNotFound`.
* **`fmt.Errorf`**: Membuat error dinamis dengan tambahan konteks variabel.
* **Wrapping (`%w`)**: Menggunakan verb `%w` untuk membungkus (*wrap*) error lama di dalam error baru. Ini menjaga "rantai kejadian" agar tidak putus.

### **B. Investigasi dengan `errors.Is`**
Fungsi ini membongkar seluruh lapisan pembungkus (*unwrapping*) untuk mencari akar masalah asli (*root cause*).
* **Analogi:** Membuka kotak kado berlapis-lapis untuk melihat apakah hadiah di dalamnya sesuai dengan yang dicari.
* **Keunggulan:** Meskipun pesan error sudah diubah menjadi detail (misal: "Gagal ambil data: ID 505: not found"), `errors.Is` tetap bisa mendeteksi bahwa ini adalah tipe `ErrNotFound`.



---

## 📊 Tabel Perbandingan Final

| Fungsi / Sintaks | Input Utama | Tujuan / Output | Fitur Unggulan |
| :--- | :--- | :--- | :--- |
| **Println** | `interface{}` | Terminal (Stdout) | Cepat & baris baru otomatis. |
| **Printf** | `format, verbs` | Terminal (Stdout) | Formatting data sangat presisi. |
| **Sprintf** | `format, verbs` | `string` (Variabel) | Menyiapkan teks (tidak langsung print). |
| **Fprintf** | `io.Writer` | Stream (Stderr/Stdout) | Pemisahan jalur data vs error. |
| **errors.New** | `string` | `error` object | Definisi error tetap (Sentinel). |
| **fmt.Errorf** | `format, %w` | `error` object | Menambah konteks & Wrapping. |
| **errors.Is** | `err, target` | `bool` | Melacak akar masalah (Root Cause). |

---

## 💡 Insight Tambahan: Best Practice
Selalu gunakan `fmt.Fprintf(os.Stderr, ...)` untuk log kesalahan. Gunakan kombinasi `%w` dan `errors.Is` agar kode kamu tetap fleksibel; kita bisa memberikan pesan error yang manusiawi ke user, namun tetap memiliki data teknis yang akurat di backend.