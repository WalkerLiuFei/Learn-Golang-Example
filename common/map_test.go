package main

import (
	"testing"
	"fmt"
)



func TestMapPassToFunc(t *testing.T){
	mapper := make(map[int]int)
	fmt.Printf("%p \n",&mapper)
	fillUpMap(mapper)
	fmt.Println(mapper[20])
}

func fillUpMap(m map[int]int) {
	fmt.Printf("%p \n",&m)
	for count := 0;count < 100;count++{
		m[count] = count + 1
	}
}