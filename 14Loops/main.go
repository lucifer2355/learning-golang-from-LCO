package main

import "fmt"

func main() {
	fmt.Println("Welcome to LOOPS in golang")

	days := []string{"Sunday", "Tuesday", "Wednesday", "Friday", "Saturday"}

	for d := 0; d < len(days); d++ {
		fmt.Println(days[d])
	}

	for i := range days {
		fmt.Println(days[i])
	}

	for index, day := range days {
		fmt.Printf("index is %v and value is %v\n", index, day)
	}

	rougeValue := 1

	for rougeValue < 10 {
		if rougeValue == 2 {
			goto lco
		}

		if rougeValue == 5 {
			// break
			rougeValue++
			continue
		}

		fmt.Println("Value is: ", rougeValue)
		rougeValue++
	}

lco:
	fmt.Println("Jumping to learncodeonline.in")
}
