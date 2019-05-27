package main

import "fmt"

func main() {

	var point = 7

	switch  {
	case point == 8:
		fmt.Print("Perfect")
	case (point < 8) && (point > 5 ):
		fmt.Print("Awesome")
		fallthrough
	default:
		{
			fmt.Print(" Not bad ")
			fmt.Print(" You need to learn more ")
		}
		
	}
	
}