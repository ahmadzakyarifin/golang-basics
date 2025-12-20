package main

import "fmt"

func main() {
    jalankanProsesPenting()
    fmt.Println(" Aplikasi Selesai dengan Aman ") 
}

func jalankanProsesPenting() {
    // Jalankan Defer LIFO
    defer fmt.Println("=> [Defer 1] Koneksi Database ditutup")

	// pakai Anonymous Function (fungsi tanpa nama) atau sering juga disebut Closure
    defer func() {
        r := recover(); 
		if r != nil {
            fmt.Println(" [Recover] Menangkap error:", r)
        }
    }()

    fmt.Println("Proses: Mengecek validasi data...")

    adaError := true
    if adaError {
        panic("ERROR FATAL: Data tidak valid!")
    }

    // Baris ini TIDAK akan dicetak jika terjadi panic
    fmt.Println("Proses: Data Berhasil diproses")
}