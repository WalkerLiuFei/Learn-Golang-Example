package ethereum

import (
	"testing"
	"io/ioutil"
	"github.com/ethereum/go-ethereum/rlp"
	"github.com/ethereum/go-ethereum/core/types"
	"bytes"
	"fmt"
	"github.com/ethereum/go-ethereum/common/hexutil"
)

func TestDecodeRLP(t *testing.T) {
	rlpFile, err := ioutil.ReadFile("E:\\GOPATH\\src\\MyTest\\ethereum\\transactions.rlp")
	if err != nil {
		panic(err)
	}
	array := []*types.Transaction{}
	rlp.Decode(bytes.NewReader(rlpFile), array)
	fmt.Println(array)
}

func TestRLPEncoding1024(t *testing.T) {
	str := ""
	for len(str) < 1024 {
		str = str + "1"
	}
	data, err := rlp.EncodeToBytes(str)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s \n", hexutil.Encode(data[:3]))
	fmt.Printf("%x",1024)
}
