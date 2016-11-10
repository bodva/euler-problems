package main

import (
	"fmt"
)

type Fibo struct {
	a, b, current int
}

func main() {
	var limit int = 4.0e6
	var sum = 2
	fibo := []int{1, 1, 2}
	for fibo[2] < limit {
		fibo[0], fibo[1], fibo[2] = fibo[1], fibo[2], fibo[1]+fibo[2]
		if Even(fibo[2]) {
			sum += fibo[2]
		}
	}

	fmt.Println(fibo, sum)
}

func Even(number int) bool {
	return number%2 == 0
}
