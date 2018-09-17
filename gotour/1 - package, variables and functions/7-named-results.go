package main

import "fmt"

// we named the return variables, and we don't explicitly return them below
func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return
}

func main() {
	fmt.Println(split(27))
}
