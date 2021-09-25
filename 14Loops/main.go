package main

import "fmt"

func main() {
	fmt.Println("Welcome to LOOPS in golang")

	days := []string{"Sunday", "Tuesday", "Wednesday", "Friday", "Saturday"}

	for d := 0; d < len(days); d++ {
		fmt.Println(days[d])
	}
}
