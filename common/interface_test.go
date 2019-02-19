package main

import (
	"reflect"
	"unsafe"
	"testing"
	"sync"
	"runtime"
	"sync/atomic"
	"time"
	"strconv"
)

func ExampleInterface() {

}

func TestSliceToArray(t *testing.T) {
	slice := make([]byte, 32)
	println(reflect.TypeOf(slice).Name())

	// Convert our destination slice to a byte buffer
	header := *(*reflect.SliceHeader)(unsafe.Pointer(&slice))
	header.Len *= 4
	header.Cap *= 4
	dataset := *(*[]byte)(unsafe.Pointer(&header))

	println(reflect.TypeOf(dataset).Name())
}

func TestMutiThreadRun(t *testing.T) {
	threads := runtime.NumCPU()
	var pend sync.WaitGroup
	println(threads)
	pend.Add(threads)

	var progress uint32
	for i := 0; i < threads; i++ {
		go func(id int) {
			defer pend.Done()
			time.Sleep(3 * time.Second)
			// Calculate the dataset segment
			atomic.AddUint32(&progress, 1)
			println("Id is " + strconv.Itoa(id) + " progress : " + strconv.Itoa(int(progress)))
		}(i)
	}
	time.Sleep(3 * time.Second)
}

func TestCommon(t *testing.T){
	println(0 ^ 1)
}
