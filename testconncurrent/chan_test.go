package main

import (
	"testing"
	"sync"
	"runtime"
	"time"
	"fmt"
	"os"
	"os/signal"
	"syscall"
)

func TestSingleDirect(t *testing.T) {

}

type receiver struct {
	sync.WaitGroup
	data chan int
}

func newReceiver() *receiver {
	r := &receiver{
		data: make(chan int),
	}
	r.Add(1)
	go func() {
		for d := range r.data {
			println(d)
		}
	}()
	return r
}
func TestFactoryPattern(t *testing.T) {
	r := newReceiver()
	r.data <- 1
	r.data <- 2
	close(r.data)
	r.Wait()
}

//信号量
func TestSemaphore(t *testing.T) {
	runtime.GOMAXPROCS(4)
	var wg sync.WaitGroup
	sem := make(chan struct{}, 2)
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			sem <- struct{}{}        // 获取信号量
			defer func() { <-sem }() //释放信号量
			time.Sleep(time.Second * 2)
			fmt.Println(id, time.Now())
		}(i)
	}
	wg.Wait()
}

// ticker

func TestTicker(t *testing.T) {
	go func() {
		for {
			select {
			case t := <-time.After(time.Second):
				fmt.Println(t)
			}
		}
	}()
	go func() {
		ticker := time.Tick(time.Second)
		count := 0
		for {
			select {
			case t := <-ticker:
				count++
				fmt.Println(t)
				if count > 10 {
					break
				}
			}
		}
	}()
	<-(chan struct{})(nil)
}

//signal wait
var exist = &struct {
	sync.RWMutex
	funcs   []func()
	signals chan os.Signal
}{}

func atexit(f func()) {
	exist.Lock()
	defer exist.Unlock()
	exist.funcs = append(exist.funcs, f)
}

func waitExist() {
	if exist.signals == nil {
		exist.signals = make(chan os.Signal)
		signal.Notify(exist.signals, syscall.SIGABRT, syscall.SIGINT)
	}
	for {
		select {
		case e := <-exist.signals:
			exist.RLock()
			fmt.Println(e.String())
			for _, f := range exist.funcs {
				 f()
			}
			exist.RUnlock()
			<- exist.signals
		}
	}
}

func TestSignalListen(t *testing.T) {
	exist.funcs = append(exist.funcs, func() {
		println("function 1 exist ")
	})

	exist.funcs = append(exist.funcs, func() {
		println("function 2 exist ")
	})
	waitExist()
}


func TestLeak(t *testing.T){
	c := make(chan int)
	for i := 0;i < 10; i++{
		go func() {
			<- c
		}()
	}
	for {
		time.Sleep(time.Second)
		runtime.GC()
	}
}
