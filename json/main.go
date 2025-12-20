package main

import (
	"encoding/json"
	"fmt"
)


type User struct {
	Nama   string `json:"nama"`
	Alamat string `json:"alamat"`
	Umur   int    `json:"umur"`
}

func main() {

	fmt.Println("=== JSON di Golang ===")

	// 1. JSON -> Struct (PALING UMUM)
	jsonToStruct()

	// 2. JSON Array -> Slice of Struct
	jsonArrayToStruct()

	// 3. Struct -> JSON (Marshal)
	structToJson()

	// 4. JSON -> map (opsional / dinamis)
	jsonToMap()

}

// ======================
// 1. JSON -> Struct
// ======================
func jsonToStruct() {
	fmt.Println("\n[1] JSON ke Struct")

	jsonString := `{"nama":"Zaky","alamat":"Sukodadi","umur":19}`
	var user User

	err := json.Unmarshal([]byte(jsonString), &user)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println("Nama  :", user.Nama)
	fmt.Println("Alamat:", user.Alamat)
	fmt.Println("Umur  :", user.Umur)
}

// ======================
// 2. JSON Array -> Slice
// ======================
func jsonArrayToStruct() {
	fmt.Println("\n[2] JSON Array ke Slice Struct")

	jsonString := `[
		{"nama":"Zaky","alamat":"Sukodadi","umur":19},
		{"nama":"Nina","alamat":"Probolinggo","umur":22}
	]`

	var users []User

	err := json.Unmarshal([]byte(jsonString), &users)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	for i, u := range users {
		fmt.Printf(
			"%d. %s (%d tahun) - %s\n",
			i+1, u.Nama, u.Umur, u.Alamat,
		)
	}
}

// ======================
// 3. Struct -> JSON
// ======================
func structToJson() {
	fmt.Println("\n[3] Struct ke JSON")

	user := User{
		Nama:   "Zaky",
		Alamat: "Sukodadi",
		Umur:   19,
	}

	jsonData, err := json.Marshal(user)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println(string(jsonData))
}

// ======================
// 4. JSON -> map (opsional)
// ======================
func jsonToMap() {
	fmt.Println("\n[4] JSON ke map (opsional)")

	jsonString := `{"nama":"Zaky","alamat":"Sukodadi","umur":19}`
	var data map[string]interface{}

	err := json.Unmarshal([]byte(jsonString), &data)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println("Nama  :", data["nama"])
	fmt.Println("Alamat:", data["alamat"])
	fmt.Println("Umur  :", data["umur"]) // float64
}
