package main

import (
	"testing"

	"math/rand"
	"fmt"
	"container/heap"
	"sort"
)
type MockTx struct {
	nonce int32
	gasPrice int32
}
type Transactions []MockTx
type TxByNonce  Transactions
func (s TxByNonce) Len() int{return len(s)}
func (s TxByNonce) Less(i,j int) bool{return s[i].nonce > s[j].nonce}
func (s TxByNonce) Swap(i,j int) {s[i],s[j] = s[j],s[i]}

func (s *TxByNonce) Push(x interface{}) {
	*s = append(*s, x.(MockTx))
}

func (s *TxByNonce) Pop() interface{} {
	old := *s
	n := len(old)
	x := old[n-1]
	*s = old[0 : n-1]
	return x
}

func TestHeapUse(t *testing.T){
	txs := make(TxByNonce,10)
	for count := 10;count >= 0;count--{
		txs = append(txs,MockTx{
			nonce : rand.Int31n(1000),
			gasPrice : rand.Int31n(1000000),
		})
	}
	fmt.Println(txs)
	heap.Init(&txs)
	fmt.Println(txs)
	sort.Sort(txs)
	fmt.Println(txs)
}
