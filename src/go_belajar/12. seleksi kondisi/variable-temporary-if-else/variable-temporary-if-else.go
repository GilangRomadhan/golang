/*
This program is made for learning temporary variable di else if
*/
package main

import "fmt"

func main() {

	var point = 7000.0

	if percent := point / 100; percent >= 100 {
		fmt.Printf("%.1f %s perfect!\n", percent, "%")
	} else if percent >= 70 {
		fmt.Printf("%.1f %s good", percent, "%")
	}else {
		fmt.Printf("%.1f %s not bad", percent, "%")
	}
	
}