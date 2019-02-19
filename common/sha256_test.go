package main

import (
	"testing"
	"fmt"
	"crypto/sha256"
	"github.com/ethereum/go-ethereum/common/hexutil"
)

func TestSha256(t *testing.T){
	src := "0051337C281FADA1EDAE5EB3AFD827FBE56031755C"
	sha := sha256.New()

	_ ,err:= sha.Write([]byte(src))
	if err != nil{
		panic(err)
	}
	fmt.Printf("%x", sha.Sum(nil))
}

func TestSha256_2(t *testing.T){
	/* Create a new SHA256 context */

	sha256H := sha256.New()
	/* SHA256 Hash #1 */
	fmt.Println("5 - Perform SHA-256 hash on the extended RIPEMD-160 result")
	sha256H.Reset()
	sha256H.Write([]byte("51337C281FADA1EDAE5EB3AFD827FBE56031755C	"))
	hash1 := sha256H.Sum(nil)
	fmt.Println(hexutil.Encode(hash1))
	fmt.Println("=======================")
}