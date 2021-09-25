package main

import "fmt"

const LogginToken = "sjaksjajksjka" //! Public

func main() {
	var username string = "lucifer"
	fmt.Println(username)
	fmt.Printf("variable is of type: %T \n", username)

	var isLoggedIn  bool = true
	fmt.Println(isLoggedIn)
	fmt.Printf("variable is of type: %T \n", isLoggedIn)

	var smallVal  uint16 = 255
	fmt.Println(smallVal)
	fmt.Printf("variable is of type: %T \n", smallVal)

	var smallFloat  float32 = 233.32323232323
	fmt.Println(smallFloat)
	fmt.Printf("variable is of type: %T \n", smallFloat)

	//! default values and some aliace
	var anothervariableint int
	fmt.Println(anothervariableint)
	fmt.Printf("variable is of type: %T \n", anothervariableint)

	var anothervariablestring string
	fmt.Println(anothervariablestring)
	fmt.Printf("variable is of type: %T \n", anothervariablestring)

	//! implicit type
	var website = "dhruvilgajjar.engineer"
	fmt.Println(website)

	//! no var style
	numberOfUser := 3000
	fmt.Println(numberOfUser)

}