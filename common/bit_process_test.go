package main

import (
	"testing"
	"fmt"
)

func TestLeft(t *testing.T){
	fmt.Println(4 << (^uint64(0) >> 63))
}