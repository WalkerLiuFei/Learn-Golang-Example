package main

import (
	"sync"
	"testing"
	"fmt"
)

func TestLocalStorage(t *testing.T) {
	var wg sync.WaitGroup
	var gs [5]struct {
		// 用于实现类似TLS功能
		id     int // 编号
		result int // 返回值
	}

	for i := 0; i < len(gs); i++ {
		wg.Add(1)

		go func(id int) { // 使用参数避免闭包延迟求值
			defer wg.Done()

			gs[id].id = id
			gs[id].result = (id + 1) * 100
		}(i)
	}

	wg.Wait()
	fmt.Printf("%+v\n",gs)

}