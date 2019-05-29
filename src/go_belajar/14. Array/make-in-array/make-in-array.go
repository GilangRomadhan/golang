package main

import "fmt"

func main() {
	var fruits = make([]string, 2)

	fruits[0] = "apple"
	fruits[1] = "jeruk"

	fmt.Print(" buah ", fruits[0])
}