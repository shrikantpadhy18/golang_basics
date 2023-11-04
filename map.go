package main

import "fmt"

func main() {

	maps := map[string]string{
		"name":    "shrikant",
		"surname": "padhy",
	}

	fmt.Println(maps)

	for key, value := range maps {
		fmt.Println(key, "->", value)
	}

}
