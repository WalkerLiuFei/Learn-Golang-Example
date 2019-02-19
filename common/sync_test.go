package main

import (
	"fmt"
	"testing"
	"sync"
	"time"
)

func TestRunOnce(t *testing.T){
	exector := new(sync.Once)
	go exector.Do(do)
	go exector.Do(do)
	time.Sleep(time.Second)
}

func do(){
	fmt.Println("done!")
}