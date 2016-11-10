package main

import (
	"fmt"
	"strconv"
)

func main() {
	p := search()
	// fmt.Println(palindom(9009))
	// fmt.Println(p)
	max := p[0]
	for _, e := range p {
		if e > max {
			max = e
		}
	}
	fmt.Println(max)
}

func search() []int {
	var p []int
	for i := 999; i > 99; i-- {
		for j := 999; j > 99; j-- {
			// fmt.Println(i, j)
			m := i * j
			if palindom(m) {
				p = append(p, m)
			}
		}
	}

	return p
}

func palindom(number int) bool {
	str := strconv.Itoa(number)
	// fmt.Println(str)
	b := []byte(str)
	// fmt.Println(b, len(b))
	if len(b)%2 != 0 {
		return false
	}

	for i := 0; i < len(b)/2; i++ {
		j := len(b) - i - 1
		if b[i] != b[j] {
			return false
		}
	}

	return true
}
