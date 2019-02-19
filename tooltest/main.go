
package main

func test(x*int) {
	println(*x)
}

func main() {
	x:=0x100
	test(&x)
}
