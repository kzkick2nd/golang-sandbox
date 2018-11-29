package main

import "fmt"

func main() {
	type A = int
	// type A int
	var a A
	fmt.Printf("%T", a)
}
