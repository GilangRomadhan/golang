package main

import "fmt"

func main() {

	//Fungsi len

	var buah = []string{"apple","melon","magga","jeruk"}

	fmt.Print("Ada berapa banyak buah ", len(buah))//function len untuk mengetahu jumlah elemen slice

	//Fungsi cap

	var status =  []string{"makan","minum","tidur","berdoa"}
	fmt.Println()
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

	fmt.Println()
	fmt.Println(nama)
	fmt.Println(nama2)

	//contoh lain fungsi append()

	var hewan = []string{"kucing", "tikus", "kambing"}
	var hewanA = hewan[0:2]

	fmt.Println()
	fmt.Println(len(hewanA))
	fmt.Println(cap(hewanA))

	fmt.Println(hewan)
	fmt.Println(hewanA)

	var hewanC = append(hewanA, "kambing")

	fmt.Println()
	fmt.Println(hewan)
	fmt.Println(hewanA)
	fmt.Println(hewanC)

	//Fungsi copy()

	srt := make([]string, 3)
	dst := [] string{"A","B","C","D"}
	n := copy(srt, dst)

	fmt.Println()
	fmt.Println(dst)//[A B C D]
	fmt.Println(srt)//[A B C]
	fmt.Println(n)//3

	A := []string{"apple", "melon", "orange"}
	B := []string{"jake_fruit", "salak"}
	C := copy(A, B)

	fmt.Println()
	fmt.Println(A)//[jake_fruit salak orange]
	fmt.Println(B)//[jake_fruit salak]
	fmt.Println(C)//2

}