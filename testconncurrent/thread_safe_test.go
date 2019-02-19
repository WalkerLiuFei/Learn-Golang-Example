package main

import (
	"testing"
	"fmt"
)

func TestSafe(t *testing.T) {

	//sliceChan := make(chan []int)

}

func modifySlice(sliceChan chan []int) {
	select {
	case slice := <-sliceChan:
		fmt.Println(slice)
	}
}
