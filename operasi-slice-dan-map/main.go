package main

import "fmt"

func main() {
	fmt.Println("=== LABORATURIUM KOLEKSI DATA GOLANG ===")
	
	latihanSliceLengkap()

	fmt.Printf("============================================")

	latihanMapLengkap()
}

func latihanSliceLengkap() {
	fmt.Println("[1] EKSPLORASI SLICE")

	numbers := []int{1, 2, 3}
	fmt.Printf("Awal: %v | Len: %d | Cap: %d\n", numbers, len(numbers), cap(numbers))

	// 2. Ubah
	numbers[2] = 8
	fmt.Println("Setelah Ubah Index 2:", numbers)

	// 3. Tambah (Append)
	numbers = append(numbers, 10, 20, 30)
	fmt.Printf("Setelah Append: %v | Len: %d | Cap: %d\n", numbers, len(numbers), cap(numbers))

	// 4. Slicing dari Array 
	fmt.Println("\n--- Mekanisme Slicing dari Array ---")
	arrayAsli := [5]int{1, 2, 3, 4, 5}
	potongan := arrayAsli[1:3] // mengambil [2, 3]

	fmt.Println("Array Asli :", arrayAsli)
	fmt.Println("Slice      :", potongan)

	// Ubah Slice merubah Array
	potongan[0] = 100 
	fmt.Println("Efek Referensi (potongan[0]=100):")
	fmt.Println("Array Asli :", arrayAsli)

	// 5. Append melampaui Capacity (Putus Referensi)
	fmt.Printf("Cap Potongan: %d\n", cap(potongan))
	potongan = append(potongan, 10, 20, 30, 40, 50) 
	
	fmt.Println("Setelah Append Melebihi Cap:")
	fmt.Println("Array Asli :", arrayAsli) // Tetap (Aman)
	fmt.Println("Slice Baru :", potongan) // Pindah ke Alamat Memori Baru
}

func latihanMapLengkap() {
	fmt.Println("[2] EKSPLORASI MAP")

	map1 := map[string]int{"ikan": 1, "ayam": 2} // Literal
	var map2 = make(map[string]int)              // Make
	map2["kucing"] = 5

	fmt.Println("Map 1:", map1)
	fmt.Println("Map 2:", map2)

	// Membuat Map Kosong (Menghindari Nil Map)
	mapKosong1 := make(map[string]int)
	mapKosong2 := map[string]int{}
	fmt.Println("Map Kosong:", mapKosong1, mapKosong2)

	// Update & Delete
	map1["ikan"] = 10 // Update
	delete(map1, "ayam") // Delete
	fmt.Println("Map 1 (Setelah Update & Delete):", map1)

	// Cek Panjang
	fmt.Println("Panjang Map 1:", len(map1))

	// Sifat Pass-by-Reference (Sinkronisasi Memori)
	fmt.Println("\n--- Uji Referensi Map ---")
	mapAsli := map[string]int{"status": 200}
	mapAlias := mapAsli

	mapAlias["status"] = 404
	fmt.Println("Map Alias:", mapAlias)
	fmt.Println("Map Asli :", mapAsli, "<- Ikut Berubah karena satu alamat!")

	// Iterasi Map
	fmt.Println("\nIterasi Map 1:")
	for key, value := range map1 {
		fmt.Printf("- %s: %d\n", key, value)
	}
}