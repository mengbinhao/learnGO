package test

import "testing"


func TestFib(t *testing.T) {
	// var a int = 1
	// var b int = 1

	// var (
	// 	a int = 1
	// 	b int = 2
	// )

	a := 1
	b := 1
	for i:=0; i<5; i++ {
		t.Log(" ", b)
		tmp := a
		a = b
		b = tmp + a
	}
}