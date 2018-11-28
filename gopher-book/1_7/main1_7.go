package main

type T struct{}

func (_ T) F() {}

func (_ *T) G() {}

func main() {
	var _ interface{ F() } = &T{} // 値型に定義されたメソッドはポインタ型にしても引き継ぐ
	var _ interface{ G() } = T{}  // ポインタ型に定義されたメソッドは値型には引き継がない
}
