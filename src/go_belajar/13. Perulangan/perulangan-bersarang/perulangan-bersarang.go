package main

import "fmt"

func main() {

	for indexa := 1; indexa < 5; indexa++ {
		for indexb := indexa; indexb < 5; indexb++ {

			fmt.Print(indexb, " ")
			
		}

		fmt.Println()
		
	}
	
}