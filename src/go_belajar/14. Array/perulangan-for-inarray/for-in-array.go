package main

import "fmt"

func main() {
	var fruit = [4]string{"mangga", "manggis","timun","nanas",}
	for i := 0; i < len(fruit); i++ {

		fmt.Printf("Nama buah %d %s\n", i, fruit[i])
		
	}

}