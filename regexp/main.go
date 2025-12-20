package main

import (
	"fmt"
	"regexp"
)

// Inisialisasi pola di tingkat global (Best Practice)
// Memastikan compile hanya terjadi satu kali saat aplikasi startup
var (
	reUsername = regexp.MustCompile(`^\w{3,15}$`)
	reBarang   = regexp.MustCompile(`BR\d{3}`)
	reSensor   = regexp.MustCompile(`(?i)buruk`)
)

func main() {
	// Validasi Format (MatchString)
	// Berguna untuk pengecekan input user di form pendaftaran
	fmt.Println("==== 1. Validasi Pola Username ====")
	fmt.Println("ahmad_zaky :", reUsername.MatchString("ahmad_zaky")) // true
	fmt.Println("za         :", reUsername.MatchString("za"))         // false (terlalu pendek)

	// Ekstraksi Data (FindAllString)
	// Berguna untuk parsing data dari teks laporan atau log
	fmt.Println("\n==== 2. Ekstraksi Kode Barang ====")
	gudang := "Kode barang: BR001, BR002, BR003"
	hasilEkstrak := reBarang.FindAllString(gudang, -1)
	fmt.Println("Hasil Ekstrak:", hasilEkstrak)

	// Manipulasi Teks (ReplaceAllString)
	// Berguna untuk fitur moderasi konten otomatis
	fmt.Println("\n==== 3. Sensor Komentar ====")
	komentarNetizen := "Pelayanan di toko ini buruk sekali !!!"
	fmt.Println("Hasil Sensor :", reSensor.ReplaceAllString(komentarNetizen, "***"))
}