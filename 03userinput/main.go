package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv" // package to convert string to integere
	"strings" // strings package to trim the input
)



// strings.TrimSpace(input): This removes any extra whitespace (including newline characters) from the input string, so strconv.Atoi can properly convert the string to an integer.



func main() {
	welcome := "Welcome to user input"
	fmt.Println(welcome)

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter the number")

	// comma ok syntax // err ok syntax

	input , _ := reader.ReadString('\n')
	fmt.Println("Thanks for enter the number" ,input)

}

// sum of two numbers program

func main1() {
	reader := bufio.NewReader(os.Stdin)

	// Prompt for the first number
	fmt.Println("Enter first number:")
	input1, _ := reader.ReadString('\n')
	num1, _ := strconv.Atoi(strings.TrimSpace(input1)) // Trim the newline and convert to int

	// Prompt for the second number
	fmt.Println("Enter second number:")
	input2, _ := reader.ReadString('\n')
	num2, _ := strconv.Atoi(strings.TrimSpace(input2)) // Trim the newline and convert to int

	// Calculate the sum
	sum := num1 + num2
	fmt.Println("The sum is:", sum)
}
