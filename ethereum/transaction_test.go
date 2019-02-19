package ethereum

import (
	"testing"
	"github.com/ethereum/go-ethereum/core/types"
	"encoding/hex"
	"math/big"
	"github.com/ethereum/go-ethereum/common"
	"io/ioutil"
	"github.com/ethereum/go-ethereum/accounts/keystore"
)

type ChainID int

const (
	TEST_NET ChainID = 1988
	PRODICT_NET ChainID = 3
)
func TestMakeTx(t *testing.T){
	inputData := "608060405234801561001057600080fd5b50610141806100206000396000f300608060405260043610610041576000357c0100000000000000000000000000000000000000000000000000000000900463ffffffff168063dabd76e914610046575b600080fd5b6100d4600480360381019080803573ffffffffffffffffffffffffffffffffffffffff1690602001909291908035906020019092919080359060200190929190803590602001908201803590602001908080601f01602080910402602001604051908101604052809392919081815260200183838082843782019150505050505091929192905050506100ee565b604051808215151515815260200191505060405180910390f35b6000806040516020840160008287838a8c6187965a03f192505050809150509493505050505600a165627a7a72305820d05d3ac86cd86f5bc3bb7eeb4b58c30157a47d2b913be2276b14a330bc5d3bb10029"
	data,_ :=  hex.DecodeString(inputData)
	tx := types.NewContractCreation(1,big.NewInt(0),414188,big.NewInt(80000000000),data)
	createTransaction(tx)

}
func createTransaction(tx *types.Transaction) {

	data ,err := ioutil.ReadFile("./UTC--2018-08-08T09-43-55.077659100Z--e3b3b7e1aafb8a436673c9bb688fdac44500a20f")
	if err != nil{
		panic(err.Error())
	}

	priKey , err := keystore.DecryptKey(data,"1234")
	println(priKey.PrivateKey.D.String())
	/*
	if err != nil{
		panic(err.Error())
	}


	tx ,err  = types.SignTx(tx, types.NewEIP155Signer(big.NewInt(3)), priKey.PrivateKey)
	if err != nil{
		panic(err.Error())
	}

	rlpEncode, err := rlp.EncodeToBytes(tx)
	if err != nil{
		panic(err.Error())
	}
	fmt.Println(hex.EncodeToString(rlpEncode))*/
}

func TestCall(t *testing.T){
	inputData := ""
	data,err :=  hex.DecodeString(inputData)
	if err != nil{
		panic(err.Error())
	}
	to := common.HexToAddress("0xebbd27eb47d7fe38b0a29ccd33a5671434225744")
	tx := types.NewTransaction(4,to,big.NewInt(47271398999999800),500000,big.NewInt(1000000000),data)

	createTransaction(tx)
}


