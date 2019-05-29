package main

import "fmt"

func main() {
	var chicken = map[string]int{"Ayam_negri" : 10, "Ayam_kampung" : 5}

	fmt.Println(len(chicken))
	fmt.Println(chicken)
	
	delete(chicken, "Ayam_negri")

	fmt.Println(len(chicken))
	fmt.Println(chicken)

}