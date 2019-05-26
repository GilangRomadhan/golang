package main

import "fmt"

func main() {
	
	var left = false
	var right = true

	var leftAndRight = left && right
	fmt.Printf("left && right = (%t)", leftAndRight)

	var leftOrRight = left || right
	fmt.Printf("\nleft || right = (%t)", leftOrRight)

	var leftReverse = !left
	fmt.Printf("\n!left \t = (%t)", leftReverse)
}