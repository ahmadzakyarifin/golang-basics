package main

import "fmt"

// ====================
// 1. STRUCT
// ====================
type Mahasiswa struct {
	Nama string
	Umur int
}

// Fungsi mengubah umur Mahasiswa lewat pointer
func ubahUmur(m *Mahasiswa, umurBaru int) {
	m.Umur = umurBaru
}

// Fungsi ubah elemen array of struct (pakai pointer array)
func ubahArrayStruct(arr *[2]Mahasiswa) {
	arr[1].Nama = "Diganti Lewat Pointer Array"
}

// Fungsi menambah elemen ke slice pointer struct
func tambahKeSlice(slice *[]*Mahasiswa, m *Mahasiswa) {
	*slice = append(*slice, m)
}

// Fungsi ubah isi map dengan struct pointer
func ubahMapStruct(m map[string]*Mahasiswa, key string, namaBaru string) {
	if val, ok := m[key]; ok {
		val.Nama = namaBaru
	}
}

// ====================
// MAIN FUNCTION
// ====================
func main() {
	fmt.Println("=== Struct Langsung & Pointer ===")
	// Struct disimpan biasa
	m1 := Mahasiswa{Nama: "Budi", Umur: 20}
	m2 := m1
	m2.Nama = "Copy-an Budi"

	fmt.Println("Struct asli:", m1)  // Tidak berubah
	fmt.Println("Copy struct :", m2)

	// Struct lewat pointer
	ptr := &m1
	ptr.Umur = 25
	fmt.Println("Lewat pointer :", m1)

	ubahUmur(&m1, 30)
	fmt.Println("Lewat fungsi pointer:", m1)

	// ====================
	// ARRAY of STRUCT
	// ====================
	fmt.Println("\n=== Array of Struct ===")
	arr := [2]Mahasiswa{
		{Nama: "A", Umur: 18},
		{Nama: "B", Umur: 19},
	}
	fmt.Println("Sebelum:", arr)
	ubahArrayStruct(&arr)
	fmt.Println("Sesudah:", arr)

	// ====================
	// SLICE of POINTER ke STRUCT
	// ====================
	fmt.Println("\n=== Slice of Pointer Struct ===")
	slc := []*Mahasiswa{}
	m3 := &Mahasiswa{Nama: "Dina", Umur: 21}
	tambahKeSlice(&slc, m3)

	m4 := &Mahasiswa{Nama: "Rudi", Umur: 22}
	tambahKeSlice(&slc, m4)

	// Ubah salah satu struct dalam slice
	slc[0].Umur = 23

	for i, m := range slc {
		fmt.Printf("Index %d: %+v\n", i, *m)
	}

	// ====================
	// MAP of POINTER to STRUCT
	// ====================
	fmt.Println("\n=== Map of Pointer Struct ===")
	mapMhs := map[string]*Mahasiswa{
		"x": {Nama: "Toni", Umur: 20},
	}
	fmt.Println("Sebelum:", mapMhs["x"])

	ubahMapStruct(mapMhs, "x", "Anton")
	fmt.Println("Sesudah:", mapMhs["x"])
}
