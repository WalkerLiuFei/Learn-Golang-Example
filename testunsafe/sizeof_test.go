package testunsafe

import (
	"testing"
	"fmt"
	"unsafe"
)

func TestSizeOf(t *testing.T)  {
	var m map[int]int
	var p uintptr
	//从这里看出来,这里m其实是个指针
	fmt.Println(unsafe.Sizeof(m),unsafe.Sizeof(p))
}