package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println("=== KONVERSI TIPE DATA DI GOLANG ===")

	// int -> float64
	var nilaiInt int = 100
	var nilaiFloat float64 = float64(nilaiInt)
	fmt.Println("int ke float64:", nilaiFloat)

	// float64 -> int
	var desimal float64 = 123.456
	var dibulatkan int = int(desimal)
	fmt.Println("float64 ke int:", dibulatkan)

	// int -> uint
	var nilaiIntNegatif int = -50
	var nilaiUint uint = uint(nilaiIntNegatif) // hati-hati overflow
	fmt.Println("int negatif ke uint:", nilaiUint)

	// int16 -> int32
	var kecil int16 = 32000
	var besar int32 = int32(kecil)
	fmt.Println("int16 ke int32:", besar)

	// int -> string
	var angka int = 123
	var teks string = strconv.Itoa(angka)
	fmt.Println("int ke string:", teks)

	// string -> int
	var str string = "456"
	var konversiInt, err = strconv.Atoi(str)
	if err == nil {
		fmt.Println("string ke int:", konversiInt)
	} else {
		fmt.Println("Gagal konversi string ke int:", err)
	}

	//  overflow dari int
	var besarLagi int = 300
	var kecilLagi uint8 = uint8(besarLagi) // overflow! 300 > 255
	fmt.Println("int 300 ke uint8 (overflow):", kecilLagi)

	// float32 ke float64 dan sebaliknya
	var f32 float32 = 1.23
	var f64 float64 = float64(f32)
	var backToF32 float32 = float32(f64)
	fmt.Println("float32 ke float64:", f64)
	fmt.Println("float64 ke float32:", backToF32)

	//  catatan jika melebihi kapasitas maka akan bali lagi ke jumlah minimumnya 
}
