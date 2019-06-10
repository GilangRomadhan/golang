package main

import "fmt"
import "strings"

func main() {
	var name = []string{"John", "wick"}
	printMessage("Hallo", name)
	
}

func printMessage(message string, arr []string)  {
	var MessageString = strings.Join(arr, " ")
	fmt.Print(message, MessageString)
	
}