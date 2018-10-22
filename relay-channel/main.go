package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"reflect"
)

func input(r io.Reader) chan string {
	ch1 := make(chan string)
	go func() {
		s := bufio.NewScanner(r)
		for s.Scan() {
			ch1 <- s.Text()
		}
		close(ch1)
	}()
	return ch1
}

func relay(recv chan string) string {
	r := <-recv
	fmt.Println(reflect.TypeOf(r))
	return r
}

func main() {
	ch := input(os.Stdin)
	for {
		fmt.Print(">")
		fmt.Println(relay(ch))
	}
}
