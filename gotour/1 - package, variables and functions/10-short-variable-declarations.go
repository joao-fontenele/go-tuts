package main

import "fmt"

// in the package scope, the := contruct is not available, but it is in functions

func main() {
	var i, j int = 1, 2
	k := 3
	c, python, java := true, false, "no!"

	fmt.Println(i, j, k, c, python, java)
}
