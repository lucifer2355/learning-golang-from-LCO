package main

import "fmt"

func main() {
	fmt.Println("Welcome to class of pointer")

	var ptr *string
	fmt.Println("Value of pointer is ", ptr)

	myNumber := 23
	var num = &myNumber
	fmt.Println("My number is", num)
	fmt.Println("My number is", *num)

	*num = *num * 2
	fmt.Println("Multiple by 2", myNumber)

}