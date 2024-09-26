package main

import "fmt"

func main() {
	fmt.Println("Maps in golang")

	lan := make(map[string]string)

	lan["JS"] = "JavaScript"
	lan["TS"] = "TypeScript"

	fmt.Println(lan)

	//delete from the map

	// delete(lan, "TS")
	// fmt.Println(lan)

	//loops in map

	for key, value := range lan {
		fmt.Printf("For key %v, Value is %v\n", key, value)
	}

}
