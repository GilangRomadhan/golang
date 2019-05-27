package main

import "fmt"

func main() {
	var point = 10

	if point > 7 {
		switch point {
		case 10 :
			fmt.Print("Perfect")
		default :
			fmt.Print("nice!")	
		}
	} else {
		if point == 5 {
			fmt.Print("not bad")
		} else if point == 3 {
			fmt.Print("Keep Trying")
		} else {
			fmt.Print("You can do it")
			if point == 0 {
				fmt.Print(" Try more harder")
			}	 
		}
	}
}