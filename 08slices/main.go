package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println("Welcome to video on slices")

	var fruitList = []string{"Apple", "Mango", "Banana"}
	fmt.Println(len(fruitList))

	//the append method

	fruitList = append(fruitList, "Test1", "Test2")
	fmt.Println(fruitList)

	//slice a slice or cut

	fruitList = append(fruitList[1:3], "Test3") // start from 1 and end on the index 2
	fmt.Println(fruitList)

	highScores := make([]int, 4)

	highScores[0] = 123
	highScores[1] = 124
	highScores[2] = 125
	highScores[3] = 126
	// highScores[4] = 127
	//slices with make will not allow you to create the extra value because it allows the only the memory slots that are defined but with append it handles
	// this smartly it appends values with append method and allocate new memory spaces to the variable values but by default it allocate only limited for performance

	highScores = append(highScores, 444, 855, 666)
	fmt.Println(highScores)

	sort.Ints(highScores) //sort the slice
	fmt.Println(highScores)
	fmt.Println(sort.IntsAreSorted(highScores)) // return true if the slice is sorted
}
