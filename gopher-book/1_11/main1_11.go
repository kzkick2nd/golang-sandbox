package main

func main() {
	ch := make(chan int, 10)
	println(len(ch))
	println(cap(ch))
}
