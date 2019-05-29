package main

import "fmt"

func main() {
	var buah = []string{"apple", "melon", "jake_fruit", "watermelon"}

	var buah1 = buah[0:3]
	var buah2 = buah[1:4]

	var buah1a = buah1[0:2]
	var buah2a = buah2[0:1] 

	fmt.Println(buah1) //apple, melon, jake_fruit
	fmt.Println(buah2) //melon, jake_fruit, watermelon
	fmt.Println(buah1a) //apple, melon
	fmt.Println(buah2a) //melon

	buah2[0] = "Pineapple"

	fmt.Println()
	fmt.Println(buah1) //apple, pineaple, jake_fruit
	fmt.Println(buah2) //pineaple, jake_fruit, watermelon
	fmt.Println(buah1a) //apple, pineaple
	fmt.Println(buah2a) //watermelon
}