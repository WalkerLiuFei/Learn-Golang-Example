package main

import (
	"testing"
	"crypto/hmac"
	"crypto/sha256"
	"fmt"
	"github.com/ethereum/go-ethereum/common/hexutil"
)

func TestMac(t *testing.T){
	message := "1234567890"
	key := "abcde"
	mac := hmac.New(sha256.New, []byte(key))
	mac.Write([]byte(message))
	expectedMAC := mac.Sum(nil)
	fmt.Println(hexutil.Encode(expectedMAC))
	//hmac.Equal(messageMAC, expectedMAC)
}
