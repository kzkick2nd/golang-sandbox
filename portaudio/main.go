package main

import (
	"github.com/gordonklaus/portaudio"
	"math/rand"
	"time"
)

func main() {
	portaudio.Initialize()
	defer portaudio.Terminate()

	h, err := portaudio.DefaultHostApi()
	if err != nil {
		panic(err)
	}

	stream, err := portaudio.OpenStream(portaudio.HighLatencyParameters(nil, h.DefaultOutputDevice), func(out []int32) {
		for i := range out {
			out[i] = int32(rand.Uint32())
		}
	})
	if err != nil {
		panic(err)
	}
	defer stream.Close()

	stream.Start()
	time.Sleep(3 * time.Second)
	stream.Stop()

}
