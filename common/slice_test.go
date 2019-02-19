package main

import (
	"testing"
	"fmt"
	//"reflect"
	"unsafe"
	"reflect"
)


func TestPointerToSlice(t *testing.T){

	//s := make([]byte, 200)
	//ptr := unsafe.Pointer(&s[0])

	backingArray := [10]byte{1,2,3,4,5,6,7,8,9,10}
	ptr := unsafe.Pointer(&backingArray[0])


	s := ((*[10]byte)(ptr))[:10]
	fmt.Println(s)

	var s1 = struct {
		addr uintptr
		len int
		cap int
	}{uintptr(ptr), 10, 10}
	s2 := *(*[]byte)(unsafe.Pointer(&s1))
	fmt.Println(s2)

	var o []byte
	sliceHeader := (*reflect.SliceHeader)((unsafe.Pointer(&o)))
	sliceHeader.Cap = 10
	sliceHeader.Len = 10
	sliceHeader.Data = uintptr(ptr)
	fmt.Println(o)

}
func TestSlicePass(t *testing.T){
		slice := []int{1,2,3,4,5,6}
		y := slice[1:3]

		fmt.Println(len(y))
		fmt.Println(y)
		//fmt.Println(y[3]) //error
		y = y[0:5]
		fmt.Println(len(y))
		fmt.Println(y)
		fmt.Println(y[3])
	/*fmt.Println(reflect.TypeOf(slice))
	modifySlice(&slice)
	fmt.Println(slice)*/
}

func modifySlice(slice *[]int){
	/*pointer := reflect.ValueOf((*slice)[1]).Pointer()
	fmt.Println(pointer)*/
}

func TestSlicePassFunc(t *testing.T)  {
	slice := make([]int,10)
	changeSlice(slice)
	fmt.Printf("%p \n",&slice)
	fmt.Println(slice[0])
}
func fillUpSlice(slice []int){
	fmt.Printf("%p \n",&slice)
	for count := 0 ;count < len(slice);count ++{
		slice[count] = 1
	}
}

func changeSlice(slice []int)  {
	slice = make([]int,40)
	slice[0] = 2
}