/*
Belajar pengakesesan slice dengan 3 indeks
*/
package main

import "fmt"

func main() {
	aktivitas := []string{"Makan", "Minum", "Tidur"}
	aktivitas1 := aktivitas[0:2]
	aktivitas2 := aktivitas[0:2:2]

	fmt.Println()
	fmt.Println(aktivitas)
	fmt.Println(len(aktivitas))
	fmt.Println(cap(aktivitas))

	fmt.Println()
	fmt.Println(aktivitas1)
	fmt.Println(len(aktivitas1))
	fmt.Println(cap(aktivitas1))

	fmt.Println()
	fmt.Println(aktivitas2)
	fmt.Println(len(aktivitas2))
	fmt.Println(cap(aktivitas2))

}
