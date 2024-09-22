package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Welcome to our Caffe")
	fmt.Println("Please rate our Coffe between 1 and 5")

	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')

	fmt.Println("Thanks for rating ", input)

	numRating, err := strconv.ParseFloat(strings.TrimSpace(input), 64)

	if(err != nil){
		fmt.Println(err)
	} else{
		fmt.Println("Added 1 to ur rating", numRating + 1)
	}
}
