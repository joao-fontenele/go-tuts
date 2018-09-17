package main

import "fmt"

func main() {
	sum := 1
	// init and post statements are optional
	// by removing them, we get a "while" loop
	// for ; sum < 1000;  { // would still work
	for sum < 1000 {
		sum += sum
	}
	fmt.Println(sum)
}
