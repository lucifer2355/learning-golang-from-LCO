package main

import "fmt"

func main() {
	fmt.Println("Welcome to FUNCTION in golang")
	//! Can't defined function inside the function
	greeter()

	result := adder(3, 5)
	fmt.Println("Result is:", result)

	proResult, myString := proAdder(2, 5, 6, 7)
	fmt.Println(myString, proResult)

}

func adder(valOne int, valTwo int) int {
	return valOne + valTwo
}

func proAdder(values ...int) (int, string) {
	total := 0

	for _, value := range values {
		total += value
	}

	return total, "Pro result is:"
}

func greeter() {
	fmt.Println("Namstey from golang")
}
