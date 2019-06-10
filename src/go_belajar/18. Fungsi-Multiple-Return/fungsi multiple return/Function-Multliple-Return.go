package main

import "fmt"
import "math"

func main() {
var diameter float64 = 15
var area, circumference = calculate(diameter)

fmt.Printf("Luas lingkaran \t\t: %2.f ", area)
fmt.Printf("\nKeliling Lingkaran \t: %2.f \n ", circumference)

	
}

func  calculate(d float64) (float64, float64)  {
	//Hitung Luas
	var area = math.Pi * math.Pow(d / 2, 2)

	//Hitung Keliling
	var circumference = math.Pi * d

	return area, circumference
	
}