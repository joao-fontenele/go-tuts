package main

import (
	"fmt"
	"math"
)

func pow(x, n, lim float64) float64 {
	// if can use the init statement just like for loops, and, the variables declared
	// in it have only the if's scope
	if v := math.Pow(x, n); v < lim {
		return v
	}

	// doesn't exist outside if scope
	// fmt.Println(v)

	return lim
}

func main() {
	fmt.Println(
		pow(3, 2, 10),
		pow(3, 3, 20),
	)
}
