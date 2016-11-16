package main

import (
	"fmt"
)

func main() {
	var n = 20
	for i := 1; i < 60949320; i++ {
		c := make(chan int, 100)
		go func() {
			if valid(i, n) {
				fmt.Println(i)
			}
			c <- -1
		}()
		<-c
	}
	fmt.Println(n)
}

func check(i, n int) {
	if valid(i, n) {
		fmt.Println(i)
	}
}

func valid(current, n int) bool {
	for i := 1; i < n+1; i++ {
		if current%i != 0 {
			return false
		}
	}

	return true
}
