package main

import "fmt"

func main() {

//inisialisasi map horisontal
	var nama = map[string]int{"Gilang" : 20, "Tri" : 30, "sigit" : 40}

//inisialisasi map vertical
	var nama2 = map[string]int{
		"Aziz" : 20,
		"Dwiki" : 30,
		"Catur" : 25,
	}
//inisialisasi map kosong
	nama3 := map[string]int{}
	nama4 := make(map[string]int)
	nama5 := *new(map[string]int)

fmt.Println(nama)
fmt.Println(nama2)
fmt.Println(nama3)
fmt.Println(nama4)
fmt.Println(nama5)
	
}