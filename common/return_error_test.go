package main

import (
	"testing"
	"time"
	"fmt"
)
// MyError is an error implementation that includes a time and message.
type MyError struct {
	When time.Time
	What string
}

func (e MyError) Error() string {
	return fmt.Sprintf("%v: %v", e.When, e.What)
}
func TestReturn(t *testing.T) {
	retR ,usedGasR ,failedR,errR := call()
	fmt.Println(retR)
	fmt.Println(usedGasR)
	fmt.Println(failedR)
	fmt.Println(errR)
}
func call()(ret []byte, usedGas uint64, failed bool, err error){
	if err = subCall(); err != nil {
		return
	}
	return nil,0,false,nil
}
func subCall() error{
	return MyError{
		What:"my error",
		When:time.Now(),
	}
}