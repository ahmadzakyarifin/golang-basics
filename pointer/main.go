package main

import "fmt"

func main() {
	// Ini Variabel Biasa (Value)
	name := "Zaky"
	
	// Ini Variabel Pointer (Reference)
	// menyimpan ALAMAT memori 'name' ke dalam variabel 'namePtr'
	var namePtr *string = &name

	fmt.Println("--- Kondisi Awal ---")
	fmt.Println("Nilai name      :", name)    // Zaky
	fmt.Println("Alamat name (&) :", &name)   // Contoh: 0xc000010250
	fmt.Println("Isi namePtr     :", namePtr) // Sama dengan alamat name

	// Mengubah nilai melalui Pointer (Dereferencing)
	// Menggunakan operator * untuk "mendatangi" alamat dan mengganti isinya
	*namePtr = "Zaky Arifin"

	fmt.Println("\n--- Setelah Diubah via Pointer ---")
	fmt.Println("Nilai name sekarang :", name)    // Ikut berubah jadi Zaky Arifin
	fmt.Println("Nilai via *namePtr  :", *namePtr) // Zaky Arifin

	//Pass by Value vs Pass by Reference
	score := 80
	
	tambahPointValue(score)
	fmt.Println("\nScore setelah Pass by Value:", score) // Tetap 80

	tambahPointReference(&score)
	fmt.Println("Score setelah Pass by Reference:", score) // Jadi 100
}

func tambahPointValue(s int) {
	s = s + 20 // Hanya mengubah salinan
}

func tambahPointReference(s *int) {
	*s = *s + 20 // Mengubah data asli di alamat memori
}