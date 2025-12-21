package main

import (
	"bufio"
	"fmt"
	"os"
)


func main(){
	
	// CREATE TXT
	namaFile := "coba.txt"
	c,err := os.Create(namaFile)
	if err != nil {
		fmt.Println("Gagal membuat file : ", err)
		return
	}
	c.WriteString("Buku bankir : ")
	defer c.Close()

	fmt.Println()
	
	// UPDATE TXT
	u, err := os.OpenFile(namaFile,os.O_APPEND|os.O_WRONLY,0644)
	if err != nil {
		fmt.Println("Gagal membuka file : ", err)
		return
	}
	u.WriteString(" Uang keluar Rp 50.000 (26/03/21)")
	defer u.Close()

	
	// READ TXT
	r,err := os.Open(namaFile)
	if err != nil {
		fmt.Println("Gagal membuka file : ", err)
		return
	}
	defer r.Close()

	scanner := bufio.NewScanner(r)

	for scanner.Scan(){
		fmt.Println(scanner.Text())
	}

	// Ini menimpa laporan sebelumnya
	dataFinal := []byte("=== LAPORAN FINAL BANK ===\nTotal Saldo: Rp750.000\nStatus: SELESAI\n")
	err = os.WriteFile(namaFile, dataFinal, 0644)
	if err != nil {
		fmt.Println(err)
	}

	
	//  DELETE TXT
	// os.Remove(namaFile)
	// fmt.Println(" File telah dihapus")

	fmt.Println("Selesai! Silakan cek file:", namaFile)

}
