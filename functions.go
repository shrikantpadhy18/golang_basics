package main

import (
	"fmt"
	"math"
)

func greeting(n string) {
	fmt.Printf("The value of n is %v\n", n)
}

func greetingone(n string) {
	fmt.Printf("The value of n is %v", n)
}

func cycle(n []string, f func(string)) {
	for i := 0; i < len(n); i++ {
		var data = n[i]
		f(data)

	}
}

func calculateArea(r float64) float64 {
	return math.Pi * r * r
}

func main() {
	//greeting("shrikant")
	//greetingone("padhy")
	cycle([]string{"shrikant", "padhy", "trilochan"}, greeting)
	fmt.Printf("area is %0.3f ", calculateArea(5))
	//fmt.Printf("hello")
}
