/*
program ini dibuat untuk memahami switch dan case
*/

package main

import "fmt"

func main() {
	
	var point = 7

	switch point {
	case 10:
		fmt.Print("Perfect")
	case 8:
		fmt.Print("Good")
	default:
		fmt.Print("Not Bad")
		
	}

}