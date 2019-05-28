package main

import "fmt"

func main() {
	var berhitung = [...]string {"satu", "dua", "tiga", "empat"}

	fmt.Print("Beritung", berhitung)
	fmt.Print("\nBerhitung ", len(berhitung))
}