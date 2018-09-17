package main

import "fmt"

// since x and y are both int, we can declare a single int
func add(x, y int) int {
	return x + y
}

func main() {
	fmt.Println(add(42, 13))
}
