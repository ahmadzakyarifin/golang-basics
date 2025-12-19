package main

import "fmt"

type Mahasiswa struct {
	Nama string
}

// --- KUMPULAN FUNGSI UPDATE ---

func updateVariabelFunc(x *int)      { *x = 40 }
func updateArrayFunc(arr *[3]int)    { arr[2] = 30 }
func updateSliceFunc(s []int)        { s[0] = 500 } // Slice isi otomatis reference
func updateMapFunc(m map[string]int) { m["a"] = 20 }  // Map otomatis reference
func (m *Mahasiswa) updateReceiver(baru string) { m.Nama = baru }

func main() {
	// ---------------------------------------------------------
	// 1. VARIABEL BIASA (int)
	// ---------------------------------------------------------
	fmt.Println("=== 1. VARIABEL ===")
	a := 10

	a = 20                  // Langsung
	
	temp := a               // Lewat Variabel Lain (Copy)
	temp = 100              // (a TIDAK berubah karena temp adalah fotokopi)

	ptrA := &a              // Lewat Variabel Pointer
	*ptrA = 30

	updateVariabelFunc(&a)  // Lewat Fungsi Pointer
	fmt.Printf("Hasil a: %d (temp tadi: %d)\n", a, temp)

	// ---------------------------------------------------------
	// 2. ARRAY
	// ---------------------------------------------------------
	fmt.Println("\n=== 2. ARRAY ===")
	arr := [3]int{1, 2, 3}

	arr[0] = 10             // Langsung
	
	ptrArr := &arr          // Lewat Pointer
	ptrArr[1] = 20

	updateArrayFunc(&arr)   // Lewat Fungsi Pointer
	fmt.Println("Hasil Array:", arr)

	// ---------------------------------------------------------
	// 3. SLICE (Pointer ke Array Dasar)
	// ---------------------------------------------------------
	fmt.Println("\n=== 3. SLICE (Update Isi) ===")
	slc := []int{1, 2, 3}

	slc[0] = 100            // Langsung
	
	aliasSlc := slc         // Lewat Variabel Lain (Reference)
	aliasSlc[1] = 200       // (slc asli IKUT berubah karena menunjuk array yang sama)

	updateSliceFunc(slc)    // Lewat Fungsi (Tanpa * pun berubah)
	fmt.Println("Hasil Slice:", slc)

	// ---------------------------------------------------------
	// 4. MAP (Murni Reference)
	// ---------------------------------------------------------
	fmt.Println("\n=== 4. MAP ===")
	m := map[string]int{"a": 1}

	m["a"] = 10             // Langsung
	
	aliasM := m             // Variabel Lain (Reference)
	aliasM["a"] = 15

	updateMapFunc(m)        // Lewat Fungsi
	fmt.Println("Hasil Map:", m)

	// ---------------------------------------------------------
	// 5. STRUCT (Value Type)
	// ---------------------------------------------------------
	fmt.Println("\n=== 5. STRUCT ===")
	mhs := Mahasiswa{Nama: "Zaky"}

	mhs.Nama = "Budi"       // Langsung
	
	ptrMhs := &mhs          // Variabel Pointer
	ptrMhs.Nama = "Andi"

	mhs.updateReceiver("Ahmad") // Pointer Receiver
	fmt.Println("Hasil Struct:", mhs.Nama)
}