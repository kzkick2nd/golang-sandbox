package main

import "fmt"

// type Hoge struct{
// 	name int
// }

func main() {
	type Hoge struct{ int }
	var h Hoge
	fmt.Println(h)
	fmt.Println(h.num)
}
