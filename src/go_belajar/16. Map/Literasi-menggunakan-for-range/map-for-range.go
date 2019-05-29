package main

import "fmt"

func main() {
	var name = map[string]int{
		"robert" : 10,
		"Tony" : 4,
		"issable" : 5,
		"granger" : 7,
	}

	for key, val := range name {
		fmt.Println(key, "\t:", val,)
		
	}
	
}