package main

import "./internal"
import "fmt"

func init(){
	fmt.Println("init package")
}

func main(){
	fmt.Println("main")
	internal.TestInternal()
}