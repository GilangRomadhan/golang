package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().Unix())
	var randomValue int

	randomValue = randomWithRange(2, 10)
	fmt.Println("random Number", randomValue)
	randomValue = randomWithRange(2, 10)
	fmt.Println("random Number", randomValue)
	randomValue = randomWithRange(2, 10)
	fmt.Println("random Number", randomValue)

}

func randomWithRange(min, max int) int  {
	var value = rand.Int() % (max - min + 1) + min
	return value
	
}