package main

import "fmt"

func main() {

	var point = 3

	switch  {
	case point == 8:
		fmt.Print("perfect")
	case (point < 8) && (point > 5):
		fmt.Print("Awesome")
	default :
	    {
			fmt.Print("Not bad ")
			fmt.Print("You need to learn more")

		}
		
	}
	
}