package main

import (
	"fmt"
	"sync"
)

var saldo = 0
var mutex sync.Mutex
var wg sync.WaitGroup

func kasir(nama string, jumlah int) {
	defer wg.Done()

	for i := 0; i < jumlah; i++ {
		mutex.Lock()
		saldo += 1
		mutex.Unlock()
	}

	fmt.Println(nama, "selesai mengisi saldo")
}

func main() {
	jumlahKasir := 5
	jumlahIsi := 1000

	for i := 0; i < jumlahKasir; i++ {
		wg.Add(1)
		go kasir(fmt.Sprintf("Kasir %d", i+1), jumlahIsi)
	}

	wg.Wait()
	fmt.Println("Kasir telah selesai semua")
	fmt.Println("Total saldo:", saldo)
}
