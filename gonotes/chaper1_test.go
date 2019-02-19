package gonotes

import (
	"testing"
	"fmt"
)

func TestReferenceType(t *testing.T) {
	s := make([]int, 0, 100)
	s = append(s, 100)
}

func TestEvent(t *testing.T) {
	type (
		user struct {
			name string
			age  int8
		}
		event func(string) bool
	)
	u := user{"Tom", 20}
	fmt.Println(u)
	var f event = func(s string) bool {
		println(s)
		return s != ""
	}
	f("abc")
}
