package main

import "fmt"

func main() {
	result := adder(2, 4)

	fmt.Println("The sum is", result)
}

func adder(val int, val1 int) int {
	return val + val1
}
