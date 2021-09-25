package main

import (
	"fmt"
)

func main() {
	fmt.Println("Maps in golang")	

	languages := make(map[string]string)

	languages["JS"] = "JavaScript"
	languages["PY"] = "Python"
	languages["RB"] = "Ruby"

	fmt.Println("Lists of all languages",languages)
	fmt.Println("JS shorts for:", languages["JS"])

	delete(languages, "RB")
	fmt.Println("Lists of all languages",languages)
	

	//! Loops 
	for key, value := range languages {
		fmt.Printf("For key %v, value is %v\n", key, value)
	}
}
