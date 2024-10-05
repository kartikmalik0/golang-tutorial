package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

const url = "https://kartikmalik.tech"

func main() {
	fmt.Println("")
	res, err := http.Get(url)

	if err != nil {
		panic(err)
	}

	fmt.Printf("type of response %T\n", res)

	// res.Body.Close()

	dataBytes, err := ioutil.ReadAll(res.Body)

	if err != nil {
		panic(err)
	}

	htmlContent := string(dataBytes)
	fmt.Println(htmlContent)
}
