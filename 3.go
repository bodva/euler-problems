package main

import (
	"fmt"
	"math"
)

func main() {
	var a uint = 13195
	var b uint = 600851475143
	// fmt.Println(a)
	// fmt.Println(b, math.Sqrt(float64(b)))
	fmt.Println(PrimeFactor(a))
	fmt.Println(PrimeFactor(b))

}
func PrimeFactor(value uint) uint {
	for i := uint(math.Sqrt(float64(value))); i > 1; i-- {
		if value%i == 0 && IsPrime(i) {
			return i
		}
	}

	return 0
}
func IsPrime(value uint) bool {
	for i := uint(math.Sqrt(float64(value))); i > 1; i-- {
		if value%i == 0 {
			return false
		}
	}

	return true
}
