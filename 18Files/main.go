package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func main() {
	fmt.Println("Welcome to FILES in golang")
	content := "This need to go in file"

	file, err := os.Create("./mylcogofile.txt")
	checkNilErr(err)

	length, err := io.WriteString(file, content)
	checkNilErr(err)

	fmt.Println("length is:", length)
	defer file.Close()

	readFile("./mylcogofile.txt")
}

func readFile(fileName string) {
	dataByte, err := ioutil.ReadFile(fileName)
	checkNilErr(err)

	fmt.Println("Tex data inside the file is:", string(dataByte))
}

func checkNilErr(err error) {
	if err != nil {
		panic(err)
	}
}