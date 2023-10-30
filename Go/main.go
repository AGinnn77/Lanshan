package main

import (
	"fmt"
	"math"
	_ "math"
)

func Add(a int, b int) int {
	sum := a + b
	return sum
}

func Circle(r float64) float64 {
	S := float64(math.Pi) * r * r
	return S
}

func PrimeNumber(n int) bool {
	if n <= 1 {
		return false
	}

	for i := 2; i < n; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}
func main() {
	var a, b int
	fmt.Scanln(&a, &b)
	fmt.Println(Add(a, b))

	var r float64
	fmt.Scanln(&r)
	fmt.Println(Circle(r))

	var Pri int
	fmt.Scanln(&Pri)
	if PrimeNumber(Pri) == true {
		fmt.Println("true")
	}
	if PrimeNumber(Pri) == false {
		fmt.Println("false")
	}
}
