package runtimetest

import (
	"fmt"
	"runtime"
	"testing"
)

func say(s string) {
	for i := 0; i < 5; i++ {
		runtime.Gosched()
		fmt.Println(s)
	}
}

func TestGosched(t *testing.T) {
	runtime.GOMAXPROCS(1)

	go say("world")
	say("hello")
}
