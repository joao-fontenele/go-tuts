package main

import "fmt"

// package level scope
// variable declaration, in this case, all are boolean
var c, python, java bool

func main() {
	// function level scope
	var i int

	fmt.Println(i, c, python, java)
}
