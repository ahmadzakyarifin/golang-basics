package main

import "fmt"


// 1. INTERFACE SEBAGAI KONTRAK (Kasus: Service & Repo)
// "Pokoknya siapa pun yang mau jadi Repository, harus punya fungsi Simpan"
// =============================================================
type Repository interface {
	Simpan(data string) // Kontrak: Harus ada fungsi Simpan
}

// Implementasi 1: Database MySQL
type MySQL struct{}
func (m MySQL) Simpan(data string) {
	fmt.Println(">>> Berhasil menyimpan '" + data + "' ke MySQL")
}

// Implementasi 2: Database MongoDB
type MongoDB struct{}
func (mo MongoDB) Simpan(data string) {
	fmt.Println(">>> Berhasil menyimpan '" + data + "' ke MongoDB")
}

// Service: Dia hanya tahu 'Repository'. Gak peduli mau MySQL atau Mongo.
func UserService(repo Repository, namaUser string) {
	fmt.Println("Service: Menerima request untuk user:", namaUser)
	repo.Simpan(namaUser) // Memanggil fungsi sesuai kontrak
}

func main() {
	fmt.Println("=== 🏛️ CONTOH BACKEND (SERVICE & REPO) ===")
	
	db1 := MySQL{}
	db2 := MongoDB{}

	// Kita bisa ganti-ganti database dengan mudah (Decoupling)
	UserService(db1, "Zaky") 
	UserService(db2, "Zaky")

	// =============================================================
	// 2. INTERFACE KOSONG (Kasus: Wadah Data / any)
	// "Bisa nampung apa aja, tapi harus dibongkar kalau mau dipake"
	// =============================================================
	fmt.Println("\n=== 📦 CONTOH DASAR (ANY / INTERFACE KOSONG) ===")

	var wadah any // any bisa menampung tipe apa pun

	wadah = "Teks String"
	fmt.Println("Isi Wadah:", wadah)

	wadah = 100 
	fmt.Println("Isi Wadah:", wadah)

	// Cara bongkar (Type Assertion)
	// Kita cek: "Apakah benar isinya int?"
	angka, ok := wadah.(int)
	if ok {
		fmt.Println("Berhasil bongkar angka! Hasil + 10 =", angka+10)
	}
}