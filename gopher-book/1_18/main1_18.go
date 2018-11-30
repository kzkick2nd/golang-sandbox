package main

import "fmt"

type Hoge struct {
	N int
}

func (h *Hoge) F() {
	fmt.Println(h)
}

func main() {
	var h *Hoge
	h.F()
}
