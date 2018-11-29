package main

func main() {
	var ch1 chan chan int
	select {
	case <-ch1:
		println("A")
	default:
		println("B")
	}
}
