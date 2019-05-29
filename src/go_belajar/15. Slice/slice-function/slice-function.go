package main

import "fmt"

func main() {

	//Fungsi len

	var buah = []string{"apple","melon","magga","jeruk"}

	fmt.Print("Ada berapa banyak buah ", len(buah))//function len untuk mengetahu jumlah elemen slice

	//Fungsi cap

	var status =  []string{"makan","minum","tidur","berdoa"}
	fmt.Println("\nstatus len", len(status))
	fmt.Println("status cap", cap(status))

	var status1 = status[0:3]
	
	fmt.Println("status1 len", len(status1))
	fmt.Println("status2 cap", cap(status1))

	var status2 = status[1:4]

	fmt.Println("status2 len", len(status2))
	fmt.Println("status2 cap", cap(status2))

	//fungsi append

	var nama = []string{"gilang","sigit","paul","tri"}
	var nama2 = append(nama, "kassle")

	fmt.Println(nama)
	fmt.Println(nama2)

}