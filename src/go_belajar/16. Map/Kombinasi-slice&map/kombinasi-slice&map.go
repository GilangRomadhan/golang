package main

import "fmt"

func main() {
	var chicken = []map[string]string{
		{"nama" : "gilang romadhan", "gender" : "male" },
		{"nama" : "Tri", "gender" : "male" },
		{"nama" : "Sigit", "gender" : "male" },

	}

	for _, value := range chicken {
		fmt.Print("\n", value["nama"], "\t :", value["gender"])
		
	}
	//penggunaan key dan value string yang berbeda
	var chicken2 = []map[string]string{
		{"nama" : "gilang", "gender" : "male"},
		{"addres" : "wonosari", "BOD" : "07031993"},
	}

	for _, value2 := range chicken2 {
		fmt.Print("\n", value2["nama"], "\t",value2["gender"])
		
	}

}