package main

import (
	"crypto/rand"
	"fmt"
	"math/big"
)

func main() {
	fmt.Println("Welcome to maths n golang ")

	// var myNumberOne int = 2
	// var myNumberTwo float64 = 4.5

	// fmt.Println("The sum is: ", myNumberOne + int(myNumberTwo))

	//! random number from math pkg
	// rand.Seed(time.Now().UnixNano())
	// fmt.Println(rand.Intn(5) + 1)

	//! random number from crypto pkg
	myRandonNum, _ := rand.Int(rand.Reader, big.NewInt(5))
	fmt.Println("My random number", myRandonNum)
}