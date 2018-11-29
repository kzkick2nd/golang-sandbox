package main

func main() {
	ch := make(chan int, 10)
	ch <- 100
	close(ch)
	close(ch)
	println(len(ch))
	println(cap(ch))
}
