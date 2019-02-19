package main

import (
	"testing"
	"os"
	"strconv"
)

func TestAppendFile(t *testing.T){

	if _,err := os.Stat("./append.txt") ;os.IsNotExist(err){
		os.Create("./append.txt")
	}
	file,err := os.OpenFile("./append.txt",os.O_APPEND|os.O_WRONLY, 0600)
	if err != nil{
		panic(err)
	}
	defer file.Close()
	for count := 0;count < 100;count++{
		if _,err := file.WriteString(strconv.Itoa(count) +"\r");err != nil{

			panic(err)
		}
	}
}
