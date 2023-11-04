package main

import "fmt"

func main() {
	var shrikant = "shrikant"
	var hi string = "padhy"
	var hello string

	fmt.Println(shrikant, hi, hello)

	hello = "galileo"
	fmt.Println(shrikant, hi, hello)

	variable3 := "padhy" //rhis syntax wont work outside function
	fmt.Println(variable3)

	//int

	var age int = 30
	fmt.Print(age)

	var ag int8 = 12

	var score float32 = 23.23
	var res float64 = 56

	z := 1.5 //its by default float64
	fmt.Print(ag, score, res, z)

}
