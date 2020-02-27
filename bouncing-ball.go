package main

import (
	"fmt"
	"math"
)

func BouncingBall(h, bounce, window float64) int {
	seen := 2
	for n := 2; n <= int(h); n++ {
		seen++
		h = math.Ceil(h * bounce)
		fmt.Println("H", h)
		if h < window {
			break
		}
	}

	return seen
}

func main() {
	fmt.Println(BouncingBall(6, 0.66, 1.5))
}
