package main

import "fmt"

func swap(s1, s2 string) (string, string) {
	return s2, s1
}

func main() {
	fmt.Println(swap("world", "hello"))
}
