package main

import (
	"fmt"
	"math"
)

func main() {
	// all exported names of a package start with a capital letter, if it's not
	// in capital letter, it's not accessible from outside the package
	fmt.Println(math.Pi) // works
	// fmt.Println(math.pi) // doesnt work
}
