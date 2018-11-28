package main

type Hoge struct{}

func (h *Hoge) A() string {
	return "Hoge"
}

func (h *Hoge) B() {
	println(h.A())
}

type Fuga struct {
	Hoge
}

func (f *Fuga) A() string {
	return "Fuga"
}

// func (f *Fuga) B() {
// 	println(f.A())
// }

func main() {
	var f Fuga
	f.B()
	// f.B()
	// f.Hoge.B()
}
