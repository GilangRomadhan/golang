package main

import "fmt"

func main() {

	// deklarasi multi variable menggunakan var
	var nama, alamat, umur string
	nama, alamat, umur = "Gilang", "wonosari", "25"

	//deklarasi multi variable menggunakan type interaface
	satu, dua, tiga := "satu", "dua", "tiga" 

	fmt.Printf("Who are you ? %s %s %s", nama, alamat, umur)
	fmt.Printf("\nBerhitung %s %s %s", satu, dua, tiga)
	
}