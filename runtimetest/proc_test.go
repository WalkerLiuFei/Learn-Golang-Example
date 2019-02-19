package runtimetest

import (
	"testing"
	"runtime"
	"time"
	"fmt"
	"math/big"
)

func TestGetPcPointer(t *testing.T){
	for count,_ := big.NewInt(0).SetString("10000000",10);count.Cmp(big.NewInt(0)) > 0 ;{
		go test()
		count.Sub(count,big.NewInt(1))
		numGo := runtime.NumGoroutine()
		fmt.Println(numGo)
	}

}
func test(){
	time.Sleep(time.Second * 100)
}
