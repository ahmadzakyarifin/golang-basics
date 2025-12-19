# 🚀 Go Data Structures: Array, Slice, & Map

Materi ini memberikan panduan mendalam tentang bagaimana Go mengelola kumpulan data melalui Array, Slice, dan Map, serta strategi mengoptimalkan penggunaan memori dan CPU menggunakan `cap` (capacity).

---

## 1. Array (Fixed Length)
Array adalah struktur data statis. Sekali dideklarasikan, ukurannya tidak dapat diubah (immutable size).
* **Alokasi Memori:** Dilakukan sekali di awal sesuai jumlah elemen yang ditentukan.
* **Karakteristik:** Sangat cepat, namun tidak fleksibel.
* **Ketentuan `cap`:** Nilai `len` (panjang) dan `cap` (kapasitas) akan selalu **SAMA**.

---

## 2. Slice (Dynamic Length)
Slice adalah struktur data yang merujuk pada sebuah **Underlying Array** di memori.

### Memahami `len` vs `cap`
* **`len` (Length):** Jumlah elemen aktif yang bisa diakses saat ini.
* **`cap` (Capacity):** Ukuran total *underlying array* sebelum Go terpaksa melakukan reallokasi.

### Strategi Pertumbuhan Memori ("Pindah Rumah")
Saat kita melakukan `append` dan `len` melampaui `cap`, Go akan melakukan reallokasi otomatis:
1. **Buat Array Baru:** Go mencari memori baru (biasanya 2x lipat dari kapasitas saat ini).
2. **Copy Data:** CPU bekerja menyalin seluruh data dari array lama ke array baru.
3. **Pembersihan:** Array lama dihapus oleh *Garbage Collector*.

### Kenapa `make` Jauh Lebih Efisien bagi CPU?
Menggunakan `make([]type, len, cap)` memberikan kontrol penuh yang signifikan terhadap performa:

* **Memotong Tahapan Alokasi:** Jika menggunakan **Literal** yang mulai dari 1, Go harus merangkak melalui banyak tahap: 1 → 2 → 4 → 8 → 16. Setiap tanda panah adalah proses "pindah rumah" yang melelahkan bagi CPU.
* **Lompatan Kapasitas Lebih Besar:** Dengan `make`, kita bisa langsung mulai dari kapasitas besar (misal: 10). Saat meluap, Go akan langsung melompat ke 20. 
* **Efisiensi CPU:** Karena kelipatannya dimulai dari angka yang lebih besar, **jumlah total array baru yang dibuat menjadi jauh lebih sedikit**. CPU tidak perlu sering-sering melakukan proses alokasi dan penyalinan data.

---

## 3. Map (Key-Value Pair)
Map digunakan untuk menyimpan data dengan pasangan kunci unik dan nilai.

* **Inisialisasi via `make`:** `map1 := make(map[string]int)`. Sangat disarankan untuk menyiapkan struktur *hash table* agar siap diisi tanpa sering melakukan penyesuaian memori di tengah jalan.
* **Karakteristik:** Sangat efisien untuk pencarian data berdasarkan kata kunci.

---

## 4. Iterasi (Looping)
Cara paling aman untuk menampilkan isi data adalah menggunakan `for range`.

```go
// Pada Array/Slice
for index, item := range slice {
    fmt.Println(index, item)
}

// Pada Map
for key, value := range myMap {
    fmt.Println(key, value)
}