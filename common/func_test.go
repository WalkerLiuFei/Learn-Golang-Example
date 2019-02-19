package main

import "testing"

func TestFuncCall(t *testing.T) {
	fun := function{
		myfunc:getFunc(10),
	}
	println(fun.myfunc(1,2))
}

type MyTypeFunc func(num1,num2 int) int
type function struct{
	myfunc MyTypeFunc
}
func getFunc(num int)MyTypeFunc{
	return func(num1,num2 int) int{
		return num1 + num2
	}
}