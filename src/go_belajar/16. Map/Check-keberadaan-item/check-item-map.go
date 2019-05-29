package main

import "fmt"

func main() {

	month := map[string]int {"Januari" : 1, "Februari" : 2}

	var value, isExist = month ["Januari"]
	
	if isExist {
		fmt.Println(value)
	
	}else {
		fmt.Println("Item not exist")
	}
	
}