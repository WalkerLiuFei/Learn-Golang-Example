package main

import (
	"testing"
	"time"
	"fmt"
)

func TestTicker(t *testing.T){
	report := time.NewTicker(8 * time.Second)

	for {
		select {
		case <-report.C:
			println(time.Now().Unix())
		}
	}
}

// 实现生产消费模型非常简单！
func TestRangeTicker(t *testing.T){
	msgChan := make(chan string)
	go func(){
		for  t := range time.Tick(1 * time.Second){
			msgChan <- " from ticker 1 " + t.String()
		}
	}()

	go func(){
		for  t := range time.Tick(2 * time.Second){
			msgChan <- " from ticker 2 " + t.String()
		}
	}()

	for {
		select {
		case str := <- msgChan:
			fmt.Println(str)
		}
	}
}
