package main

import (
	"fmt"
	"time"
)

func main() {
	presentTime := time.Now()
	fmt.Println(presentTime)

	fmt.Println(presentTime.Format("01-02-2006 15:04:05 Monday"))

	createDate := time.Date(2021, time.April, 10, 23, 23, 0, 0, time.UTC)
	fmt.Println("Created Date", createDate)
}
