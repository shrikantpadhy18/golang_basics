package main

import (
	"fmt"
)

type bill struct {
	name   string
	amount float64
}

func (b bill) getBill() string {
	x := fmt.Sprintln("The name of the bill :", b.name, " and amount :", b.amount)
	return x
}

func main() {
	billobj := bill{
		name:   "balancesheet",
		amount: 245.78,
	}
	response := billobj.getBill()
	fmt.Println(response)
}
