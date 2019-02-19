package benchmarktest

import (
	"testing"
	"fmt"
)

func add(x, y int) byte {
	return byte(x + y)
}

func heap()[]byte{
	return make([]byte,1024 * 1024 * 10)
}

func BenchmarkAdd(b *testing.B) {
	fmt.Println("b.N=", b.N)
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		_ = add(1, 2)
		_ = heap()
	}
}
