package main

import "fmt"

func main() {
	var point = 10

	switch point {
	case 10:
		fmt.Println("Perfect")
	case 7,8,9 :
		fmt.Println("Good")
	case 5,6 :
		fmt.Println("Good enough")
	default :
	{
		fmt.Println("You are not qualified")
		fmt.Println("You point are to much")
	}
		
	}
}