package main

import "fmt"

type Key struct {
	V1 int
	V2 [2]int
}

func main() {
	var m map[Key]int
	key := Key{V1: 100, V2: [2]int{1, 2}}
	m[key] = 100
	fmt.Println(m[key])
}
