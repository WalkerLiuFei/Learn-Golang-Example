package main

import (
	"testing"
	"math/big"
	"fmt"
	"github.com/ethereum/go-ethereum/common/math"
	"github.com/ethereum/go-ethereum/core/vm"
)

func TestBigInt1( t *testing.T)  {
	newInt := big.NewInt(100)
	fmt.Println(math.U256(newInt).Int64())
	fmt.Println(newInt.Text(10))
}

func TestBigInt2(t *testing.T)  {
	i := big.NewInt(0)
	memory := new(vm.Memory)
	mStart, val := big.NewInt(0x40), big.NewInt(0x8000000000000)
	memory.Resize(0x40 + 32)
	memory.Set32(mStart.Uint64(), val)
	cpy := memory.Get(mStart.Int64(),32)
	i.SetBytes(cpy)
	fmt.Printf("%x",i.Int64())
}


func Test2(t *testing.T){
	pow := math.BigPow(10,18)
	s := "ad78ebc5ac6200000"
	data := big.NewInt(0)
	data.SetString(s,16)
	fmt.Println(data.Text(10))
	data = data.Div(data,pow)
	fmt.Println(data.Text(10))
}