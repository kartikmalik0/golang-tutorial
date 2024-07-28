package main

import "fmt"


const LoginToken string = "login94t9tyokentkds95093kasdfk"

func main() {
	var username string = "kartik"
	fmt.Println(username)
	fmt.Printf("Variable is the type of : %T \n", username)

	var isLoggedIn bool = true
	fmt.Println(isLoggedIn)
	fmt.Printf("Is user logged in: %T \n", isLoggedIn)

	var smallInt int = 255
	fmt.Println(smallInt)
	fmt.Printf("The type of smallInt is: %T \n", smallInt)

	var floatInt float32 = 255.398493874
	fmt.Println(floatInt)
	fmt.Printf("The type of floatInt is: %T \n", floatInt)

	var newVariable int
	fmt.Println((newVariable))

	var newVariableString string
	fmt.Println(newVariableString)

	//implisit type

	var implishit = "it's a string without type declaration"
	fmt.Println(implishit)

	//no var keyword

	numberOfUsers := 5656  // you cannot use this volouras operator outside the method like global scope or anything
	fmt.Println(numberOfUsers)

	fmt.Println(LoginToken)
}


