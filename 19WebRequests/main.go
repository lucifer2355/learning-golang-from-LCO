package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

const url = "https://lco.dev" 

func main() {
	fmt.Println("Web Request in golang")

	response, err := http.Get(url)
	checkNilErr(err)

	fmt.Printf("Response is type of %T\n", response)
	defer response.Body.Close()

	dataBytes, err := ioutil.ReadAll(response.Body)
	checkNilErr(err)

	fmt.Println("Response is:", string(dataBytes))
} 

func checkNilErr(err error) {
	if err != nil {
		panic(err)
	}
}