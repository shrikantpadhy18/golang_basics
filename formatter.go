package main

import "fmt"

func main() {
	x := "shrikant"
	y := "padhy"

	fmt.Print("The firstname is ", x, "The second name is\n ", y)

	fmt.Print("The firstname is %v and lastname is %v\n", x, y)
	fmt.Print("The firstname is %q and lastname is %q\n", x, y)

	var h = fmt.Sprintf("The firstname is %q and lastname is %q\n", x, y)
	fmt.Print(h)
}
