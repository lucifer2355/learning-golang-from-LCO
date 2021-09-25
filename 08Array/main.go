package main

import "fmt"

func main() {
	fmt.Println("Welcome to array in golang")

	var fruitsList [4]string

	fruitsList[0] = "Apple"
	fruitsList[1] = "Tomato"
	fruitsList[3] = "Peach"

	fmt.Println("Fruit list is: ", fruitsList)
	fmt.Println("Fruit list length: ", len(fruitsList))

	var vegList = [3]string{"Potato", "Beans", "Mushroom"}

	fmt.Println("Vegy list is: ", vegList)
	fmt.Println("Vegy list length: ", len(vegList))
}