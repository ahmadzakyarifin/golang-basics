package main

import (
	"errors"
	"fmt"
	"os"
	"time"
)

type Produk struct {
	ID    int
	Nama  string
	Harga float64
}

func main() {
	// =============================================================
	fmt.Print("Cetak polos (tanpa enter)")
	fmt.Println(" - Println otomatis tambah baris baru.")

	// =============================================================
	currentTime := time.Now().Format("15:04:05")
	pesanLog := fmt.Sprintf("[%s] INFO: Log disimpan ke variabel string", currentTime) //Analogi: Menulis surat tapi disimpan dalam variabel (String)
	fmt.Println("Hasil Sprintf:", pesanLog)

	// =============================================================
	fmt.Fprintln(os.Stdout, "[Stdout] Menulis ke jalur output normal.")
	fmt.Fprintf(os.Stderr, "[Stderr] CRITICAL: Menulis ke jalur error sistem!\n")

	// =============================================================
	p := Produk{ID: 101, Nama: "Monitor 4K", Harga: 4500000.75}
	
	fmt.Println("\n--- LAPORAN DATA PRODUK ---")
	fmt.Printf("ID Produk   : %d\n", p.ID)
	fmt.Printf("Nama Produk : %-15s (Tipe: %T)\n", p.Nama, p.Nama)
	fmt.Printf("Harga       : Rp %.2f\n", p.Harga)
	
	fmt.Printf("Object Detail: %+v\n", p)
	fmt.Printf("Alamat Memori: %p\n", &p)


	// =============================================================
	fmt.Println("\n--- ANALISIS ERROR HANDLING ---")

	errNotFound := errors.New("ErrNotFound: data tidak ada di DB")

	errFinal := fmt.Errorf("Gagal ambil produk ID %d: %w", p.ID, errNotFound)

	fmt.Println("Pesan Akhir (Wrapped):", errFinal)

	if errors.Is(errFinal, errNotFound) {
		fmt.Println("Akar Masalah: System mendeteksi identitas Sentinel Error asli.")
	}
}