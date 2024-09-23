package main

import "fmt"

func main() {
	fmt.Println("Welcome to array in golang")

	var fruitList [3]string

	fruitList[0] = "apple"
	fruitList[1] = "apple2"
	fruitList[2] = "apple3"

	fmt.Println("fruit list is :", fruitList)
	fmt.Println("fruit list length is :", len(fruitList))

	var numList = [3]int{1, 2, 3}

	fmt.Println("this is the numList", numList)
}
