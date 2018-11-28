package main

import "fmt"

func main() {
	ns := []int{10, 20, 30}
	defer fmt.Println(ns)
	ns[1] = 200
	ns = append(ns, 40, 50)
}
