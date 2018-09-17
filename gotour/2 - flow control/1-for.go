package main

import "fmt"

func main() {
	sum := 0
	// any variable declared in the for loop has only the loop scope
	for i := 0; i < 10; i++ {
		sum += i
	}
	fmt.Println(sum)
}
