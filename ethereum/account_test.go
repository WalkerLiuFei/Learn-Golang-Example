package ethereum

import (
	"testing"
	"github.com/ethereum/go-ethereum/crypto"
	"fmt"
)

func TestGenerateAddrFromKey(t *testing.T) {
	key, _ := crypto.HexToECDSA("7a7b9d6a32ea351f760928bbbf3ec093d3e8437905f1c027069be1484d46fe84")
	fmt.Println(crypto.PubkeyToAddress(key.PublicKey).Hex())
}
