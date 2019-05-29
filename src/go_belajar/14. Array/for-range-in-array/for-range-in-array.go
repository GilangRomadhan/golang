package main

import "fmt"

func main() {
	var fruits = [4]string{"mangga", "salak", "jeruk","apple"}
	for i, fruit := range fruits {
		fmt.Printf("\nBuah %d %s", i, fruit)
	} 
}