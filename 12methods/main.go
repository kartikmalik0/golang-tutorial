package main

import "fmt"

func main() {
	fmt.Println("Structs in golang")
	//there is no inheritance in golang
	//  no super or parent concept in golang

	kartik := User{"Kartik", "kartik@go.dev", true, 20}

	fmt.Println(kartik)
	fmt.Printf("Kartik details are : %+v\n", kartik)
	fmt.Printf("User name is : %v\n", kartik.Name)
	kartik.GetStatus()
}

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}

func (u User) GetStatus() {
	fmt.Println("The user status is", u.Status)
}
