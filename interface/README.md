# 🏛️ Memahami Interface di Go: Kontrak & Wadah

Interface bukan tentang "apa itu benda tersebut", tapi tentang **"apa yang bisa dilakukan benda tersebut"**.

---

## 🔍 1. Interface di Backend (Service & Repo)
Dalam pembuatan API/Backend, Interface digunakan untuk membuat **"Kontrak"** agar kode kita tidak kaku.

* **Repository:** Tempat urusan database (MySQL, PostgreSQL, dll).
* **Service:** Tempat logika aplikasi.
* **Manfaat:** Jika kita ingin pindah dari MySQL ke MongoDB, kita tidak perlu mengubah kode di bagian **Service**. Cukup buat implementasi baru yang mengikuti "Kontrak" yang sama.



---

## 📦 2. Interface Kosong (`any`)
Interface kosong (`any`) adalah tipe data yang paling sakti karena bisa menampung **data apa pun**.

* **Kegunaan:** Dipakai saat kita tidak tahu tipe data apa yang akan masuk (misal: isi Slice yang campur-campur antara teks dan angka).
* **Aturan Main:** Karena dia "bungkusan universal", datanya harus dibongkar dulu menggunakan **Type Assertion** (`v.(int)`) sebelum bisa dioperasikan secara matematika atau logika tipe spesifik.



---

## 🛠️ Ringkasan Poin Penting
1. **Implicit:** Di Go, kita tidak perlu nulis `implements`. Cukup buat fungsi yang namanya sama, maka otomatis dianggap memenuhi interface.
2. **Decoupling:** Memisahkan urusan logika (Service) dengan urusan database (Repo).
3. **Flexible:** Menggunakan `any` untuk membuat fungsi yang bisa menerima input macam-macam.