package main

import (
	"testing"
	"time"
	"fmt"
)

var intChan = make(chan int,10)
func TestChan(t *testing.T){
	go readChan()
	for count := 0;count < 10;count++{
		select {
			case  intChan <- count:
		}
	}
	fmt.Println("finish send")
	for len(intChan) > 0{}

}

func readChan(){
	for true{
		time.Sleep(time.Second)
		select {
		case msg := <-intChan:
			fmt.Println(msg)
		default:
			fmt.Println("no message received")
		}
	}

}
