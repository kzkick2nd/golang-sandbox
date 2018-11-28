package main

import "fmt"

type Hoge interface {
	F() int
}

type Piyo struct {
	N int
}

func (p *Piyo) F() int {
	p.Inc()
	return p.N
}

func (p Piyo) Inc() {
	p.N++
}

func main() {
	var h Hoge = Piyo{N: 1000}
	fmt.Println(h.F(), h.F())
}
