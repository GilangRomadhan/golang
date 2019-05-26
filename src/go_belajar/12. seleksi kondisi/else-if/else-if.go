package main

import "fmt"

func main() {
	var nilai = 3

	if nilai == 10 {
		fmt.Println(`Selamat kamu lulus dengan nilai sempurna`)
	} else if nilai > 6 {
		fmt.Println("Selamat kamu lulus")
	} else if nilai == 5 {
		fmt.Println("Maaf kamu hampir lulus")
	} else {
		fmt.Printf("Maaf kamu tidak lulus nilai kamu = %d", nilai)
	}	
}