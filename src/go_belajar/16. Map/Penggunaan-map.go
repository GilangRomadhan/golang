package main

import "fmt"

func main() {

	var chicken map[string]int // initialisasi map
	chicken = map[string]int{}//set default map string

	chicken["januari"] = 50
	chicken["februari"] = 30

	fmt.Print("chicken", chicken["januari"])
	fmt.Print("chicken", chicken["februari"])

}