package main

import "fmt"

func main() {
	// var ptr *string

	// fmt.Println(ptr)
	myNumber := 23
	var ptr = &myNumber

	fmt.Println("Valur of actual pointer is", ptr)
	fmt.Println("Value of actual pointer", *ptr)

	*ptr = *ptr + 2

	fmt.Println("The new value after adding", *ptr)
}
