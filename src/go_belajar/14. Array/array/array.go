package main

import "fmt"

func main() {

	var name [4]string
	name[0] = "satu"
	name[1] = "dua"
	name[2] = "tiga"
	name[3] = "empat"

	fmt.Println(name[1], name[2], name[3], name[0])
	
}