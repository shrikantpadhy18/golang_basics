package main

import (
	"fmt"
	"sort"
	"strings"
)

func main() {
	var x = "Hello Mumbai"
	fmt.Println(strings.Contains(x, "Hello"))

	fmt.Println(strings.ReplaceAll(x, "Hello", "Hi"))
	fmt.Println(x)

	var age = []int{2, 13, 34, 56, 8, 7, 7}

	sort.Ints(age)
	fmt.Println(age)

	fmt.Println(sort.SearchInts(age, 7))

	name := []string{"shrikant", "padhy", "amit", "saurabh", "kenya"}
	sort.Strings(name)

	fmt.Print(name)
}
