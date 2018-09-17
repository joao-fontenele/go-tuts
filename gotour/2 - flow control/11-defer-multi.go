package main

import "fmt"

func main() {
	// defer is a stack, that doesn't run until the surrounding
	// function returns

	for i := 0; i < 10; i++ {
		defer fmt.Println(i)
	}

	fmt.Println("done")
}
