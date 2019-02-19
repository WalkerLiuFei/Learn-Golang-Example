package main

import (
	"testing"
	"fmt"
	"time"
)

func TestRateLimitingRate(t *testing.T){
	requests := make(chan  int ,5)
	for i := 1 ;i < 5;i++{
		requests <- i
	}
	close(requests)
	for req := range requests{
		fmt.Println("request",req,time.Now())
	}
}