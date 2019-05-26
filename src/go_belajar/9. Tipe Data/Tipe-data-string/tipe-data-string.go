/*
Program ini dibuat untuk belajar tipe data strings

*/

package main

import "fmt"
func main() {

	var nama string = "Gilang"
	fmt.Printf("What is your name ? \nMy Name is %s\n", nama)

	var message = // using `` untuk menggunakan multiline strings
	`	Nama saya "Gilang Romadhan".
	Salam kenal
	Mari belajar golang berasama` 

	fmt.Println(message)
}