package main

import (
	"fmt"
	"os"
	"syscall"
	"time"
)

func main() {
	f, _ := os.Stat("test")
	t := f.Sys().(*syscall.Stat_t).Birthtimespec
	d := time.Unix(t.Sec, t.Nsec)
	fmt.Println(d)
}
