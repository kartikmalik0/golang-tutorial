package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	fmt.Println("Welcome to the files in the golang")
	content := "This is the new file created by golang"

	file, err := os.Create("./demo.txt")

	if err != nil {
		panic(err)
	}

	lenght, err := io.WriteString(file, content)

	if err != nil {
		panic(err)
	}

	fmt.Println("Length is", lenght)
	defer file.Close()

	ReadFile("./demo.txt")
}

func ReadFile(filename string) {
	databytes, err := os.ReadFile(filename)

	if err != nil {
		panic(err)
	}

	// fmt.Println("Here are the data bytes", databytes) to convert this into string we need to wrap into the string() method

	fmt.Println("Here are the string of databytes", string(databytes))
}
