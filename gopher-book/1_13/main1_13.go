package main

func main() {
	ch1 := make(chan int)
	ch2 := make(chan int)
	go func() {
		<-ch1
		print("A")
		close(ch2)
	}()
	close(ch1)
	<-ch2
	print("B")
}
