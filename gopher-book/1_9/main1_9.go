package main

import "fmt"

func main() {
	ns := [...]int{10, 20, 30}
	ptr := &ns
	fmt.Println(len(ptr))
}
