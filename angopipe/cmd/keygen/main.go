package main

import (
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"io"
)

func main() {
	key := make([]byte, 32)
	io.ReadFull(rand.Reader, key)
	readbleKey := base64.StdEncoding.EncodeToString(key)
	fmt.Println(readbleKey)
}
