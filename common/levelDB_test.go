package main

import (
	"testing"
	"github.com/ethereum/go-ethereum/ethdb"
	"strconv"
	"fmt"
)

func TestLevelDB(t *testing.T){
	var table ethdb.Database
	dbfile := "./db_test"
	db,err := ethdb.NewLDBDatabase(dbfile,16,16)
	defer db.Close()
	if err != nil{
		panic(err)
	}
	for count := 0;count < 10000;count++{
		if count % 100 == 0{
			table = ethdb.NewTable(db,strconv.Itoa(count))
		}
		table.Put([]byte("_key_"+strconv.Itoa(count)), []byte("value_"+strconv.Itoa(count)))
	}
	read(db,"100")
}

func read(db *ethdb.LDBDatabase,prefix string){
	iterator := db.NewIteratorWithPrefix([]byte(prefix))
	defer iterator.Release()
	for iterator.Next(){
		value_pair := iterator.Value()
		fmt.Println(string(value_pair))
	}
}

func TestReWrite(t *testing.T){
	dbfile := "./db_test"
	db,err := ethdb.NewLDBDatabase(dbfile,16,16)
	defer db.Close()
	if err != nil{
		panic(err)
	}
	for count := 0;count < 100;count++{
		db.Put([]byte("_key_"+strconv.Itoa(count)), []byte("value2_"+strconv.Itoa(count)))
	}

	for count := 0;count < 100;count++{
		db.Put([]byte("_key_"+strconv.Itoa(count)), []byte("value_"+strconv.Itoa(count)))
	}
	result ,_ := db.Get([]byte("_key_"+strconv.Itoa(50)))
	fmt.Println(string(result))
}