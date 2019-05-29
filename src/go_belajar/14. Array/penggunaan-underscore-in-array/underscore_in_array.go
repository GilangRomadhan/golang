package main

import "fmt"

func main() {
	var fruits = [4]string{"apple", "jeruk","melon","manggis"}
	for _, fruit := range fruits {
		fmt.Printf("\nBuah %s", fruit)

	}

	for i := range fruits {
		fmt.Printf("\nBuah %d", i)
		
	}
	
}