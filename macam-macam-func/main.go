package main

import (
	"fmt"
	"strings"
)

// ==========================================
// 1. Fungsi Standar (Basic)
// ==========================================
func sapa(nama string) {
	fmt.Println("Halo,", nama)
}

// ==========================================
// 2. Multiple Return Values & Error Handling
// ==========================================
func bagi(a, b float64) (float64, error) {
	if b == 0 {
		return 0, fmt.Errorf("kesalahan: tidak bisa membagi dengan nol")
	}
	return a / b, nil
}

// ==========================================
// 3. Named Return Values
// ==========================================
func hitungPersegi(sisi int) (luas int, keliling int) {
	luas = sisi * sisi
	keliling = 4 * sisi
	return // otomatis mengembalikan variabel luas & keliling
}

// ==========================================
// 4. Variadic Function (Parameter Tak Terbatas)
// ==========================================
func jumlahkan(namaDaftar string, angka ...int) {
	total := 0
	for _, nilai := range angka {
		total += nilai
	}
	fmt.Printf("%s: %d\n", namaDaftar, total)
}

// ==========================================
// 5. Function as Parameter (Callback)
// ==========================================
func filterNama(nama string, filter func(string) string) {
	namaFiltered := filter(nama)
	fmt.Println("Nama asli:", nama, "| Setelah filter:", namaFiltered)
}

func main() {
	fmt.Println("=== EKSPERIMEN RAGAM FUNGSI GO ===")

	// --- Test Basic ---
	sapa("Zaky")

	// --- Test Multiple Return ---
	hasil, err := bagi(10, 2)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Hasil bagi:", hasil)
	}

	// --- Test Named Return ---
	l, k := hitungPersegi(5)
	fmt.Printf("Persegi -> Luas: %d, Keliling: %d\n", l, k)

	// --- Test Variadic ---
	jumlahkan("Total Belanja", 1000, 2500, 5000, 1500)

	// --- Test Anonymous Function (Fungsi Tanpa Nama) ---
	blacklist := func(name string) bool {
		return name == "admin" || name == "root"
	}
	fmt.Println("Apakah 'admin' dilarang?", blacklist("admin"))

	// --- Test Function as Parameter ---
	sensor := func(name string) string {
		if name == "kasar" {
			return "*****"
		}
		return strings.ToUpper(name)
	}
	filterNama("zaky", sensor)
	filterNama("kasar", sensor)
}