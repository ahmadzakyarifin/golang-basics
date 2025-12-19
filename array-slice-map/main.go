package main

import "fmt"

func main() {
	// Array 1
	array1 := [3]string{"1", "2", "3"}
	fmt.Println(array1)
	fmt.Println(len(array1))
	fmt.Println(cap(array1))

	array11 := [3]string{
		"1",
		"2",
		"3",
	}
	fmt.Println(array11)
	fmt.Println(len(array11))
	fmt.Println(cap(array11))

	array111 := [3]string{}
	array111[0] = "1"
	array111[1] = "2"
	array111[2] = "3"
	fmt.Println(array111)
	fmt.Println(len(array111))
	fmt.Println(cap(array111))

	array1111 := [...]string{"1", "2"}
	fmt.Println(array1111)
	fmt.Println(len(array1111))
	fmt.Println(cap(array1111))

	array11111 := [...]string{
		"1",
		"2",
	}
	fmt.Println(array11111)
	fmt.Println(len(array11111))
	fmt.Println(cap(array11111))

	// Array 2
	var array2 = [3]string{"a", "b", "c"}
	fmt.Println(array2)
	fmt.Println(len(array2))
	fmt.Println(cap(array2))

	var array22 = [3]string{
		"a",
		"b",
		"c",
	}
	fmt.Println(array22)
	fmt.Println(len(array22))
	fmt.Println(cap(array22))

	var array222 = [3]string{}
	array222[0] = "a"
	array222[1] = "b"
	array222[2] = "c"
	fmt.Println(array222)
	fmt.Println(len(array222))
	fmt.Println(cap(array222))

	var array2222 = [...]string{"1"}
	fmt.Println(array2222)
	fmt.Println(len(array2222))
	fmt.Println(cap(array2222))

	var array22222 = [...]string{
		"1",
	}
	fmt.Println(array22222)
	fmt.Println(len(array22222))
	fmt.Println(cap(array22222))

	// Slice 1
	slice1 := []string{"ikan", "ayam", "sapi"}
	fmt.Println(slice1)
	fmt.Println(len(slice1))
	fmt.Println(cap(slice1))

	slice11 := []string{
		"ikan",
		"ayam",
		"sapi",
	}
	fmt.Println(slice11)
	fmt.Println(len(slice11))
	fmt.Println(cap(slice11))

	// slice kosong, tidak bisa langsung akses index
	slice111 := make([]string, 3)
	slice111[0] = "ikan"
	slice111[1] = "ayam"
	slice111[2] = "sapi"
	fmt.Println(slice111)
	fmt.Println(len(slice111))
	fmt.Println(cap(slice111))

	// Slice 2
	var slice2 = []string{"ikan", "ayam", "sapi"}
	fmt.Println(slice2)
	fmt.Println(len(slice2))
	fmt.Println(cap(slice2))

	var slice22 = []string{
		"ikan",
		"ayam",
		"sapi",
	}
	fmt.Println(slice22)
	fmt.Println(len(slice22))
	fmt.Println(cap(slice22))

	var slice222 = make([]string, 3)
	slice222[0] = "ikan"
	slice222[1] = "ayam"
	slice222[2] = "sapi"
	fmt.Println(slice222)
	fmt.Println(len(slice222))
	fmt.Println(cap(slice222))

	// Umumnya slice kosong biar bisa di-append
	var data1 []string
	data1 = append(data1, "satu")

	data2 := []string{}
	data2 = append(data2, "dua")

	var data3 = []string{}
	data3 = append(data3, "tiga")

	fmt.Println("data1:", data1) // ["satu"]
	fmt.Println("data2:", data2) // ["dua"]
	fmt.Println("data3:", data3) // ["tiga"]

	// make slice
	slice3 := make([]string, 2, 10)
	//  2 nilai awal 
	//  10 capasitas semuanya
	slice3[0] = "unta"
	slice3[1] = "kambing"

	fmt.Println(slice3)
	fmt.Println(len(slice3))
	fmt.Println(cap(slice3))

	// Map 1: RECOMMENDED buat map kosong
	map1 := make(map[string]int)
	map1["apel"] = 5

	fmt.Println("Map 1:")
	fmt.Println("apel =", map1["apel"])
	fmt.Println("Jumlah elemen:", len(map1))
	fmt.Println()

	// Map 2: langsung inisialisasi isi
	buah := map[string]int{
		"jeruk":  3,
		"mangga": 7,
	}

	fmt.Println("Map 2:")
	fmt.Println("jeruk =", buah["jeruk"])
	fmt.Println("mangga =", buah["mangga"])
	fmt.Println("Jumlah elemen:", len(buah))
	fmt.Println()

	// Map 3: map kosong literal
	kosong := map[string]string{}
	kosong["nama"] = "Dino"

	fmt.Println("Map 3:")
	fmt.Println("nama =", kosong["nama"])
	fmt.Println("Jumlah elemen:", len(kosong))


	//   for range di array , slice , map
	for index,item := range slice111{
		fmt.Println(index,item)
	}
	for key,value := range map1{
		fmt.Println(key,value)
	}
	
}


// catatan 
// make([]string, 2, 10)     // ❌ Error: hasil make tidak digunakan
// make(map[string]int)      // ❌ Error: hasil make tidak digunakan
// harus simpen variabel

//  buat nampilin array da slice bisa pakai for i := range misalnya{slice[i]}, atau for manual atau bisa satu satu berdasarkan indek misal slice[0],dst
