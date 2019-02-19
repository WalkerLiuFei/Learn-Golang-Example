package main

import (
	"testing"
	"reflect"
	"fmt"
)
type S struct{}

type T struct{
	S                  // 匿名嵌入字段
}

func(S)sVal()  {
	fmt.Println("123")
}
func(*S)sPtr() {
	fmt.Println("123")
}
func(T)tVal()  {
	fmt.Println("123")
}
func(*T)tPtr() {	fmt.Println("123")}
func methodGet(a interface{}){
	t := reflect.TypeOf(a)
	fmt.Println(t.PkgPath(),t.NumMethod())
	for i,n := 0,t.NumMethod();i < n;i++{
		fmt.Println(t.Method(i).Type.Name(),t.Method(i).PkgPath,t.Method(i).Name)
	}
}

func TestCallMethod(t *testing.T){
	var t1 T
	methodGet(t1)
	methodGet(&t1)
}

type ST struct {}
func TestKindAndType(t *testing.T){
	var i int = 99
	tp  := reflect.TypeOf(i)
	kind := reflect.Kind(i)
	fmt.Println(tp.Name(),tp.Kind(),kind.String())
	fmt.Println(reflect.TypeOf(map[string]int{}).Elem())
	fmt.Println(reflect.TypeOf([]int32{}).Elem())

}
