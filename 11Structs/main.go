package main

import "fmt"

//! No inheritance in golang; No super or parent
func main() {
	fmt.Println("Welcome to structs in golang")

	dhruvil := User{"Dhruvil", "dhruvil@go.dev", true, 22}
	fmt.Println(dhruvil)
	fmt.Printf("Dhruvil details are %+v\n", dhruvil)
	fmt.Printf("Name is %v and Email id %v\n", dhruvil.Name, dhruvil.Email)
}

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}
