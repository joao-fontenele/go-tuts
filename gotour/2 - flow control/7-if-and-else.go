package main

import (
	"fmt"
	"math"
)

func pow(x, n, lim float64) float64 {
	if v := math.Pow(x, n); v < lim {
		return v
	} else {
		// notice v is in the scope for the else statement
		fmt.Printf("%g >= %g\n", v, lim)
	}

	// v is out of scope here
	return lim
}

func main() {
	fmt.Println(
		pow(3, 3, 20),
		pow(3, 2, 10),
	)
}
