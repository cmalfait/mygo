package main

import "fmt"

func myshit(murphy int) int {
	fmt.Printf("you sent us: %d\n", murphy)
	return murphy + 1
}

func main() {
	junk := myshit(5)
	junk1 := myshit(10)
	fmt.Printf("junk: %d junk2: %d\n", junk, junk1)
}
