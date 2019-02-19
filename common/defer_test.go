package main

import (
	"testing"
	"fmt"
)

func TestDefer1(t *testing.T) {
	fmt.Println(test1())
	fmt.Println(test2())
	fmt.Println(test3())

	fmt.Println("----------------------")
	fmt.Println(test1Equal())
	fmt.Println(test2Equal())
	fmt.Println(test3Equal())
}

func test1Equal()(result int){
	result = 0  //return语句不是一条原子调用，return xxx其实是赋值＋ret指令
	func() { //defer被插入到return之前执行，也就是赋返回值和ret指令之间
		result++
	}()
	return
}

func test2Equal()(result int){
	t := 5
	result = t //赋值指令
	func(){
		t = t + 5
	}()
	return
}

func test3Equal()(r int){
	r = 1  //给返回值赋值
	func(r int) {        //这里改的r是传值传进去的r，不会改变要返回的那个r值
		r = r + 5
	}(r)
	return        //空的return
}
func test1() (result int) {
	defer func() {
		result++
	}()
	return 0
}

func test2() (r int) {
	t := 5
	defer func() {
		t = t + 5
	}()
	return t
}

func test3() (r int) {
	defer func(r int) {
		r = r + 5
	}(r)
	return 1
}
