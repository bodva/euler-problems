package main

import (
	"fmt"
)

func main() {
	fmt.Println(recursive(0, 3, 1000))
}

func recursive(sum, iteration, n int) int {
	if iteration < n {
		sum += div3or5(iteration)
		return recursive(sum, iteration+1, n)
	} else {
		return sum
	}
}

func div3or5(value int) int {
	if bool(value%3 == 0) || bool(value%5 == 0) {
		return value
	} else {
		return 0
	}
}
