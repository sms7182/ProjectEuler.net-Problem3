package main

import (
	"fmt"
	"math"
)

func main() {
	i := 3
	n := 600851475143
	max := 0
	c := math.Sqrt(float64(n))
	for i <= int(c) {
		if n%i == 0 {
			if i > max {
				max = i
			}

		}
		for n%i == 0 {
			n /= i
		}
		i = i + 2
	}
	fmt.Println(max)

}
