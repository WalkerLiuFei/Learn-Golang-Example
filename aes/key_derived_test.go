package aes

import (
	//"crypto/sha256"
	"crypto/rand"
	//"golang.org/x/crypto/pbkdf2"
	"golang.org/x/crypto/pbkdf2"
	"crypto/sha256"
	"testing"
	"crypto/aes"
	"crypto/cipher"
	"encoding/hex"
	"fmt"
	"github.com/ethereum/go-ethereum/common/math"
	"math/big"
)

func TestAESEncode(t *testing.T) {
	password := "1234567890"
	content := "我是刘飞"
	math.PaddedBigBytes(big.NewInt(0).SetBytes([]byte(content)), 32)
	result := make([]byte, len([]byte(content)))
	key, salt := KeyDerivation(password)
	iv := make([]byte, 16)
	rand.Read(iv)
	fmt.Println(hex.EncodeToString(key))
	aesBlock, _ := aes.NewCipher(key)
	stream := cipher.NewCTR(aesBlock, iv)
	stream.XORKeyStream(result, []byte(content))
	println(hex.EncodeToString(result), hex.EncodeToString(salt), hex.EncodeToString(iv))
}

func TestAESDecode(t *testing.T) {
	content, _ := hex.DecodeString("13588f97e1e0e1e375aee1df")
	keySrc := "5bd3e964fe6c35c100833dcbbd24378e"
	salt, _ := hex.DecodeString("cb9bed2c2d67b84dbce19b444c8d35f6ee1e665fa5778549d3f7083206f5c2b0")
	iv, _ := hex.DecodeString("06421ef9063ddeb5a0c58632813a2f7b")
	password := "1234567890"
	key := GetKey([]byte(password), salt)
	fmt.Println(hex.EncodeToString(key) == keySrc)
	aesBlock, _ := aes.NewCipher(key)
	//FIXME
	decrypter := cipher.NewCTR(aesBlock, iv)
	paddedPlainText := make([]byte, len(content))
	decrypter.XORKeyStream(paddedPlainText, content)
	fmt.Println(string(paddedPlainText))
}

func GetAesCTR() {

}
func KeyDerivation(password string) ([]byte, []byte) {
	salt := make([]byte, 32)
	rand.Read(salt)
	//TODO : salt 一般不直接用，而是还要在在外面gener
	b := pbkdf2.Key([]byte(password), salt, 1, 32, sha256.New)
	return pbkdf2.Key([]byte(password), b, 1, 16, sha256.New), salt
}
func GetKey(password []byte, salt []byte) []byte {
	b := pbkdf2.Key([]byte(password), salt, 1, 32, sha256.New)
	return pbkdf2.Key(password, b, 1, 16, sha256.New)
}
