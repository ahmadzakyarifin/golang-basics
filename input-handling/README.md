# 📥 Go Input Handling: Scan vs Scanln vs Bufio

Materi ini membahas teknik mengambil input dari terminal dan membedah tuntas fenomena **Input Buffer** yang sering menyebabkan fungsi input terlewati (skip) otomatis.

---

## 🔍 Perbandingan Metode & Karakteristik

### 1. `fmt.Scan()`
* **Sifat:** "Haus Input".
* **Perilaku:** Ia akan terus menunggu sampai semua variabel terisi. Meskipun user menekan Enter berkali-kali, ia tidak akan berhenti sebelum ada data yang masuk untuk mengisi variabel yang diminta.
* **Pemisah:** Menggunakan spasi sebagai pemisah antar input.
* **Kelemahan:** Meninggalkan karakter `\n` (Enter) tetap tertinggal di dalam Buffer memori.

### 2. `fmt.Scanln()` (Spesifik & Tegas)
* **Sifat:** "Sensitif terhadap Enter".
* **Perilaku:** Berhenti tepat saat user menekan Enter (`\n`), apa pun kondisinya. Jika variabel yang diminta ada 2 tapi user hanya mengisi 1 lalu menekan Enter, maka variabel kedua akan **kosong**.
* **Masalah Utama (The Ghost Newline):** Inilah penyebab utama input terlewati. Jika sebelumnya ada `fmt.Scan`, sisa `\n` di buffer akan langsung dibaca oleh `Scanln`. Akibatnya, `Scanln` mengira user sudah menekan Enter dan langsung selesai tanpa menunggu input baru.

### 3. `bufio.NewReader()`
* **Sifat:** "Pembaca Baris Penuh".
* **Keunggulan:** Paling stabil untuk input kalimat panjang yang mengandung spasi. Ia membaca segala sesuatu sampai menyentuh karakter `\n`.
* **Tips:** Hasilnya sering membawa karakter `\n` di ujung string, gunakan `strings.TrimSpace()` untuk membersihkannya.

---

## ⚠️ Visualisasi Masalah Buffer (The Skip Phenomenon)



**Kronologi Input yang Terlewati:**
1. User menjalankan `fmt.Scan(&nama)`.
2. User mengetik `Zaky` lalu menekan **Enter**.
3. `nama` mengambil `Zaky`, tetapi karakter **Enter (`\n`)** masih nyangkut di buffer.
4. Program menjalankan `fmt.Scanln(&alamat)`.
5. `Scanln` melihat buffer masih ada sisa `\n`, ia menganggap user sudah menekan Enter, lalu ia langsung berhenti.
6. Hasilnya: `alamat` menjadi kosong dan user tidak sempat mengetik apa pun.

---

## 🚀 Kesimpulan & Solusi
* **Gunakan `Scan`** jika hanya ingin mengambil kata per kata.
* **Gunakan `Scanln`** jika ingin input selesai tepat saat user menekan Enter, tapi pastikan buffer sudah bersih.
* **Gunakan `bufio`** untuk input paling aman (kalimat panjang).
* **Solusi Buffer:** Jika berpindah dari `Scan` ke `Scanln` atau `bufio`, selalu bersihkan sisa `\n` menggunakan `bufio.NewReader(os.Stdin).ReadString('\n')` atau `fmt.Scanln()` kosong.

---