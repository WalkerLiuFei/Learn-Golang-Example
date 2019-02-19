package main

import (
	"testing"
	"fmt"
)

func TestClosure(t *testing.T) {
	fmt.Println(closureFunc(0)())
	fmt.Println(closureFunc(1)())
}

func closureFunc(i int) func() int {
	return func() int {
		i++
		return i
	}
}
