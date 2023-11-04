package main

import "fmt"

func main() {
	//x := 0

	/*for x < 5 {
		fmt.Println(x)
		x += 1
	}*/

	names := []string{"shrikant", "raja", "prashant", "chaubey"}
	for i := 0; i < 5; i++ {
		fmt.Println(i)
	}

	for i := 0; i < len(names); i++ {
		fmt.Println(names[i])
	}

	for index, value := range names {
		fmt.Println("The value at index %v is %v", index, value)
	}

}
