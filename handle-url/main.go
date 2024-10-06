package main

import (
	"fmt"
	"net/url"
)

const myUrl string = "https://kartikmalik.tech/docs?title=nextjs&id=48kdkf4erkdkdrkww4r"

func main() {
	fmt.Println("Urls in goLang")

	//parsing the url to extract the values

	result, _ := url.Parse(myUrl)

	fmt.Println("this is the result", result)
	fmt.Println(result.Scheme)
	fmt.Println(result.Host)
	fmt.Println(result.Path)
	fmt.Println(result.RawQuery)

	qParams := result.Query()

	fmt.Println(qParams) // it gives you a map like this  map[id:[48kdkf4erkdkdrkww4r] title:[nextjs]]

	for _, val := range qParams {
		fmt.Println("the values are", val)
	}

	//create a url

	partsOrUrl := &url.URL{
		Scheme:  "https",
		Host:    "kartikmalik.tech",
		Path:    "/user",
		RawPath: "id=lkjhklj8ukj89y68",
	}

	anotherUrl := partsOrUrl.String()

	fmt.Println("The new url is", anotherUrl)

}
