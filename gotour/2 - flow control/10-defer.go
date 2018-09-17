package main

import "fmt"

func main() {
	// defer is a stack, that doesn't run until the surrounding
	// function returns
	defer fmt.Println("!")
	defer fmt.Println("world")

	fmt.Println("hello")
}
