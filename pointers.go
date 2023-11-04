package main

import "fmt"

func greet(n *string) {
	*n = "shrikant"
}

func main() {
	var res string = "padhy"
	name := &res
	var k *string = &res
	fmt.Println(k)
	greet(name)
	fmt.Println(res)
}
