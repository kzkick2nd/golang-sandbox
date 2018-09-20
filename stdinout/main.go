package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"time"
)

func main() {
	fmt.Println("Enter a string")

	ch := input(os.Stdin)
	// ループして時間経過したら抜ける
	// 問題を出すのと正答数を積み上げる
	for {
		select {
		case <-time.After(3 * time.Second):
			fmt.Println("Time Over")
			return
		case answer := <-ch:
			fmt.Println(answer)
		}
	}

}

func input(r io.Reader) <-chan string {
	// 受取りルーチンを初期化
	ch := make(chan string)
	go func() {
		s := bufio.NewScanner(r)
		for s.Scan() {
			ch <- s.Text()
		}
		close(ch)
	}()
	return ch
}
