package main

import (
	"fmt"
	"math"
	"math/rand"
	_ "sort"
	"time"
)

func Add(a int, b int) int {
	sum := a + b
	return sum
}

func Circle(r float64) float64 {
	S := math.Pi * r * r
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

func Binary(target int) int {
	low, high := 0, 100

	for low <= high {
		mid := (low + high) / 2
		if mid == target {
			return mid
		}
		if mid < target {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}
	return -1
}

func random() int {
	rand.Seed(time.Now().UnixNano())
	r := rand.Intn(100)
	return r
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
	} else {
		fmt.Println("false")
	}

	RandomNum := random()
	fmt.Println(Binary(RandomNum))

}
