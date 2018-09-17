package main

import (
	"fmt"
	"math"
)

func Sqrt(x float64) (result float64) {
	lim := 1e-8  // how accurate we want the result to be
	result = 1.0 // initial guess for the sqrt of x

	// while the result is not accurate enough, do
	for math.Abs(result*result-x) > lim {
		result -= (result*result - x) / (2 * result)
	}

	return
}

func main() {
	fmt.Println(Sqrt(4))
}
