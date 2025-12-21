package main

import (
	"fmt"
	"sync"
)

// kasir bertindak sebagai Producer
func kasir(id int, jumlahKoin int, pipa chan<- int, wg *sync.WaitGroup) {
	defer wg.Done()

	for i := 1; i <= jumlahKoin; i++ {
		// Kasir memasukkan koin ke pipa. 
		// Karena ada buffer, kasir tidak perlu menunggu petugas selama laci belum penuh.
		pipa <- 1 
	}
	fmt.Printf("Kasir %d: Selesai mengirim %d koin ke pipa\n", id, jumlahKoin)
}

func main() {
	const jumlahKasir = 5
	const koinPerKasir = 1000
	
	// Membuat Channel dengan BUFFER (Laci Penampung) sebesar 100.
	// Kasir bisa "nyicil" taruh koin sampai 100 biji tanpa harus nunggu petugas.
	pipa := make(chan int, 100) 

	var wg sync.WaitGroup
	totalSaldo := 0

	// 1. Jalankan para Kasir (Producer)
	for i := 1; i <= jumlahKasir; i++ {
		wg.Add(1)
		go kasir(i, koinPerKasir, pipa, &wg)
	}

	// 2. Jalankan Petugas Pencatat (Consumer)
	// Kita gunakan channel 'finish' untuk tahu kapan Petugas benar-benar selesai.
	finish := make(chan bool)
	go func() {
		// Petugas terus mengambil koin selama pipa belum ditutup
		for koin := range pipa {
			totalSaldo += koin
		}
		finish <- true
	}()

	// 3. Tunggu sampai semua Kasir selesai menaruh koin
	wg.Wait()

	// 4. Tutup pipa (Sinyal bahwa tidak ada lagi koin yang akan masuk)
	close(pipa)

	// 5. Tunggu petugas selesai mencatat sisa koin yang ada di buffer
	<-finish

	fmt.Println("--------------------------------------")
	fmt.Printf("Semua Kasir telah selesai.\n")
	fmt.Printf("Total Saldo Akhir di Buku Besar: %d\n", totalSaldo)
}