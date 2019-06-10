package main

import "fmt"

func main() {
	divideNumber(2, 0)
	divideNumber(10, 2)
	divideNumber(0, 2)
	divideNumber(6, 2)

	
}

func divideNumber(m, n int)  { //mendeklarasikan m, dan n dengan tipe data integer.
	if n == 0 {
		fmt.Printf("Invalid divider. %d cannot divided by %d\n", m, n)
		return // return berfungsi untuk menstop proses divider.
	}
	var res = m / n // variable res berisi rumus pembagian m / n
	fmt.Printf("%d divided by %d = %d\n", m, n, res) // tampilan
	
	
}