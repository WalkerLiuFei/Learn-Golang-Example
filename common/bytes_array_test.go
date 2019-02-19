package main

import (
	"testing"
	"math/big"
	"github.com/ethereum/go-ethereum/common"
	"fmt"
)

func TestBytesArray(t *testing.T){
	byteArray := []byte{0,0,0,0,0,0,0,0,0,0,0,0,104}
	fmt.Println(new(big.Int).SetBytes(byteArray[:common.HashLength-common.AddressLength]).Cmp(big.NewInt(0)) == 0)
}
