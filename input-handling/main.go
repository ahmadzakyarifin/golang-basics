package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	var kata1, kata2 string

	fmt.Println("=== 1. Contoh fmt.Scan ===")
	fmt.Print("Masukkan dua kata (dipisah spasi): ")
	fmt.Scan(&kata1, &kata2)
	fmt.Printf("Hasil Scan -> Kata 1: %s | Kata 2: %s\n", kata1, kata2)

	// Membersihkan sisa newline (\n) agar tidak merusak input selanjutnya
	bufio.NewReader(os.Stdin).ReadString('\n')
	fmt.Println()

	// Reset Variabel
	kata1, kata2 = "", ""

	fmt.Println("=== 2. Contoh fmt.Scanln ===")
	fmt.Print("Masukkan dua kata (dipisah spasi): ")
	// Scanln berhenti tepat saat tombol Enter ditekan
	fmt.Scanln(&kata1, &kata2)
	fmt.Printf("Hasil Scanln -> Kata 1: %s | Kata 2: %s\n", kata1, kata2)
	fmt.Println()

	fmt.Println("=== 3. Contoh bufio.Reader (Input Kalimat) ===")
	fmt.Print("Masukkan kalimat panjang: ")
	reader := bufio.NewReader(os.Stdin)
	// Membaca sampai Enter, aman untuk kalimat dengan spasi
	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(input)
	
	fmt.Println("Hasil bufio.Reader:")
	fmt.Printf("Kalimat: %s\n", input)
}