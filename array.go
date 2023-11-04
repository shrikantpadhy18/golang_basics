package main

import "fmt"

func main() {
	var arr = [3]int{10, 20, 30}
	fmt.Printf("%v\n", arr)

	fmt.Println("v", len(arr))

	//slices

	var arr2 = []int{10, 12, 13}
	arr2[2] = 25
	fmt.Print(arr2)
	arr2 = append(arr2, 25)

	var data = arr2[1:2]
	fmt.Println(data)
	fmt.Print(arr2)
}
