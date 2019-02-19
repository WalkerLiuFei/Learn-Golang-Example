package main


import (
	"testing"
	"bufio"
	"os"
	"strings"
	"github.com/ethereum/go-ethereum/crypto"
	"encoding/hex"
)


func TestModifyLog(t *testing.T) {
	inFile, _ := os.Open("./src_log")
	outFile, _ := os.Create("./dest_log")
	defer inFile.Close()
	defer outFile.Close()
	scanner := bufio.NewScanner(inFile)
	scanner.Split(bufio.ScanLines)
	flag := false
	for scanner.Scan() {
		str := scanner.Text()
		var parts []string
		if strings.Contains(str, "executeOperation") {
			parts = strings.Split(str, "executeOperation")
		} else if strings.Contains(str, "stack value") {
			parts = strings.Split(str, "stack value")
		}
		str = strings.Trim(parts[len(parts)-1], " ")
		str = strings.Split(str,"=")[1]
		outFile.WriteString(str)
		if !flag{
			outFile.WriteString("\n")
		}else {
			outFile.WriteString("\n\n")
		}
		flag = !flag

	}
}

func TestKeccakHash(t *testing.T){
	println(hex.EncodeToString(crypto.Keccak256([]byte("Transfer(address,address,uint256)"))))
}
