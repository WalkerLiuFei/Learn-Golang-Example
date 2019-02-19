package aes

import (
	"testing"
	"fmt"
	"io"
	 crand "crypto/rand"
)

func TestCTREncode(t *testing.T){
	var content string
	fmt.Print("input content:")
	fmt.Scanln(&content)

	var password string
	fmt.Print("input password:")
	fmt.Scanln(&password)


	salt := make([]byte, 32)
	_, err := io.ReadFull(crand.Reader, salt)
	if err != nil {
		panic("reading from crypto/rand failed: " + err.Error())
	}

	fmt.Printf("salt is : %s", salt)

	//derivedKey, err := scrypt.Key([]byte(password), salt, scryptN, scryptR, scryptP, scryptDKLen)
}