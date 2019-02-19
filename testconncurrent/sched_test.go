package main

import (
	"testing"
	"runtime"
	"fmt"
	"time"
)

func TestSched(t *testing.T) {
	runtime.GOMAXPROCS(1)
	dataChan := make(chan int)
	for count := 10; count >= 0; count-- {
		go f(dataChan)
	}

	for count := 10; count >= 0; count-- {
		dataChan <- count
	}
	time.Sleep(1)
}

func f(dataChan chan int) {
	select {
		case data := <- dataChan:
			if data % 2 == 0{
				runtime.Gosched()
			}
			fmt.Println(data)
	}
}


