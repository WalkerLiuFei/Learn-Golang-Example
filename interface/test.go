package main


/*type iface struct{
	tab *itab          // 类型信息
	data unsafe.Pointer      // 实际对象指针
}

type itab struct{
	inter *interfacetype // 接口类型
	_type *_type   // 实际对象类型
	fun   [1]uintptr     // 实际对象方法地址
}*/
type Ner interface{
	a()
	b(int)
	c(string)string
}

type N int
func(N) a() {}
func(*N)b(int) {}
func(*N)c(string)string{return"" }

func main() {
	var n N
	var t Ner= &n

	t.a()
}

