package main

import "fmt"

func main(){
	// ==============================================
	// 				PERCABANGAN
	// ==============================================
	fmt.Println("===== Percabangan =====")
	
	//  if else
	nilai := 20
	if nilai >= 80 {
		fmt.Println("Anda mendapatkan nilai A")
	}else if nilai >=60 {
		fmt.Println("Anda mendapatkan nilai B")
	}else {
		fmt.Println("Anda mendapatkan nilai C")
	}

	// switch case
	hari := "jum'at"
	switch hari {
	case "minggu","senin","selasa","rabu" :
		fmt.Println("kategori : Hari masuk kuliah")
	case "kamis","jum'at","sabtu":
		fmt.Println("kategori : Hari libur kuliah")
	default:
		fmt.Println("Kategori: Hari tidak valid")
	}

	// ==============================================
	// 				PERULANGAN 
	// ==============================================

	println("===== Perulangan =====")

	
	// for standart
	for i := 1; i <= 3; i++ {
		fmt.Printf("%d ", i)
	}

	fmt.Println()

	//  for sebagi while
	counter := 1
	for counter <= 10 {
		fmt.Printf("%d",counter)
		counter ++
	}

	fmt.Println()

	//  for range
	buahs := []string {"Mangga","Nanas","Apel"}
	for i,b := range buahs {
		fmt.Printf("  [%d] %s\n", i, b)
	}

	// =================================================
			// KOMBINASI: LOOP + IF + BREAK/CONTINUE
	// =================================================

	fmt.Println("Mencatak angka ganjil (1-10), tapi berhenti di angka 7")
	for i := 1 ; i <= 10; i++ {
		if i % 2 == 0{
			continue
		}
		if i > 7 {
			break
		}
		fmt.Printf("%d",i)
	}
	fmt.Println()

}