package main

import "fmt"

func main() {
	var angka1 = [3][2]int{[2]int{1, 2}, [2]int{4, 5}, [2]int{7, 8}}
	var angka2 = [2][3]int{{1, 2, 3}, {4, 5, 6}}
	
	fmt.Print("Angka1", angka1)
	fmt.Print("\nAngka2", angka2)

}