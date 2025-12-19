# 🛠️ Ragam Penulisan Fungsi (Function) di Golang

Dokumentasi ini mencakup berbagai cara penulisan fungsi di Go, mulai dari tingkat dasar hingga penggunaan tingkat lanjut seperti *High-Order Functions*.

## 📌 Jenis Fungsi yang Dipelajari

### 1. Dasar & Multiple Return
Go mewajibkan tipe data parameter ditulis secara eksplisit. Fitur **Multiple Return** sangat krusial karena Go tidak menggunakan `try-catch`, melainkan mengembalikan nilai bersamaan dengan `error`.



### 2. Named Return Values
Memungkinkan kita mendeklarasikan variabel return langsung di *header* fungsi. Hal ini membantu dokumentasi kode agar lebih mudah dibaca tanpa harus mencari variabel di dalam badan fungsi.

### 3. Variadic Function (`...`)
Fungsi yang fleksibel menerima jumlah argumen berapapun. Parameter variadic akan ditangkap sebagai sebuah **Slice** di dalam fungsi.

### 4. Function as First-Class Citizen
Di Go, fungsi adalah tipe data. Hal ini memungkinkan:
* **Anonymous Function:** Fungsi tanpa nama untuk logika sementara.
* **Function as Parameter:** Mengirim fungsi ke fungsi lain (pola *Callback*).
* **Function Value:** Menyimpan fungsi ke dalam variabel.



---

## 🛠️ Ringkasan Cepat

| Teknik | Kegunaan Utama |
| :--- | :--- |
| **Named Return** | Efisiensi kode & keterbacaan (Self-documenting). |
| **Variadic** | Menangani input dinamis (seperti `fmt.Println`). |
| **Anonymous** | Digunakan di dalam *Closure* atau variabel lokal. |
| **Function Param** | Membuat fungsi yang sangat fleksibel (*Reusable*). |

---

## 💡 Insight: Filosofi Go
Go menekankan pada kejelasan (*Explicitness*). Penggunaan fungsi di Go didesain agar aliran data (input dan output) terlihat sangat jelas, terutama dalam menangani kondisi error melalui nilai kembalian ganda.