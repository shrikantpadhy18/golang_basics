package main

import "fmt"

type bill struct {
	name     string
	amount   float64
	validity bool
}

func getBill() bill {
	x := bill{
		name:     "shrikant",
		amount:   23.67,
		validity: true,
	}

	return x
}

func main() {
	obj := getBill()
	fmt.Println(obj.name)
}
