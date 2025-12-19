package main

import "fmt"

// 1. Ketentuan Penamaan (Visibility)
// Nama Besar (Person) = Public/Exported
// Nama Kecil (umur) = Private/Unexported
type Person struct {
	Nama   string
	umur   int   
	Gender string
}

// --- METODE RECEIVER (Gaya OOP di Go) ---

// Value Receiver: Hanya meminjam salinan data (Read-Only)
func (p Person) TampilkanProfil() {
	fmt.Printf("[Value Receiver] Nama: %s, Umur: %d, Gender: %s\n", p.Nama, p.umur, p.Gender)
}

// Pointer Receiver: Mengakses data asli di memori (Update/Setter)
func (p *Person) SetUmur(umurBaru int) {
	p.umur = umurBaru
}

// --- FUNGSI BIASA ---

func ubahNamaValue(p Person, namaBaru string) {
	p.Nama = namaBaru
	fmt.Println("   -> Inside Func (Value): Nama berubah jadi", p.Nama)
}

func ubahNamaPointer(p *Person, namaBaru string) {
	p.Nama = namaBaru
}

func main() {
	// --- BERBAGAI CARA MEMBUAT STRUCT ---

	// A. Literal dengan Nama Field (Rekomendasi Utama)
	user1 := Person{
		Nama:   "Zaky",
		umur:   19,
		Gender: "Laki-laki",
	}

	// B. Literal tanpa Nama Field (Harus urut)
	user2 := Person{"Kayla", 18, "Perempuan"}

	// C. Menggunakan 'new' (Menghasilkan Pointer)
	user3 := new(Person)
	user3.Nama = "Arifin"

	// D. Anonymous Struct (Sekali pakai)
	user4 := struct {
		Role string
	}{
		Role: "Developer",
	}

	fmt.Println("=== 1. EKSPERIMEN PASS BY VALUE ===")
	ubahNamaValue(user1, "Zaky Palsu")
	fmt.Println("   Hasil di Main:", user1.Nama)
	 // Tetap Zaky

	fmt.Println("\n=== 2. EKSPERIMEN PASS BY REFERENCE ===")
	ubahNamaPointer(&user1, "Zaky Arifin")
	fmt.Println("   Hasil di Main:", user1.Nama) 
	// Berubah jadi Zaky Arifin

	fmt.Println("\n=== 3. EKSPERIMEN METHOD RECEIVER ===")
	user1.SetUmur(20)      
	// Mengubah data asli via Pointer Receiver
	user1.TampilkanProfil() 
	// Menampilkan data via Value Receiver

	fmt.Println("\n=== DATA LAINNYA ===")
	fmt.Println("User 2:", user2)
	fmt.Println("User 3 (Pointer):", *user3)
	fmt.Println("User 4 (Anonymous):", user4)
}