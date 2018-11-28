package main

import "fmt"

func f(ns []int) {
	ns[1] = 200             // 背後の配列にポインタ経由で代入
	ns = append(ns, 40, 50) // capを超えたら配列再確保。元配列には影響しない。
	fmt.Println("関数内", ns)
}

func main() {
	ns := []int{10, 20, 30}
	fmt.Println("初期化", ns)
	f(ns)
	fmt.Println("main", ns)
}
